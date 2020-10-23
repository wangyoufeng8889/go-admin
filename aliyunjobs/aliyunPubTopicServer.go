package aliyunjobs

import (
	"encoding/base64"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/iot"
	"strings"
)
//参数说明，请参见AMQP客户端接入说明文档。
var Product_Key = "a188ncLCDbX"
var Product_Secert = "76VASKZrFqnVHgix"
var Topic ="/user/get"
var DeviceName = "LSDDTU01012008190007"
func AliyunPubTopic(devicename string,buf string)(err error)  {
	AliyunContorClient, err := iot.NewClientWithAccessKey(region, accessKey,accessSecret)

	request := iot.CreatePubRequest()
	request.AcceptFormat = "json"
	var builder strings.Builder
	builder.WriteString("/")
	builder.WriteString(Product_Key)
	builder.WriteString("/")
	builder.WriteString(devicename)
	builder.WriteString(Topic)
	request.TopicFullName = builder.String()
	payload := "{\"message\":\""+buf+"\"}"
	request.MessageContent = base64.StdEncoding.EncodeToString([]byte(payload))
	request.Qos = "1"
	request.ProductKey=Product_Key

	response, err := AliyunContorClient.Pub(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
	return err
}