package aliyunjobs

import (
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"go-admin/global"
	"go-admin/tools/config"
	"pack.ag/amqp"
	"strings"
	"time"
)
//参数说明，请参见AMQP客户端接入说明文档。
var uid string
var accessKey string
var accessSecret string
var region string
const consumerGroupId = "DEFAULT_GROUP"
//iotInstanceId：购买的实例请填写实例ID，公共实例请填空字符串""。
const iotInstanceId = ""
//控制台服务端订阅中消费组状态页客户端ID一栏将显示clientId参数。
//建议使用机器UUID、MAC地址、IP等唯一标识等作为clientId。便于您区分识别不同的客户端。
var clientId string
var messageChan chan ModbusMessage
var AliyunClient *sdk.Client
func AliyunServerRun() {
	uid = config.AliyunConfig.Uid
	accessKey = config.AliyunConfig.AccessKey
	accessSecret = config.AliyunConfig.AccessSecret
	region = config.AliyunConfig.Region
	clientId = config.AliyunConfig.ClientId
	//接入域名，请参见AMQP客户端接入说明文档。
	address := fmt.Sprintf("amqps://%s.iot-amqp.%s.aliyuncs.com:5671", uid, region)
	timestamp := time.Now().Nanosecond() / 1000000
	//userName组装方法，请参见AMQP客户端接入说明文档。
	userName := fmt.Sprintf("%s|authMode=aksign,signMethod=Hmacsha1,consumerGroupId=%s,authId=%s,iotInstanceId=%s,timestamp=%d|",
		clientId, consumerGroupId, accessKey, iotInstanceId, timestamp)
	stringToSign := fmt.Sprintf("authId=%s&timestamp=%d", accessKey, timestamp)
	hmacKey := hmac.New(sha1.New, []byte(accessSecret))
	hmacKey.Write([]byte(stringToSign))
	//计算签名，password组装方法，请参见AMQP客户端接入说明文档。
	password := base64.StdEncoding.EncodeToString(hmacKey.Sum(nil))

	amqpManager := &AmqpManager{
		address:address,
		userName:userName,
		password:password,
	}

	AliyunClient, err := sdk.NewClientWithAccessKey(region ,accessKey,accessSecret)
	if err != nil {
		// Handle exceptions
		panic(err)
	}
	global.Logger.Info(AliyunClient)

	//如果需要做接受消息通信或者取消操作，从Background衍生context。
	ctx := context.Background()
	global.Logger.Info("start run aluyun amqp server")
	//初始化channel
	messageChan = make(chan ModbusMessage,100)
	amqpManager.startReceiveMessage(ctx)
}

//业务函数。用户自定义实现，该函数被异步执行，请考虑系统资源消耗情况。
func (am *AmqpManager) processMessage(message *amqp.Message) {
	//fmt.Printf("topic=%s\r\n",message.ApplicationProperties["topic"])
	//fmt.Printf("generateTime=%d\r\n",message.ApplicationProperties["generateTime"])

	//global.Logger.Info("data received:", string(message.GetData()), " properties:", message.ApplicationProperties)
	//1、解析产品ID 设备ID topic

	var modbusMessage ModbusMessage
	temp := message.ApplicationProperties["topic"]
	var topicdata string
	switch v:=temp.(type) {
	case string:
		topicdata = v
	}
	topicslic:=strings.Split(topicdata,"/")
	if topicslic[1] == "as" && topicslic[2] == "mqtt" && topicslic[3] == "status" {
		modbusMessage.ProductID =  topicslic[4]
		modbusMessage.DtuID =topicslic[5]
		modbusMessage.Topic ="/"+topicslic[1]+"/"+topicslic[2]+"/"+topicslic[3]
		//2、解析payload
		var messagePayload struct{LastTime string
									UtcLastTime string
									ClientIp string
									UtcTime string
									Time string
									ProductKey string
									DeviceName string
									Status string}
		if err:=json.Unmarshal(message.GetData(),&messagePayload);err!=nil{
			fmt.Println("err=",err)
		}
		global.Logger.Info(messagePayload.Status)
		if messagePayload.Status == "offline" {
			modbusMessage.Payload = []byte{0}
		}else {
			modbusMessage.Payload = []byte{1}
		}

		switch v:=message.ApplicationProperties["generateTime"].(type) {
		case int64:
			modbusMessage.Timestamp = v
		}
		messageChan <- modbusMessage
		go ModbusServer(messageChan)
	}else if topicslic[1] == "a188ncLCDbX"{
		modbusMessage.ProductID =  topicslic[1]
		modbusMessage.DtuID =topicslic[2]
		modbusMessage.Topic ="/"+topicslic[3]+"/"+topicslic[4]
		//2、解析payload
		var messagePayload struct{Message string}
		if err:=json.Unmarshal(message.GetData(),&messagePayload);err!=nil{
			fmt.Println("err=",err)
		}
		//global.Logger.Info(messagePayload.Message)
		bytePayload,err:=hex.DecodeString(messagePayload.Message)
		if err!=nil {
			fmt.Println(err)
		}

		modbusMessage.Payload = bytePayload
		switch v:=message.ApplicationProperties["generateTime"].(type) {
		case int64:
			modbusMessage.Timestamp = v
		}
		messageChan <- modbusMessage
		go ModbusServer(messageChan)
	}

}

