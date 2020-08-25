package aliyunjobs

import (
	"fmt"
	"github.com/iancoleman/strcase"
	orm "go-admin/global"
	"go-admin/models/batterymanage"
	"go-admin/tools/gps"
	"reflect"
	"time"
	//"time"
)

type ModbusMessage struct {
	Payload []byte
	Timestamp int64
	ProductID string
	DtuID string
	PkgID string
	Topic string
}
//mao[dtu_id]pkg_id
var Dtu_Pkg_map map[string]string

func readDtuPkgMapFromDB()  {
	time.Sleep(5000 * time.Millisecond)
	var dtuPkg_list []batterymanage.Dtu_specInfo

	orm.Eloquent.Find(&dtuPkg_list)
	for _, dtupkg := range dtuPkg_list{
		Dtu_Pkg_map[dtupkg.Dtu_id] = dtupkg.Pkg_id
	}


}
func init()  {
	Dtu_Pkg_map = make(map[string]string)
	go readDtuPkgMapFromDB()
}
func ModbusServer(msg chan ModbusMessage) {
	message := <-msg
	if message.Topic == "/user/update" {
		fmt.Println(message.DtuID)
		//addr, reglen, reg, err := modbusParseTcp(message)
		addr, reglen, reg, err := modbusParseTcp(message)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("ModbusServer=", addr, reglen, reg)
	}else if message.Topic == "/as/mqtt/status"{
		aliyunOnOffprocess(message)
	}else {
		fmt.Println("ModbusServer topic is err=",message.Topic)
	}
}
func Struct2Map(obj interface{},ingore []int ) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	var i,j int = 0,0
	var data = make(map[string]interface{})
	for i = 0; i < t.NumField(); i++ {
		if i == ingore[j] || i == t.NumField() + ingore[j]{
			j++
		}else{
			data[strcase.ToSnake(t.Field(i).Name)] = v.Field(i).Interface()
		}
	}
	return data
}
func aliyunOnOffprocess(msg ModbusMessage)  {
	var dtu_specInfo,dtu_specInfotemp batterymanage.Dtu_specInfo
	dtu_specInfo.Dtu_uptime=time.Unix(msg.Timestamp/1000, 0)
	dtu_specInfo.Dtu_aliyunStatus = uint8(msg.Payload[0])

	dtu_specInfotemp = dtu_specInfo
	//orm.Eloquent.Create(&dtu_aliyun)
	if err:= orm.Eloquent.Where(&batterymanage.Dtu_specInfo{Dtu_id:msg.DtuID}).FirstOrCreate(&dtu_specInfotemp).Error;err!=nil{
		fmt.Println(err)
	}else {
		dtu_aliyunmap:=Struct2Map(dtu_specInfo,[]int{0,2,3,4,5,6,7,8,9,10,11,12,13,-3,-2,-1})
		if err := orm.Eloquent.Model(batterymanage.Dtu_specInfo{}).Where(&batterymanage.Dtu_specInfo{Dtu_id:msg.DtuID}).Updates(dtu_aliyunmap).Error;err!=nil{
			fmt.Println(err)
		}
		fmt.Println("dtu=",msg.DtuID,"|aliyun=",dtu_specInfo.Dtu_aliyunStatus,time.Now())
	}
}
func modbusParseTcp(msg ModbusMessage)(addr uint16,reglen uint8,reg []uint16,err error)  {
	if len(msg.Payload) < 8 {
		return 0, 0, nil,fmt.Errorf("payload is short")
	}
	addr = uint16(msg.Payload[0])<<8 +uint16(msg.Payload[1])
	datalen := uint16(msg.Payload[4])<<8 +uint16(msg.Payload[5])
	datalen = datalen+6
	if int(datalen) != len(msg.Payload) {
		return addr, reglen, reg,fmt.Errorf("payload len is err")
	}
	devId :=msg.Payload[6]
	if devId != 0x01 {
		return addr, reglen, reg,fmt.Errorf("payload slaveid is err")
	}
	cmd:=msg.Payload[7]
	if cmd == 0x03 {
		//读寄存器
		byteNbr:=int(msg.Payload[8])
		reglen= uint8(byteNbr/2)
		for i :=0;i<byteNbr;i+=2 {
			reg = append(reg,uint16(msg.Payload[i+9])<<8+uint16(msg.Payload[i+10]))
		}
		switch addr {
		case 30000:
			modbusProcess30000(reg,reglen,msg)
			break
		case 30027:
			modbusProcess30027(reg,reglen,msg)
			break
		case 30100:
			modbusProcess30100(reg,reglen,msg)
			break
		case 30113:
			modbusProcess30113(reg,reglen,msg)
		case 30123:
			modbusProcess30123(reg,reglen,msg)
			break
		case 30200:
			modbusProcess30200(reg,reglen,msg)
			break
		case 30300:
			modbusProcess30300(reg,reglen,msg)
			break
		case 30400:
			break
		case 30500:
			modbusProcess30500(reg,reglen,msg)
			break
		case 30600:
			modbusProcess30600(reg,reglen,msg)
			break
		case 30647:
			modbusProcess30647(reg,reglen,msg)
			break
		case 30700:
			break
		default:
			fmt.Println("default 30xxx")
			break
		}
	}else if cmd == 0x10 {
		//写寄存器
	}else {
		return addr, reglen, reg,fmt.Errorf("payload cmd is err")
	}
	return addr, reglen, reg,nil
}
func sliceUin16Tobyte(sl_in []uint16)(sl_out []byte,err error)  {
	var temp []byte
	sl_in_len := len(sl_in)
	for i:=0;i<sl_in_len;i++ {
		temp = append(temp,byte(sl_in[i]>>8),byte(sl_in[i]))
	}
	sl_in_len = len(temp)
	for i:=0;i<sl_in_len;i++{
		if temp[i] != 0x00{
			sl_out = append(sl_out,temp[i])
		}
	}
	return sl_out,nil
}
func Dtu_BMS_map_Init(msg ModbusMessage)(bool)  {
	var dtuPkg_list batterymanage.Dtu_specInfo

	orm.Eloquent.Where(&batterymanage.Dtu_specInfo{Dtu_id: msg.DtuID}).First(&dtuPkg_list)
	if len(dtuPkg_list.Pkg_id)>0 {
		Dtu_Pkg_map[msg.DtuID] = dtuPkg_list.Pkg_id
		return false
	}else{
		//Dtu_Pkg_map[msg.DtuID] = ""
		return false
	}
}
func modbusProcess30000(reg []uint16,reglen uint8,msg ModbusMessage)  {
	regTemp, _ := sliceUin16Tobyte(reg[0:10])
	pkg_id:= string(regTemp)
	regTemp, _ = sliceUin16Tobyte(reg[10:20])
	bms_id:= string(regTemp)

	var bms_specinfo batterymanage.Bms_specInfo
	bms_specinfo.Dtu_uptime = time.Unix(msg.Timestamp/1000, 0)
	bms_specinfo.Pkg_id=pkg_id
	bms_specinfo.Dtu_id=msg.DtuID
	bms_specinfo.Bms_id=bms_id
	bms_specinfo.Pkg_count=uint8(reg[20]>>8)
	bms_specinfo.Pkg_type=uint8(reg[20])
	bms_specinfo.Pkg_capacity=uint16(reg[21])
	bms_specinfo.Pkg_nominalVoltage=uint16(reg[22])
	bms_specinfo.Pkg_ntcCount=uint8(reg[23]>>8)
	var manufactureyear int = int(int(2000) + int(uint8(reg[23])))
	loc, _ := time.LoadLocation("Local")
	bms_specinfo.Pkg_manufactureDate=time.Date(manufactureyear,time.Month(uint8(reg[24]>>8)),int(uint8(reg[24])),0,0,0,0,loc)
	bms_specinfo.Bms_hardVer=uint8(reg[25]>>8)
	bms_specinfo.Bms_softVer=uint8(reg[25])
	data1 := reg[26]/0xFF
	data2 := reg[26]%0xFF
	bms_specinfo.Bms_protocolVer =  fmt.Sprintf("%d.%02d",data1,data2)
	bms_specinfotemp := bms_specinfo
	if err:=orm.Eloquent.Where(&batterymanage.Bms_specInfo{Pkg_id: pkg_id}).FirstOrCreate(&bms_specinfotemp).Error;err != nil {
		fmt.Println(err)
	}else {
		bms_specinfomap:=Struct2Map(bms_specinfo,[]int{0,-3,-2,-1})
		if err:=orm.Eloquent.Model(batterymanage.Bms_specInfo{}).Where(&batterymanage.Bms_specInfo{Pkg_id: pkg_id}).Updates(bms_specinfomap).Error;err != nil {
			fmt.Println(err)
		}
	}
	if reglen == 72 {
		regTemp, _ = sliceUin16Tobyte(reg[27:37])
		dtu_id:= string(regTemp)
		regTemp, _ = sliceUin16Tobyte(reg[41:51])
		dtu_devid:= string(regTemp)
		regTemp, _ = sliceUin16Tobyte(reg[51:61])
		dtu_sim_iccid:= string(regTemp)
		regTemp, _ = sliceUin16Tobyte(reg[61:71])
		dtu_imei:= string(regTemp)
		if dtu_id != msg.DtuID {
			fmt.Println("dtu id is err")
		}else {
			Dtu_Pkg_map[dtu_id]=pkg_id
		}

		var dtu_specinfo batterymanage.Dtu_specInfo
		dtu_specinfo.Dtu_uptime=bms_specinfo.Dtu_uptime
		dtu_specinfo.Dtu_id= dtu_id
		dtu_specinfo.Pkg_id= pkg_id
		dtu_specinfo.Dtu_type= uint8(reg[37]>>8)
		dtu_specinfo.Dtu_setupType= uint8(reg[37])
		dtu_specinfo.Dtu_coreVer=uint16(reg[38])
		dtu_specinfo.Dtu_hardVer=uint8(reg[39]>>8)
		dtu_specinfo.Dtu_softVer=uint8(reg[39])
		data1 := reg[40]/0x100
		data2 := reg[40]%0x100
		dtu_specinfo.Dtu_protocolVer =  fmt.Sprintf("%d.%02d",data1,data2)
		dtu_specinfo.Dtu_devID=dtu_devid
		dtu_specinfo.Dtu_simIccid=dtu_sim_iccid
		dtu_specinfo.Dtu_imei=dtu_imei
		dtu_specinfo.Dtu_bmsBindStatus=uint8(reg[71])
		dtu_specinfotemp:=dtu_specinfo
		if err:=orm.Eloquent.Where(&batterymanage.Dtu_specInfo{Dtu_id: dtu_id}).FirstOrCreate(&dtu_specinfotemp).Error;err != nil {
			fmt.Println(err)
		}else {
			dtu_specinfomap:=Struct2Map(dtu_specinfo,[]int{0,-4,-3,-2,-1})
			if err:=orm.Eloquent.Model(batterymanage.Dtu_specInfo{}).Where(&batterymanage.Dtu_specInfo{Dtu_id: dtu_id}).Updates(dtu_specinfomap).Error;err != nil {
				fmt.Println(err)
			}
		}
	}
}
func modbusProcess30027(reg []uint16,reglen uint8,msg ModbusMessage)  {
	regTemp, _ := sliceUin16Tobyte(reg[0:10])
	dtu_id:= string(regTemp)
	regTemp, _ = sliceUin16Tobyte(reg[14:24])
	dtu_devid:= string(regTemp)
	regTemp, _ = sliceUin16Tobyte(reg[24:34])
	dtu_sim_iccid:= string(regTemp)
	regTemp, _ = sliceUin16Tobyte(reg[34:44])
	dtu_imei:= string(regTemp)
	if dtu_id != msg.DtuID {
		//需要加上错误检测
		fmt.Println("dtu id is not aliyun id")
	}else {
		Dtu_Pkg_map[dtu_id]="0"
	}
	pkg_id:= Dtu_Pkg_map[msg.DtuID]
	if len(pkg_id)<5{
		res := Dtu_BMS_map_Init(msg)
		if res != true {
			pkg_id=""
		}else {
			pkg_id=Dtu_Pkg_map[msg.DtuID]
		}
	}

	var dtu_specinfo batterymanage.Dtu_specInfo
	dtu_specinfo.Dtu_uptime=time.Unix(msg.Timestamp/1000, 0)
	dtu_specinfo.Dtu_id= dtu_id
	dtu_specinfo.Dtu_type= uint8(reg[10]>>8)
	dtu_specinfo.Dtu_setupType= uint8(reg[10])
	dtu_specinfo.Dtu_coreVer=uint16(reg[11])
	dtu_specinfo.Dtu_hardVer=uint8(reg[12]>>8)
	dtu_specinfo.Dtu_softVer=uint8(reg[12])
	data1 := reg[13]/0x100
	data2 := reg[13]%0x100
	dtu_specinfo.Dtu_protocolVer =  fmt.Sprintf("%d.%02d",data1,data2)
	dtu_specinfo.Dtu_devID=dtu_devid
	dtu_specinfo.Dtu_simIccid=dtu_sim_iccid
	dtu_specinfo.Dtu_imei=dtu_imei
	dtu_specinfo.Dtu_bmsBindStatus=uint8(reg[44])
	if dtu_specinfo.Dtu_bmsBindStatus == 0 {
		dtu_specinfo.Pkg_id= ""
		Dtu_Pkg_map[msg.DtuID] = ""
	}else {
		dtu_specinfo.Pkg_id= pkg_id
	}
	dtu_specinfotemp:=dtu_specinfo
	if err:=orm.Eloquent.Where(&batterymanage.Dtu_specInfo{Dtu_id: dtu_id}).FirstOrCreate(&dtu_specinfotemp).Error;err != nil {
		fmt.Println(err)
	}else {
		dtu_specinfomap:=Struct2Map(dtu_specinfo,[]int{0,-4,-3,-2,-1})
		if err:=orm.Eloquent.Model(batterymanage.Dtu_specInfo{}).Where(&batterymanage.Dtu_specInfo{Dtu_id: dtu_id}).Updates(dtu_specinfomap).Error;err != nil {
			fmt.Println(err)
		}
	}
}
func modbusProcess30100(reg []uint16,reglen uint8,msg ModbusMessage)  {
	pkg_id:= Dtu_Pkg_map[msg.DtuID]
	if len(pkg_id)<5{
		res := Dtu_BMS_map_Init(msg)
		if res != true {
			pkg_id=""
		}else {
			pkg_id=Dtu_Pkg_map[msg.DtuID]
		}
	}
	var bms_statusinfo batterymanage.Bms_statusInfo
	bms_statusinfo.Dtu_uptime=time.Unix(msg.Timestamp/1000, 0)
	bms_statusinfo.Pkg_id= pkg_id
	bms_statusinfo.Dtu_id=msg.DtuID
	bms_statusinfo.Bms_chargeStatus=uint8(reg[0]>>8)
	bms_statusinfo.Bms_soc=uint8(reg[0])
	bms_statusinfo.Bms_errStatus=uint8(reg[1]>>8)
	bms_statusinfo.Bms_errNbr=uint8(reg[1])
	bms_statusinfo.Bms_errCode=uint32(reg[2])<<16+uint32(reg[3])
	bms_statusinfo.Bms_voltage=uint16(reg[4])
	bms_statusinfo.Bms_current=uint16(reg[5])
	bms_statusinfo.Bms_maxCellVoltage=uint16(reg[6])
	bms_statusinfo.Bms_minCellVoltage=uint16(reg[7])
	bms_statusinfo.Bms_averageCellVoltage=uint16(reg[8])
	bms_statusinfo.Bms_maxTemperature=uint8(reg[9]>>8) - 40
	bms_statusinfo.Bms_minTemperature=uint8(reg[9]) - 40
	bms_statusinfo.Bms_mosTemperature=uint8(reg[10]>>8) - 40
	bms_statusinfo.Bms_balanceResistance=uint8(reg[10]) - 40
	bms_statusinfo.Bms_chargeMosStatus=uint8(reg[11]>>8)
	bms_statusinfo.Bms_dischargeMosStatus=uint8(reg[11])
	bms_statusinfo.Bms_otaBufStatus=uint8(reg[12]>>8)
	bms_statusinfo.Bms_magneticCheck=uint8(reg[12])


	var bms_statusinfolog batterymanage.Bms_statusInfoLog
	bms_statusinfolog.Dtu_uptime =	bms_statusinfo.Dtu_uptime
	bms_statusinfolog.Pkg_id =  bms_statusinfo.Pkg_id
	bms_statusinfolog.Dtu_id=bms_statusinfo.Dtu_id
	bms_statusinfolog.Bms_chargeStatus=bms_statusinfo.Bms_chargeStatus
	bms_statusinfolog.Bms_soc=bms_statusinfo.Bms_soc
	bms_statusinfolog.Bms_errStatus=bms_statusinfo.Bms_errStatus
	bms_statusinfolog.Bms_errNbr=bms_statusinfo.Bms_errNbr
	bms_statusinfolog.Bms_errCode=bms_statusinfo.Bms_errCode
	bms_statusinfolog.Bms_voltage=bms_statusinfo.Bms_voltage
	bms_statusinfolog.Bms_current=bms_statusinfo.Bms_current
	bms_statusinfolog.Bms_maxCellVoltage=bms_statusinfo.Bms_maxCellVoltage
	bms_statusinfolog.Bms_minCellVoltage=bms_statusinfo.Bms_minCellVoltage
	bms_statusinfolog.Bms_averageCellVoltage=bms_statusinfo.Bms_averageCellVoltage
	bms_statusinfolog.Bms_maxTemperature=bms_statusinfo.Bms_maxTemperature
	bms_statusinfolog.Bms_minTemperature=bms_statusinfo.Bms_minTemperature
	bms_statusinfolog.Bms_mosTemperature=bms_statusinfo.Bms_mosTemperature
	bms_statusinfolog.Bms_balanceResistance=bms_statusinfo.Bms_balanceResistance
	bms_statusinfolog.Bms_chargeMosStatus=bms_statusinfo.Bms_chargeMosStatus
	bms_statusinfolog.Bms_dischargeMosStatus=bms_statusinfo.Bms_dischargeMosStatus
	bms_statusinfolog.Bms_otaBufStatus=bms_statusinfo.Bms_otaBufStatus
	bms_statusinfolog.Bms_magneticCheck=bms_statusinfo.Bms_magneticCheck
	if err:=orm.Eloquent.Create(&bms_statusinfolog).Error;err!=nil{
		fmt.Println(err)
	}
	bms_statusinfotemp:=bms_statusinfo
	if err:=orm.Eloquent.Where(&batterymanage.Bms_statusInfo{Pkg_id: pkg_id}).FirstOrCreate(&bms_statusinfotemp).Error;err != nil {
		fmt.Println(err)
	}else {
		bms_statusinfomap:=Struct2Map(bms_statusinfo,[]int{0,-3,-2,-1})
		if err:=orm.Eloquent.Model(batterymanage.Bms_statusInfo{}).Where(&batterymanage.Bms_statusInfo{Pkg_id: pkg_id}).Updates(bms_statusinfomap).Error;err != nil {
			fmt.Println(err)
		}
	}
	if reglen == 25 {
		var dtu_statusinfo batterymanage.Dtu_statusInfo
		dtu_statusinfo.Dtu_uptime = time.Unix(msg.Timestamp/1000, 0)
		dtu_statusinfo.Pkg_id=pkg_id
		dtu_statusinfo.Dtu_id=msg.DtuID
		if uint8(reg[13] >> 8) == 'N' {
			dtu_statusinfo.Dtu_latitudeType = "N"
		}else{
			dtu_statusinfo.Dtu_latitudeType = "S"
		}
		if uint8(reg[13]) == 'E' {
			dtu_statusinfo.Dtu_longitudeType = "E"
		}else{
			dtu_statusinfo.Dtu_longitudeType = "W"
		}
		Dtu_latitude := int(reg[14])<<16 + int(reg[15])
		Dtu_longitude := int(reg[16])<<16 + int(reg[17])
		var latitude_WGS84 float64 = float64(Dtu_latitude)/1000000
		var longitude_WGS84 float64 = float64(Dtu_longitude)/1000000

		latitude_GCJ02,longitude_GCJ02 := gps.WGS84ToGCJ02(latitude_WGS84,longitude_WGS84)
		dtu_statusinfo.Dtu_latitude =  fmt.Sprint(latitude_GCJ02)
		dtu_statusinfo.Dtu_longitude = fmt.Sprint(longitude_GCJ02)

		dtu_statusinfo.Dtu_csq = uint8(reg[18] >> 8)
		dtu_statusinfo.Dtu_locateMode = uint8(reg[19] >> 8)
		dtu_statusinfo.Dtu_gpsSateCnt = uint8(reg[19])
		dtu_statusinfo.Dtu_speed=uint16(reg[20])/100
		dtu_statusinfo.Dtu_altitude=uint16(reg[21])
		dtu_statusinfo.Dtu_pluginVoltage=uint8(reg[22] >> 8)
		dtu_statusinfo.Dtu_selfInVoltage=uint8(reg[22])
		dtu_statusinfo.Dtu_errStatus=uint8(reg[23] >> 8)
		dtu_statusinfo.Dtu_errNbr=uint8(reg[23])
		dtu_statusinfo.Dtu_errCode=uint16(reg[24])

		var dtu_statusinfolog batterymanage.Dtu_statusInfoLog
		dtu_statusinfolog.Dtu_uptime = dtu_statusinfo.Dtu_uptime
		dtu_statusinfolog.Dtu_id = dtu_statusinfo.Dtu_id
		dtu_statusinfolog.Pkg_id = dtu_statusinfo.Pkg_id
		dtu_statusinfolog.Dtu_latitudeType = dtu_statusinfo.Dtu_latitudeType
		dtu_statusinfolog.Dtu_longitudeType = dtu_statusinfo.Dtu_longitudeType
		dtu_statusinfolog.Dtu_latitude = dtu_statusinfo.Dtu_latitude
		dtu_statusinfolog.Dtu_longitude = dtu_statusinfo.Dtu_longitude
		dtu_statusinfolog.Dtu_csq = dtu_statusinfo.Dtu_csq
		dtu_statusinfolog.Dtu_locateMode = dtu_statusinfo.Dtu_locateMode
		dtu_statusinfolog.Dtu_gpsSateCnt = dtu_statusinfo.Dtu_gpsSateCnt
		dtu_statusinfolog.Dtu_speed = dtu_statusinfo.Dtu_speed
		dtu_statusinfolog.Dtu_altitude = dtu_statusinfo.Dtu_altitude
		dtu_statusinfolog.Dtu_pluginVoltage = dtu_statusinfo.Dtu_pluginVoltage
		dtu_statusinfolog.Dtu_selfInVoltage = dtu_statusinfo.Dtu_selfInVoltage
		dtu_statusinfolog.Dtu_errStatus = dtu_statusinfo.Dtu_errStatus
		dtu_statusinfolog.Dtu_errNbr = dtu_statusinfo.Dtu_errNbr
		dtu_statusinfolog.Dtu_errCode = dtu_statusinfo.Dtu_errCode
		if err:=orm.Eloquent.Create(&dtu_statusinfolog).Error;err!=nil{
			fmt.Println(err)
		}
		dtu_statusinfotemp:=dtu_statusinfo
		if err:=orm.Eloquent.Where(&batterymanage.Dtu_statusInfo{Dtu_id: msg.DtuID}).FirstOrCreate(&dtu_statusinfotemp).Error;err != nil {
			fmt.Println(err)
		}else {
			dtu_statusinfomap:=Struct2Map(dtu_statusinfo,[]int{0,-3,-2,-1})
			if err:=orm.Eloquent.Model(batterymanage.Dtu_statusInfo{}).Where(&batterymanage.Dtu_statusInfo{Dtu_id: msg.DtuID}).Updates(dtu_statusinfomap).Error;err != nil {
				fmt.Println(err)
			}
		}
	}
}
func modbusProcess30113(reg []uint16,reglen uint8,msg ModbusMessage)  {
	pkg_id:= Dtu_Pkg_map[msg.DtuID]
	if len(pkg_id)<5{
		res:= Dtu_BMS_map_Init(msg)
		if res != true {
			fmt.Println("find no bmsID")
			pkg_id=""
		}else {
			pkg_id= Dtu_Pkg_map[msg.DtuID]
		}
	}
	var dtu_statusinfo batterymanage.Dtu_statusInfo
	dtu_statusinfo.Dtu_uptime = time.Unix(msg.Timestamp/1000, 0)
	dtu_statusinfo.Pkg_id=pkg_id
	dtu_statusinfo.Dtu_id=msg.DtuID
	if uint8(reg[0] >> 8) == 'N' {
		dtu_statusinfo.Dtu_latitudeType = "N"
	}else{
		dtu_statusinfo.Dtu_latitudeType = "S"
	}
	if uint8(reg[0]) == 'E' {
		dtu_statusinfo.Dtu_longitudeType = "E"
	}else{
		dtu_statusinfo.Dtu_longitudeType = "W"
	}

	Dtu_latitude := int(reg[1])<<16 + int(reg[2])
	Dtu_longitude := int(reg[3])<<16 + int(reg[4])
	var latitude_WGS84 float64 = float64(Dtu_latitude)/1000000
	var longitude_WGS84 float64 = float64(Dtu_longitude)/1000000

	latitude_GCJ02,longitude_GCJ02 := gps.WGS84ToGCJ02(latitude_WGS84,longitude_WGS84)
	dtu_statusinfo.Dtu_latitude =  fmt.Sprint(latitude_GCJ02)
	dtu_statusinfo.Dtu_longitude = fmt.Sprint(longitude_GCJ02)
	dtu_statusinfo.Dtu_csq = uint8(reg[5] >> 8)
	dtu_statusinfo.Dtu_locateMode = uint8(reg[6] >> 8)
	dtu_statusinfo.Dtu_gpsSateCnt = uint8(reg[6])
	dtu_statusinfo.Dtu_speed=uint16(reg[7])
	dtu_statusinfo.Dtu_altitude=uint16(reg[8])
	dtu_statusinfo.Dtu_pluginVoltage=uint8(reg[9] >> 8)
	dtu_statusinfo.Dtu_selfInVoltage=uint8(reg[9])
	dtu_statusinfo.Dtu_errStatus=uint8(reg[10] >> 8)
	dtu_statusinfo.Dtu_errNbr=uint8(reg[10])
	dtu_statusinfo.Dtu_errCode=uint16(reg[11])

	var dtu_statusinfolog batterymanage.Dtu_statusInfoLog
	dtu_statusinfolog.Dtu_uptime = dtu_statusinfo.Dtu_uptime
	dtu_statusinfolog.Dtu_id = dtu_statusinfo.Dtu_id
	dtu_statusinfolog.Pkg_id = dtu_statusinfo.Pkg_id
	dtu_statusinfolog.Dtu_latitudeType = dtu_statusinfo.Dtu_latitudeType
	dtu_statusinfolog.Dtu_longitudeType = dtu_statusinfo.Dtu_longitudeType
	dtu_statusinfolog.Dtu_latitude = dtu_statusinfo.Dtu_latitude
	dtu_statusinfolog.Dtu_longitude = dtu_statusinfo.Dtu_longitude
	dtu_statusinfolog.Dtu_csq = dtu_statusinfo.Dtu_csq
	dtu_statusinfolog.Dtu_locateMode = dtu_statusinfo.Dtu_locateMode
	dtu_statusinfolog.Dtu_gpsSateCnt = dtu_statusinfo.Dtu_gpsSateCnt
	dtu_statusinfolog.Dtu_speed = dtu_statusinfo.Dtu_speed
	dtu_statusinfolog.Dtu_altitude = dtu_statusinfo.Dtu_altitude
	dtu_statusinfolog.Dtu_pluginVoltage = dtu_statusinfo.Dtu_pluginVoltage
	dtu_statusinfolog.Dtu_selfInVoltage = dtu_statusinfo.Dtu_selfInVoltage
	dtu_statusinfolog.Dtu_errStatus = dtu_statusinfo.Dtu_errStatus
	dtu_statusinfolog.Dtu_errNbr = dtu_statusinfo.Dtu_errNbr
	dtu_statusinfolog.Dtu_errCode = dtu_statusinfo.Dtu_errCode
	if err:=orm.Eloquent.Create(&dtu_statusinfolog).Error;err!=nil{
		fmt.Println(err)
	}

	//orm.Eloquent.Create(&dtu_statusinfo)
	dtu_statusinfotemp:=dtu_statusinfo
	if err:=orm.Eloquent.Where(&batterymanage.Dtu_statusInfo{Dtu_id: msg.DtuID}).FirstOrCreate(&dtu_statusinfotemp).Error;err != nil {
		fmt.Println(err)
	}else {
		dtu_statusinfomap:=Struct2Map(dtu_statusinfo,[]int{0,-3,-2,-1})
		if err:=orm.Eloquent.Model(batterymanage.Dtu_statusInfo{}).Where(&batterymanage.Dtu_statusInfo{Dtu_id: msg.DtuID}).Updates(dtu_statusinfomap).Error;err != nil {
			fmt.Println(err)
		}
	}
}
func modbusProcess30123(reg []uint16,reglen uint8,msg ModbusMessage)  {
	/*
	pkg_id:= Dtu_Pkg_map[msg.DtuID]
	if len(pkg_id)<5{
		res := Dtu_BMS_map_Init(msg)
		if res != true {
			fmt.Println("find no bmsID")
			pkg_id=""
		}else {
			pkg_id= Dtu_Pkg_map[msg.DtuID]
		}
	}

	 */
	var dtu_statusinfo batterymanage.Dtu_statusInfo
	dtu_statusinfo.Dtu_uptime = time.Unix(msg.Timestamp/1000, 0)
	//dtu_statusinfo.Pkg_id=pkg_id
	dtu_statusinfo.Dtu_id=msg.DtuID
	dtu_statusinfo.Dtu_errStatus=uint8(reg[0] >> 8)
	dtu_statusinfo.Dtu_errNbr=uint8(reg[0])
	dtu_statusinfo.Dtu_errCode=uint16(reg[1])
	//orm.Eloquent.Create(&dtu_statusinfo)
	dtu_statusinfotemp:=dtu_statusinfo
	if err:=orm.Eloquent.Where(&batterymanage.Dtu_statusInfo{Dtu_id: msg.DtuID}).FirstOrCreate(&dtu_statusinfotemp).Error;err != nil {
		fmt.Println(err)
	}else {
		dtu_statusinfomap:=Struct2Map(dtu_statusinfo,[]int{0,2,4,5,6,7,8,9,10,11,12,-3,-2,-1})
		if err:=orm.Eloquent.Model(batterymanage.Dtu_statusInfo{}).Where(&batterymanage.Dtu_statusInfo{Dtu_id: msg.DtuID}).Updates(dtu_statusinfomap).Error;err != nil {
			fmt.Println(err)
		}
	}
}
func modbusProcess30200(reg []uint16,reglen uint8,msg ModbusMessage)  {
	pkg_id:= Dtu_Pkg_map[msg.DtuID]
	if len(pkg_id)<5{
		res := Dtu_BMS_map_Init(msg)
		if res != true {
			fmt.Println("find no bmsID")
			pkg_id=""
		}else {
			pkg_id= Dtu_Pkg_map[msg.DtuID]
		}
	}
	var bms_cellinfo batterymanage.Bms_cellInfo
	bms_cellinfo.Dtu_uptime= time.Unix(msg.Timestamp/1000, 0)
	bms_cellinfo.Pkg_id= pkg_id
	bms_cellinfo.Dtu_id=msg.DtuID
	bms_cellinfo.Bms_cellVoltage1= uint16(reg[0])
	bms_cellinfo.Bms_cellVoltage2= uint16(reg[1])
	bms_cellinfo.Bms_cellVoltage3= uint16(reg[2])
	bms_cellinfo.Bms_cellVoltage4= uint16(reg[3])
	bms_cellinfo.Bms_cellVoltage5= uint16(reg[4])
	bms_cellinfo.Bms_cellVoltage6= uint16(reg[5])
	bms_cellinfo.Bms_cellVoltage7= uint16(reg[6])
	bms_cellinfo.Bms_cellVoltage8= uint16(reg[7])
	bms_cellinfo.Bms_cellVoltage9= uint16(reg[8])
	bms_cellinfo.Bms_cellVoltage10= uint16(reg[9])
	bms_cellinfo.Bms_cellVoltage11= uint16(reg[10])
	bms_cellinfo.Bms_cellVoltage12= uint16(reg[11])
	bms_cellinfo.Bms_cellVoltage13= uint16(reg[12])
	bms_cellinfo.Bms_cellVoltage14= uint16(reg[13])
	bms_cellinfo.Bms_cellVoltage15= uint16(reg[14])
	bms_cellinfo.Bms_cellVoltage16= uint16(reg[15])
	bms_cellinfo.Bms_cellVoltage17= uint16(reg[16])
	bms_cellinfo.Bms_cellVoltage18= uint16(reg[17])
	bms_cellinfo.Bms_cellVoltage19= uint16(reg[18])
	bms_cellinfo.Bms_cellVoltage20= uint16(reg[19])

	var bms_cellinfolog batterymanage.Bms_cellInfoLog
	bms_cellinfolog.Dtu_uptime= bms_cellinfo.Dtu_uptime
	bms_cellinfolog.Pkg_id= bms_cellinfo.Pkg_id
	bms_cellinfolog.Dtu_id=bms_cellinfo.Dtu_id
	bms_cellinfolog.Bms_cellVoltage1= bms_cellinfo.Bms_cellVoltage1
	bms_cellinfolog.Bms_cellVoltage2= bms_cellinfo.Bms_cellVoltage2
	bms_cellinfolog.Bms_cellVoltage3= bms_cellinfo.Bms_cellVoltage3
	bms_cellinfolog.Bms_cellVoltage4= bms_cellinfo.Bms_cellVoltage4
	bms_cellinfolog.Bms_cellVoltage5= bms_cellinfo.Bms_cellVoltage5
	bms_cellinfolog.Bms_cellVoltage6= bms_cellinfo.Bms_cellVoltage6
	bms_cellinfolog.Bms_cellVoltage7= bms_cellinfo.Bms_cellVoltage7
	bms_cellinfolog.Bms_cellVoltage8= bms_cellinfo.Bms_cellVoltage8
	bms_cellinfolog.Bms_cellVoltage9= bms_cellinfo.Bms_cellVoltage9
	bms_cellinfolog.Bms_cellVoltage10= bms_cellinfo.Bms_cellVoltage10
	bms_cellinfolog.Bms_cellVoltage11= bms_cellinfo.Bms_cellVoltage11
	bms_cellinfolog.Bms_cellVoltage12= bms_cellinfo.Bms_cellVoltage12
	bms_cellinfolog.Bms_cellVoltage13= bms_cellinfo.Bms_cellVoltage13
	bms_cellinfolog.Bms_cellVoltage14= bms_cellinfo.Bms_cellVoltage14
	bms_cellinfolog.Bms_cellVoltage15= bms_cellinfo.Bms_cellVoltage15
	bms_cellinfolog.Bms_cellVoltage16= bms_cellinfo.Bms_cellVoltage16
	bms_cellinfolog.Bms_cellVoltage17= bms_cellinfo.Bms_cellVoltage17
	bms_cellinfolog.Bms_cellVoltage18= bms_cellinfo.Bms_cellVoltage18
	bms_cellinfolog.Bms_cellVoltage19= bms_cellinfo.Bms_cellVoltage19
	bms_cellinfolog.Bms_cellVoltage20= bms_cellinfo.Bms_cellVoltage20
	if err:=orm.Eloquent.Create(&bms_cellinfolog).Error;err!=nil{
		fmt.Println(err)
	}
	bms_cellinfotemp:=bms_cellinfo
	if err:=orm.Eloquent.Where(&batterymanage.Bms_cellInfo{Pkg_id: pkg_id}).FirstOrCreate(&bms_cellinfotemp).Error;err != nil {
		fmt.Println(err)
	}else {
		bms_cellinfomap:=Struct2Map(bms_cellinfo,[]int{0,2,4,5,6,7,8,9,10,11,12,-3,-2,-1})
		if err:=orm.Eloquent.Model(batterymanage.Bms_cellInfo{}).Where(&batterymanage.Bms_cellInfo{Pkg_id: pkg_id}).Updates(bms_cellinfomap).Error;err != nil {
			fmt.Println(err)
		}
	}
}
func modbusProcess30300(reg []uint16,reglen uint8,msg ModbusMessage)  {
	pkg_id:= Dtu_Pkg_map[msg.DtuID]
	if len(pkg_id)<5{
		res := Dtu_BMS_map_Init(msg)
		if res != true {
			fmt.Println("find no bmsID")
			pkg_id=""
		}else {
			pkg_id= Dtu_Pkg_map[msg.DtuID]
		}
	}
	var bms_temperatureinfo batterymanage.Bms_temperatureInfo
	bms_temperatureinfo.Dtu_uptime=time.Unix(msg.Timestamp/1000, 0)
	bms_temperatureinfo.Pkg_id= pkg_id
	bms_temperatureinfo.Dtu_id=msg.DtuID
	bms_temperatureinfo.Bms_temperature1= uint8(reg[0]>>8) - 40
	bms_temperatureinfo.Bms_temperature2= uint8(reg[0]) - 40
	bms_temperatureinfo.Bms_temperature3= uint8(reg[1]>>8) - 40
	bms_temperatureinfo.Bms_temperature4= uint8(reg[1]) - 40
	bms_temperatureinfo.Bms_temperature5= uint8(reg[2]>>8) - 40
	bms_temperatureinfo.Bms_temperature6= uint8(reg[2]) - 40

	var bms_temperatureinfolog batterymanage.Bms_temperatureInfoLog
	bms_temperatureinfolog.Dtu_uptime=bms_temperatureinfo.Dtu_uptime
	bms_temperatureinfolog.Pkg_id= pkg_id
	bms_temperatureinfolog.Dtu_id=msg.DtuID
	bms_temperatureinfolog.Bms_temperature1= bms_temperatureinfo.Bms_temperature1
	bms_temperatureinfolog.Bms_temperature2= bms_temperatureinfo.Bms_temperature2
	bms_temperatureinfolog.Bms_temperature3= bms_temperatureinfo.Bms_temperature3
	bms_temperatureinfolog.Bms_temperature4= bms_temperatureinfo.Bms_temperature4
	bms_temperatureinfolog.Bms_temperature5= bms_temperatureinfo.Bms_temperature5
	bms_temperatureinfolog.Bms_temperature6= bms_temperatureinfo.Bms_temperature6
	if err:=orm.Eloquent.Create(&bms_temperatureinfolog).Error;err!=nil{
		fmt.Println(err)
	}
	bms_temperatureinfotemp:=bms_temperatureinfo
	if err:=orm.Eloquent.Where(&batterymanage.Bms_temperatureInfo{Pkg_id: pkg_id}).FirstOrCreate(&bms_temperatureinfotemp).Error;err != nil {
		fmt.Println(err)
	}else {
		bms_temperatureinfomap:=Struct2Map(bms_temperatureinfo,[]int{0,-3,-2,-1})
		if err:=orm.Eloquent.Model(batterymanage.Bms_temperatureInfo{}).Where(&batterymanage.Bms_temperatureInfo{Pkg_id: pkg_id}).Updates(bms_temperatureinfomap).Error;err != nil {
			fmt.Println(err)
		}
	}
}
func modbusProcess30500(reg []uint16,reglen uint8,msg ModbusMessage)  {
	pkg_id:= Dtu_Pkg_map[msg.DtuID]
	if len(pkg_id)<5{
		res := Dtu_BMS_map_Init(msg)
		if res != true {
			fmt.Println("find no bmsID")
			pkg_id=""
		}else {
			pkg_id= Dtu_Pkg_map[msg.DtuID]
		}
	}
	var bms_historyinfo batterymanage.Bms_historyInfo
	bms_historyinfo.Dtu_uptime=time.Unix(msg.Timestamp/1000, 0)
	bms_historyinfo.Pkg_id= pkg_id
	bms_historyinfo.Dtu_id=msg.DtuID
	bms_historyinfo.Pkg_historyMaxCellVoltage= uint16(reg[0])
	bms_historyinfo.Pkg_historyMinCellVoltage= uint16(reg[1])
	bms_historyinfo.Pkg_historyMaxVoltageDif= uint16(reg[2])
	bms_historyinfo.Pkg_historyMaxTemperature= uint8(reg[3]>>8) - 40
	bms_historyinfo.Pkg_historyMinTemperature= uint8(reg[3]) - 40
	bms_historyinfo.Pkg_historyMaxDischargeCurrent= uint16(reg[4])
	bms_historyinfo.Pkg_historyMaxChargeCurrent= uint16(reg[4])
	bms_historyinfo.Pkg_nbrOfChargeDischarge= uint16(reg[4])
	bms_historyinfo.Pkg_nbrofChargingCycle= uint16(reg[4])
	bms_historyinfotemp:=bms_historyinfo
	if err:=orm.Eloquent.Where(&batterymanage.Bms_historyInfo{Pkg_id: pkg_id}).FirstOrCreate(&bms_historyinfotemp).Error;err != nil {
		fmt.Println(err)
	}else {
		bms_historyinfomap:=Struct2Map(bms_historyinfo,[]int{0,-3,-2,-1})
		if err:=orm.Eloquent.Model(batterymanage.Bms_historyInfo{}).Where(&batterymanage.Bms_historyInfo{Pkg_id: pkg_id}).Updates(bms_historyinfomap).Error;err != nil {
			fmt.Println(err)
		}
	}
}
func modbusProcess30600(reg []uint16,reglen uint8,msg ModbusMessage)  {
	pkg_id:= Dtu_Pkg_map[msg.DtuID]
	if len(pkg_id)<5{
		res := Dtu_BMS_map_Init(msg)
		if res != true {
			fmt.Println("find no bmsID")
			pkg_id=""
		}else {
			pkg_id= Dtu_Pkg_map[msg.DtuID]
		}
	}
	var bms_paraSetReg batterymanage.Bms_paraSetReg
	bms_paraSetReg.Dtu_uptime=time.Unix(msg.Timestamp/1000, 0)
	bms_paraSetReg.Pkg_id= pkg_id
	bms_paraSetReg.Dtu_id=msg.DtuID
	bms_paraSetReg.Bms_chargeMosCtr= uint8(reg[0]>>8)
	bms_paraSetReg.Bms_dischargeMosCtr= uint8(reg[0])
	bms_paraSetReg.Bms_chargeHighTempProtect= uint8(reg[1]>>8) - 40
	bms_paraSetReg.Bms_chargeHighTempRelease= uint8(reg[1]) - 40
	bms_paraSetReg.Bms_chargeLowTempProtect= uint8(reg[2]>>8) - 40
	bms_paraSetReg.Bms_chargeLowTempRelease= uint8(reg[2]) - 40
	bms_paraSetReg.Bms_chargeHighTempDelay= uint8(reg[3]>>8)
	bms_paraSetReg.Bms_chargeLowTempDelay= uint8(reg[3])
	bms_paraSetReg.Bms_dischargeHighTempProtect= uint8(reg[4]>>8) - 40
	bms_paraSetReg.Bms_dischargeHighTempRelease= uint8(reg[4]) - 40
	bms_paraSetReg.Bms_dischargeLowTempProtect= uint8(reg[5]>>8) - 40
	bms_paraSetReg.Bms_dischargeLowTempRelease= uint8(reg[5]) - 40
	bms_paraSetReg.Bms_dischargeHighTempDelay= uint8(reg[6]>>8)
	bms_paraSetReg.Bms_dischargeLowTempDelay= uint8(reg[6])
	bms_paraSetReg.Bms_mosHighTempProtect= uint8(reg[7]>>8) - 40
	bms_paraSetReg.Bms_mosHighTempRelease= uint8(reg[7]) - 40
	bms_paraSetReg.Bms_pkgOverVoltageProtect= uint16(reg[11])
	bms_paraSetReg.Bms_pkgOverVoltageRelease= uint16(reg[12])
	bms_paraSetReg.Bms_pkgUnderVoltageProtect= uint16(reg[13])
	bms_paraSetReg.Bms_pkgUnderVoltageRelease= uint16(reg[14])
	bms_paraSetReg.Bms_pkgUnderVoltageDelay= uint8(reg[15]>>8)
	bms_paraSetReg.Bms_pkgOverVoltageDelay= uint8(reg[15])
	bms_paraSetReg.Bms_cellOverVoltageProtect= uint16(reg[16])
	bms_paraSetReg.Bms_cellOverVoltageRelease= uint16(reg[17])
	bms_paraSetReg.Bms_cellUnderVoltageProtect= uint16(reg[18])
	bms_paraSetReg.Bms_cellUnderVoltageRelease= uint16(reg[19])
	bms_paraSetReg.Bms_cellUnderVoltageDelay= uint8(reg[20]>>8)
	bms_paraSetReg.Bms_cellOverVoltageDelay= uint8(reg[20])
	bms_paraSetReg.Bms_chargeOverCurrentProtect= uint16(reg[21])
	bms_paraSetReg.Bms_chargeOverCurrentDelay= uint8(reg[22]>>8)
	bms_paraSetReg.Bms_chargeOverCurrentRelease= uint8(reg[22])
	bms_paraSetReg.Bms_dischargeOverCurrentProtect= uint16(reg[23])
	bms_paraSetReg.Bms_dischargeOverCurrentDelay= uint8(reg[24]>>8)
	bms_paraSetReg.Bms_dischargeOverCurrentRelease= uint8(reg[24])
	bms_paraSetReg.Bms_balanceOpenVoltage= uint16(reg[29])
	bms_paraSetReg.Bms_balanceVoltageDiff= uint16(reg[30])
	bms_paraSetReg.Bms_balanceTime= uint16(reg[31])
	bms_paraSetReg.Bms_funcConfig= uint16(reg[33])
	bms_paraSetReg.Bms_hardCellOverVoltage= uint16(reg[38])
	bms_paraSetReg.Bms_hardCellUnderVoltage= uint16(reg[39])
	bms_paraSetReg.Bms_hardOverCurrentAndDelayTime= uint16(reg[40])
	bms_paraSetReg.Bms_hardUnderVoltageAndOverCurrentDelayTime= uint16(reg[41])
	bms_paraSetReg.Bms_magneticCheckEnable= uint16(reg[42])
	bms_paraSetReg.Bms_forceIntoStorageMode= uint16(reg[43])
	bms_paraSetReg.Bms_enableChargeStatus= uint16(reg[44])
	bms_paraSetRegtemp :=bms_paraSetReg
	if err:=orm.Eloquent.Where(&batterymanage.Bms_paraSetReg{Pkg_id: pkg_id}).FirstOrCreate(&bms_paraSetRegtemp).Error;err != nil {
		fmt.Println(err)
	}else {
		bms_paraSetRegmap:=Struct2Map(bms_paraSetReg,[]int{0,-3,-2,-1})
		if err:=orm.Eloquent.Model(batterymanage.Bms_paraSetReg{}).Where(&batterymanage.Bms_paraSetReg{Pkg_id: pkg_id}).Updates(bms_paraSetRegmap).Error;err != nil {
			fmt.Println(err)
		}
	}
	if reglen == 61 {
		regTemp, _ := sliceUin16Tobyte(reg[51:])
		dtu_otaIP:= string(regTemp)
		var dtu_paraSetReg batterymanage.Dtu_paraSetReg
		dtu_paraSetReg.Dtu_uptime=bms_paraSetReg.Dtu_uptime
		dtu_paraSetReg.Pkg_id=pkg_id
		dtu_paraSetReg.Dtu_id=msg.DtuID
		dtu_paraSetReg.Dtu_pkgInfoReportPeriod= uint16(reg[47])
		dtu_paraSetReg.Dtu_remoteLockCar= uint16(reg[48])
		dtu_paraSetReg.Dtu_voiceTipsOnOff=uint8(reg[49])
		dtu_paraSetReg.Dtu_voiceTipsThresholdValue=uint8(reg[50] >> 8)
		dtu_paraSetReg.Dtu_voiceTipsDownBulk=uint8(reg[50])
		dtu_paraSetReg.Dtu_otaIP=dtu_otaIP
		dtu_paraSetRegtemp:=dtu_paraSetReg
		if err:=orm.Eloquent.Where(&batterymanage.Dtu_paraSetReg{Dtu_id: msg.DtuID}).FirstOrCreate(&dtu_paraSetRegtemp).Error;err != nil {
			fmt.Println(err)
		}else {
			bms_paraSetRegmap:=Struct2Map(dtu_paraSetReg,[]int{0,-3,-2,-1})
			if err:=orm.Eloquent.Model(batterymanage.Dtu_paraSetReg{}).Where(&batterymanage.Dtu_paraSetReg{Dtu_id: msg.DtuID}).Updates(bms_paraSetRegmap).Error;err != nil {
				fmt.Println(err)
			}
		}
	}
}
func modbusProcess30647(reg []uint16,reglen uint8,msg ModbusMessage)  {
	pkg_id:= Dtu_Pkg_map[msg.DtuID]
	if len(pkg_id)<5{
		res := Dtu_BMS_map_Init(msg)
		if res != true {
			fmt.Println("find no bmsID")
			pkg_id=""
		}else {
			pkg_id= Dtu_Pkg_map[msg.DtuID]
		}
	}
	regTemp, _ := sliceUin16Tobyte(reg[4:])
	dtu_otaIP:= string(regTemp)
	var dtu_paraSetReg batterymanage.Dtu_paraSetReg
	dtu_paraSetReg.Dtu_uptime=time.Unix(msg.Timestamp/1000, 0)
	dtu_paraSetReg.Pkg_id= pkg_id
	dtu_paraSetReg.Dtu_id=msg.DtuID
	dtu_paraSetReg.Dtu_pkgInfoReportPeriod=  uint16(reg[0])
	dtu_paraSetReg.Dtu_remoteLockCar=  uint16(reg[1])
	dtu_paraSetReg.Dtu_voiceTipsOnOff= uint8(reg[2])
	dtu_paraSetReg.Dtu_voiceTipsThresholdValue=uint8(reg[3] >> 8)
	dtu_paraSetReg.Dtu_voiceTipsDownBulk=uint8(reg[3])
	dtu_paraSetReg.Dtu_otaIP= dtu_otaIP
	dtu_paraSetRegtemp:=dtu_paraSetReg
	if err:=orm.Eloquent.Where(&batterymanage.Dtu_paraSetReg{Dtu_id: msg.DtuID}).FirstOrCreate(&dtu_paraSetRegtemp).Error;err != nil {
		fmt.Println(err)
	}else {
		bms_paraSetRegmap:=Struct2Map(dtu_paraSetReg,[]int{0,-3,-2,-1})
		if err:=orm.Eloquent.Model(batterymanage.Dtu_paraSetReg{}).Where(&batterymanage.Dtu_paraSetReg{Dtu_id: msg.DtuID}).Updates(bms_paraSetRegmap).Error;err != nil {
			fmt.Println(err)
		}
	}
}
