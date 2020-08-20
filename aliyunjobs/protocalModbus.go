package aliyunjobs

import (
	"fmt"
	orm "go-admin/global"
	"go-admin/models/batterymanage"
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
var Dtu_BMS_map map[string]string

func init()  {
	Dtu_BMS_map = make(map[string]string)
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
func aliyunOnOffprocess(msg ModbusMessage)  {
	var battery_list batterymanage.Battery_list
	battery_list.Dtu_uptime=time.Unix(msg.Timestamp/1000, 0)
	battery_list.Dtu_id=msg.DtuID
	battery_list.Pkg_onOffLineStatus = uint8(msg.Payload[0])
	var temp1 batterymanage.Battery_list
	if err:=orm.Eloquent.Where(&batterymanage.Battery_list{Dtu_id: msg.DtuID}).First(&temp1).Error;err != nil {
		//orm.Eloquent.Create(&battery_list)
	}else {
		orm.Eloquent.Model(&battery_list).Where(&batterymanage.Battery_list{Dtu_id: msg.DtuID}).Update(&battery_list)
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
	sl_in_len := len(sl_in)
	for i:=0;i<sl_in_len;i++ {
		sl_out = append(sl_out,byte(sl_in[i]>>8),byte(sl_in[i]))
	}
	return sl_out,nil
}
func Dtu_BMS_map_Init(msg ModbusMessage)(bool)  {
	var dtu_specinfo batterymanage.Dtu_specinfo

	orm.Eloquent.Where(&batterymanage.Dtu_specinfo{Dtu_id: msg.DtuID}).First(&dtu_specinfo)
	if len(dtu_specinfo.Pkg_id)>0 {
		Dtu_BMS_map[msg.DtuID] = dtu_specinfo.Pkg_id
		return false
	}else{
		Dtu_BMS_map[msg.DtuID] = "0"
		return false
	}
}
func modbusProcess30000(reg []uint16,reglen uint8,msg ModbusMessage)  {
	regTemp, _ := sliceUin16Tobyte(reg[0:10])
	pkg_id:= string(regTemp)
	regTemp, _ = sliceUin16Tobyte(reg[10:20])
	bms_id:= string(regTemp)

	var bms_specinfo batterymanage.Bms_specinfo
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
	data1 := reg[26]/100
	data2 := reg[26]%100
	bms_specinfo.Bms_protocolVer =  fmt.Sprintf("%d.%02d",data1,data2)
	var temp batterymanage.Bms_specinfo
	if err:=orm.Eloquent.Where(&batterymanage.Bms_specinfo{Pkg_id: pkg_id}).First(&temp).Error;err != nil {
		orm.Eloquent.Create(&bms_specinfo)
	}else {
		orm.Eloquent.Model(&bms_specinfo).Where(&batterymanage.Bms_specinfo{Pkg_id: pkg_id}).First(&temp).Update(&bms_specinfo)
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
			Dtu_BMS_map[dtu_id]=pkg_id
		}

		var dtu_specinfo batterymanage.Dtu_specinfo
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
		var temp batterymanage.Dtu_specinfo
		if err:=orm.Eloquent.Where(&batterymanage.Dtu_specinfo{Dtu_id: dtu_id}).First(&temp).Error;err != nil {
			orm.Eloquent.Create(&dtu_specinfo)
		}else {
			orm.Eloquent.Model(&dtu_specinfo).Where(&batterymanage.Dtu_specinfo{Dtu_id: dtu_id}).First(&temp).Update(&dtu_specinfo)
		}
		var battery_list batterymanage.Battery_list
		battery_list.Dtu_uptime=bms_specinfo.Dtu_uptime
		battery_list.Pkg_id=pkg_id
		battery_list.Dtu_id=msg.DtuID
		battery_list.Pkg_count=uint8(reg[20]>>8)
		battery_list.Pkg_type=uint8(reg[20])
		battery_list.Pkg_capacity=uint16(reg[21])
		battery_list.Pkg_nominalVoltage=uint16(reg[22])
		battery_list.Dtu_type = dtu_specinfo.Dtu_type
		battery_list.Dtu_setupType = dtu_specinfo.Dtu_setupType
		var temp1 batterymanage.Battery_list
		if err:=orm.Eloquent.Where(&batterymanage.Battery_list{Pkg_id: pkg_id}).First(&temp1).Error;err != nil {
			orm.Eloquent.Create(&battery_list)
		}else {
			orm.Eloquent.Model(&battery_list).Where(&batterymanage.Battery_list{Pkg_id: pkg_id}).First(&temp1).Update(&battery_list)
		}
	}else{
		Dtu_BMS_map[msg.DtuID]=pkg_id
		var dtu_specinfo batterymanage.Dtu_specinfo
		dtu_specinfo.Dtu_uptime=bms_specinfo.Dtu_uptime
		dtu_specinfo.Dtu_id= msg.DtuID
		dtu_specinfo.Pkg_id= pkg_id
		var temp batterymanage.Dtu_specinfo
		if err:=orm.Eloquent.Where(&batterymanage.Dtu_specinfo{Dtu_id: msg.DtuID}).First(&temp).Error;err != nil {
			orm.Eloquent.Create(&dtu_specinfo)
		}else {
			orm.Eloquent.Model(&dtu_specinfo).Where(&batterymanage.Dtu_specinfo{Dtu_id: msg.DtuID}).First(&temp).Update(&dtu_specinfo)
		}
		var battery_list batterymanage.Battery_list
		battery_list.Dtu_uptime=bms_specinfo.Dtu_uptime
		battery_list.Pkg_id=pkg_id
		battery_list.Dtu_id=msg.DtuID
		battery_list.Pkg_count=uint8(reg[20]>>8)
		battery_list.Pkg_type=uint8(reg[20])
		battery_list.Pkg_capacity=uint16(reg[21])
		battery_list.Pkg_nominalVoltage=uint16(reg[22])
		var temp1 batterymanage.Battery_list
		if err:=orm.Eloquent.Where(&batterymanage.Battery_list{Pkg_id: pkg_id}).First(&temp1).Error;err != nil {
			orm.Eloquent.Create(&battery_list)
		}else {
			orm.Eloquent.Model(&battery_list).Where(&batterymanage.Battery_list{Pkg_id: pkg_id}).First(&temp1).Update(&battery_list)
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
		fmt.Println("dtu id is not aliyun id")
	}else {
		Dtu_BMS_map[dtu_id]="0"
	}
	var dtu_specinfo batterymanage.Dtu_specinfo
	dtu_specinfo.Dtu_uptime=time.Unix(msg.Timestamp/1000, 0)
	dtu_specinfo.Dtu_id= dtu_id
	dtu_specinfo.Pkg_id= "0"
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
	var temp batterymanage.Dtu_specinfo
	if err:=orm.Eloquent.Where(&batterymanage.Dtu_specinfo{Dtu_id: msg.DtuID}).First(&temp).Error;err != nil {
		orm.Eloquent.Create(&dtu_specinfo)
	}else {
		orm.Eloquent.Model(&dtu_specinfo).Where(&batterymanage.Dtu_specinfo{Dtu_id: msg.DtuID}).First(&temp).Update(&dtu_specinfo)
	}

}
func modbusProcess30100(reg []uint16,reglen uint8,msg ModbusMessage)  {
	pkg_id:= Dtu_BMS_map[msg.DtuID]
	if len(pkg_id)<5{
		res := Dtu_BMS_map_Init(msg)
		if res != true {
			pkg_id="0"
		}else {
			pkg_id=Dtu_BMS_map[msg.DtuID]
		}
	}
	var bms_statusinfo batterymanage.Bms_statusinfo
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
	orm.Eloquent.Create(&bms_statusinfo)
	if reglen == 25 {
		var dtu_statusinfo batterymanage.Dtu_statusinfo
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
		var data1,data2 int
		data1 = Dtu_latitude/1000000
		data2 = Dtu_latitude%1000000
		dtu_statusinfo.Dtu_latitude =  fmt.Sprint(data1,".",data2)
		data1 = Dtu_longitude/1000000
		data2 = Dtu_longitude%1000000
		dtu_statusinfo.Dtu_longitude = fmt.Sprint(data1,".",data2)

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
		orm.Eloquent.Create(&dtu_statusinfo)
		var battery_list batterymanage.Battery_list
		battery_list.Dtu_uptime=bms_statusinfo.Dtu_uptime
		battery_list.Pkg_id=pkg_id
		battery_list.Dtu_id=msg.DtuID
		battery_list.Bms_soc=bms_statusinfo.Bms_soc
		battery_list.Bms_chargeStatus=bms_statusinfo.Bms_chargeStatus
		battery_list.Dtu_latitudeType = dtu_statusinfo.Dtu_latitudeType
		battery_list.Dtu_longitudeType = dtu_statusinfo.Dtu_longitudeType
		battery_list.Dtu_latitude =  dtu_statusinfo.Dtu_latitude
		battery_list.Dtu_longitude = dtu_statusinfo.Dtu_longitude
		battery_list.Dtu_csq = dtu_statusinfo.Dtu_csq
		battery_list.Dtu_locateMode = dtu_statusinfo.Dtu_locateMode
		var temp1 batterymanage.Battery_list
		if err:=orm.Eloquent.Where(&batterymanage.Battery_list{Pkg_id: pkg_id}).First(&temp1).Error;err != nil {
			//orm.Eloquent.Create(&battery_list)
		}else {
			orm.Eloquent.Model(&battery_list).Where(&batterymanage.Battery_list{Pkg_id: pkg_id}).First(&temp1).Update(&battery_list)
		}
	}else {
		var battery_list batterymanage.Battery_list
		battery_list.Dtu_uptime=bms_statusinfo.Dtu_uptime
		battery_list.Pkg_id=pkg_id
		battery_list.Dtu_id=msg.DtuID
		battery_list.Bms_soc=bms_statusinfo.Bms_soc
		battery_list.Bms_chargeStatus=bms_statusinfo.Bms_chargeStatus
		var temp1 batterymanage.Battery_list
		if err:=orm.Eloquent.Where(&batterymanage.Battery_list{Pkg_id: pkg_id}).First(&temp1).Error;err != nil {
			//orm.Eloquent.Create(&battery_list)
		}else {
			orm.Eloquent.Model(&battery_list).Where(&batterymanage.Battery_list{Pkg_id: pkg_id}).First(&temp1).Update(&battery_list)
		}
	}
}
func modbusProcess30113(reg []uint16,reglen uint8,msg ModbusMessage)  {
	pkg_id:= Dtu_BMS_map[msg.DtuID]
	if pkg_id == ""{
		res:= Dtu_BMS_map_Init(msg)
		if res != true {
			fmt.Println("find no bmsID")
			pkg_id="0"
		}else {
			pkg_id= Dtu_BMS_map[msg.DtuID]
		}
	}
	var dtu_statusinfo batterymanage.Dtu_statusinfo
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
	var data1,data2 int
	data1 = Dtu_latitude/1000000
	data2 = Dtu_latitude%1000000
	dtu_statusinfo.Dtu_latitude =  fmt.Sprint(data1,".",data2)
	data1 = Dtu_longitude/1000000
	data2 = Dtu_longitude%1000000
	dtu_statusinfo.Dtu_longitude = fmt.Sprint(data1,".",data2)
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
	orm.Eloquent.Create(&dtu_statusinfo)
}
func modbusProcess30123(reg []uint16,reglen uint8,msg ModbusMessage)  {
	pkg_id:= Dtu_BMS_map[msg.DtuID]
	if pkg_id == ""{
		res := Dtu_BMS_map_Init(msg)
		if res != true {
			fmt.Println("find no bmsID")
			pkg_id="0"
		}else {
			pkg_id= Dtu_BMS_map[msg.DtuID]
		}
	}
	var dtu_statusinfo batterymanage.Dtu_statusinfo
	dtu_statusinfo.Dtu_uptime = time.Unix(msg.Timestamp/1000, 0)
	dtu_statusinfo.Pkg_id=pkg_id
	dtu_statusinfo.Dtu_id=msg.DtuID
	dtu_statusinfo.Dtu_errStatus=uint8(reg[0] >> 8)
	dtu_statusinfo.Dtu_errNbr=uint8(reg[0])
	dtu_statusinfo.Dtu_errCode=uint16(reg[1])
	orm.Eloquent.Create(&dtu_statusinfo)
}
func modbusProcess30200(reg []uint16,reglen uint8,msg ModbusMessage)  {
	pkg_id:= Dtu_BMS_map[msg.DtuID]
	if pkg_id == ""{
		res := Dtu_BMS_map_Init(msg)
		if res != true {
			fmt.Println("find no bmsID")
			pkg_id="0"
		}else {
			pkg_id= Dtu_BMS_map[msg.DtuID]
		}
	}
	var bms_cellinfo batterymanage.Bms_cellinfo
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
	orm.Eloquent.Create(&bms_cellinfo)
}
func modbusProcess30300(reg []uint16,reglen uint8,msg ModbusMessage)  {
	pkg_id:= Dtu_BMS_map[msg.DtuID]
	if pkg_id == ""{
		res := Dtu_BMS_map_Init(msg)
		if res != true {
			fmt.Println("find no bmsID")
			pkg_id="0"
		}else {
			pkg_id= Dtu_BMS_map[msg.DtuID]
		}
	}
	var bms_temperatureinfo batterymanage.Bms_temperatureinfo
	bms_temperatureinfo.Dtu_uptime=time.Unix(msg.Timestamp/1000, 0)
	bms_temperatureinfo.Pkg_id= pkg_id
	bms_temperatureinfo.Dtu_id=msg.DtuID
	bms_temperatureinfo.Bms_temperature1= uint8(reg[0]>>8) - 40
	bms_temperatureinfo.Bms_temperature2= uint8(reg[0]) - 40
	bms_temperatureinfo.Bms_temperature3= uint8(reg[1]>>8) - 40
	bms_temperatureinfo.Bms_temperature4= uint8(reg[1]) - 40
	bms_temperatureinfo.Bms_temperature5= uint8(reg[2]>>8) - 40
	bms_temperatureinfo.Bms_temperature6= uint8(reg[2]) - 40
	orm.Eloquent.Create(&bms_temperatureinfo)
}
func modbusProcess30500(reg []uint16,reglen uint8,msg ModbusMessage)  {
	pkg_id:= Dtu_BMS_map[msg.DtuID]
	if pkg_id == ""{
		res := Dtu_BMS_map_Init(msg)
		if res != true {
			fmt.Println("find no bmsID")
			pkg_id="0"
		}else {
			pkg_id= Dtu_BMS_map[msg.DtuID]
		}
	}
	var bms_historyinfo batterymanage.Bms_historyinfo
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
	var temp batterymanage.Bms_historyinfo
	if err:=orm.Eloquent.Where(&batterymanage.Bms_historyinfo{Pkg_id: pkg_id}).First(&temp).Error;err != nil {
		orm.Eloquent.Create(&bms_historyinfo)
	}else {
		orm.Eloquent.Model(&bms_historyinfo).Where(&batterymanage.Bms_historyinfo{Pkg_id: pkg_id}).First(&temp).Update(&bms_historyinfo)
	}
}
func modbusProcess30600(reg []uint16,reglen uint8,msg ModbusMessage)  {
	pkg_id:= Dtu_BMS_map[msg.DtuID]
	if pkg_id == ""{
		res := Dtu_BMS_map_Init(msg)
		if res != true {
			fmt.Println("find no bmsID")
			pkg_id="0"
		}else {
			pkg_id= Dtu_BMS_map[msg.DtuID]
		}
	}
	var bms_paraSetReg batterymanage.Bms_paraSetReg
	bms_paraSetReg.Dtu_uptime=time.Unix(msg.Timestamp/1000, 0)
	bms_paraSetReg.Pkg_id= pkg_id
	bms_paraSetReg.Dtu_id=msg.DtuID
	bms_paraSetReg.Bms_chargeMosCtr= uint8(reg[0]>>8)
	bms_paraSetReg.Bms_dischargeMosCtr= uint8(reg[0])
	bms_paraSetReg.Bms_chargeHighTempProtect= uint8(reg[1]>>8)
	bms_paraSetReg.Bms_chargeHighTempRelease= uint8(reg[1])
	bms_paraSetReg.Bms_chargeLowTempProtect= uint8(reg[2]>>8)
	bms_paraSetReg.Bms_chargeLowTempRelease= uint8(reg[2])
	bms_paraSetReg.Bms_chargeHighTempDelay= uint8(reg[3]>>8)
	bms_paraSetReg.Bms_chargeLowTempDelay= uint8(reg[3])
	bms_paraSetReg.Bms_dischargeHighTempProtect= uint8(reg[4]>>8)
	bms_paraSetReg.Bms_dischargeHighTempRelease= uint8(reg[4])
	bms_paraSetReg.Bms_dischargeLowTempProtect= uint8(reg[5]>>8)
	bms_paraSetReg.Bms_dischargeLowTempRelease= uint8(reg[5])
	bms_paraSetReg.Bms_dischargeHighTempDelay= uint8(reg[6]>>8)
	bms_paraSetReg.Bms_dischargeLowTempDelay= uint8(reg[6])
	bms_paraSetReg.Bms_mosHighTempProtect= uint8(reg[7]>>8)
	bms_paraSetReg.Bms_mosHighTempRelease= uint8(reg[7])
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
	var temp batterymanage.Bms_paraSetReg
	if err:=orm.Eloquent.Where(&batterymanage.Bms_paraSetReg{Pkg_id: pkg_id}).First(&temp).Error;err != nil {
		orm.Eloquent.Create(&bms_paraSetReg)
	}else {
		orm.Eloquent.Model(&bms_paraSetReg).Where(&batterymanage.Bms_paraSetReg{Pkg_id: pkg_id}).First(&temp).Update(&bms_paraSetReg)
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
		var temp batterymanage.Dtu_paraSetReg
		if err:=orm.Eloquent.Where(&batterymanage.Dtu_paraSetReg{Dtu_id: msg.DtuID}).First(&temp).Error;err != nil {
			orm.Eloquent.Create(&dtu_paraSetReg)
		}else {
			orm.Eloquent.Model(&dtu_paraSetReg).Where(&batterymanage.Dtu_paraSetReg{Dtu_id: msg.DtuID}).First(&temp).Update(&dtu_paraSetReg)
		}
	}
}
func modbusProcess30647(reg []uint16,reglen uint8,msg ModbusMessage)  {
	pkg_id:= Dtu_BMS_map[msg.DtuID]
	if pkg_id == ""{
		res := Dtu_BMS_map_Init(msg)
		if res != true {
			fmt.Println("find no bmsID")
			pkg_id="0"
		}else {
			pkg_id= Dtu_BMS_map[msg.DtuID]
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
	var temp batterymanage.Dtu_paraSetReg
	if err:=orm.Eloquent.Where(&batterymanage.Dtu_paraSetReg{Dtu_id: msg.DtuID}).First(&temp).Error;err != nil {
		orm.Eloquent.Create(&dtu_paraSetReg)
	}else {
		orm.Eloquent.Model(&dtu_paraSetReg).Where(&batterymanage.Dtu_paraSetReg{Dtu_id: msg.DtuID}).First(&temp).Update(&dtu_paraSetReg)
	}
}