type AmqpManager struct {
	address     string
	userName     string
	password     string
	client         *amqp.Client
	session     *amqp.Session
	receiver     *amqp.Receiver

}

func (am *AmqpManager) startReceiveMessage(ctx context.Context)  {

	childCtx, _ := context.WithCancel(ctx)
	err := am.generateReceiverWithRetry(childCtx)
	if nil != err {
		return
	}
	defer func() {
		am.receiver.Close(childCtx)
		am.session.Close(childCtx)
		am.client.Close()
	}()

	for {

		//阻塞接受消息，如果ctx是background则不会被打断。
		message, err := am.receiver.Receive(ctx)

		if nil == err {
			go am.processMessage(message)
			message.Accept()
		} else {
			global.Logger.Info("amqp receive data error:", err)

			//如果是主动取消，则退出程序。
			select {
			case <- childCtx.Done(): return
			default:
			}

			//非主动取消，则重新建立连接。
			err := am.generateReceiverWithRetry(childCtx)
			if nil != err {
				return
			}

		}
	}

}

func (am *AmqpManager) generateReceiverWithRetry(ctx context.Context) error {

	//退避重试，从10ms依次x2，直到20s。
	duration := 10 * time.Millisecond
	maxDuration := 20000 * time.Millisecond
	times := 1

	//异常情况，退避重连。
	for {
		select {
		case <- ctx.Done(): return amqp.ErrConnClosed
		default:
		}

		err := am.generateReceiver()
		if nil != err {
			time.Sleep(duration)
			if duration < maxDuration {
				duration *= 2
			}
			global.Logger.Info("amqp connect retry,times:", times, ",duration:", duration)
			times ++
		} else {
			global.Logger.Info("amqp connect init success")
			return nil
		}
	}
}

//由于包不可见，无法判断conn和session状态，重启连接获取。
func (am *AmqpManager) generateReceiver() error {

	if am.session != nil {
		receiver, err := am.session.NewReceiver(
			amqp.LinkSourceAddress("/queue-name"),
			amqp.LinkCredit(20),
		)
		//如果断网等行为发生，conn会关闭导致session建立失败，未关闭连接则建立成功。
		if err == nil {
			am.receiver = receiver
			return nil
		}
	}

	//清理上一个连接。
	if am.client != nil {
		am.client.Close()
	}

	client, err := amqp.Dial(am.address, amqp.ConnSASLPlain(am.userName, am.password), )
	if err != nil {
		return err
	}
	am.client = client

	session, err := client.NewSession()
	if err != nil {
		return err
	}
	am.session = session

	receiver, err := am.session.NewReceiver(
		amqp.LinkSourceAddress("/queue-name"),
		amqp.LinkCredit(20),
	)
	if err != nil {
		return err
	}
	am.receiver = receiver

	return nil
}