package batterymanage

import (
	"errors"
	orm "go-admin/global"
	"go-admin/models"
	"go-admin/tools"
	"time"
)

type Dtu_specInfo struct {
	Dtu_specInfoId     int    `json:"dtu_specInfoId" gorm:"size:10;primary_key;AUTO_INCREMENT"`
	Dtu_uptime time.Time  `json:"dtu_uptime"`
	Dtu_id      string `json:"dtu_id" gorm:"size:20;primary_key;unique;not null;"`
	Pkg_id   string `json:"pkg_id" gorm:"size:20;"`
	Dtu_type   uint8    `json:"dtu_type" gorm:"Type：uint8"`
	Dtu_setupType   uint8    `json:"dtu_setupType" gorm:"Type：uint8"`
	Dtu_coreVer   uint16    `json:"dtu_coreVer" gorm:"Type：uint16"`
	Dtu_hardVer   uint8    `json:"dtu_hardVer" gorm:"Type：uint8"`
	Dtu_softVer   uint8    `json:"dtu_softVer" gorm:"Type：uint8"`
	Dtu_protocolVer   string    `json:"dtu_protocolVer" gorm:"Type：size:10"`
	Dtu_devID      string `json:"dtu_devID" gorm:"size:20;"`
	Dtu_simIccid      string `json:"dtu_simIccid" gorm:"size:20;"`
	Dtu_imei      string `json:"dtu_imei" gorm:"size:20;"`
	Dtu_bmsBindStatus   uint8    `json:"dtu_bmsBindStatus" gorm:"Type：uint8"`

	//aliyun，pkg去查询是否有dtu绑定，有绑定判断是否在线，没有绑定就不在线
	Dtu_aliyunStatus uint8    `json:"dtu_aliyunStatus" gorm:"Type：uint8"`


	DataScope  string `json:"dataScope" gorm:"-"`
	UpdateBy  string `gorm:"size:128;" json:"updateBy"`
	models.BaseModel
}
func (Dtu_specInfo) TableName() string {
	return "user_dtu_specinfo"
}

type DtuListInfo struct {
	//Dtu_specInfo
	Dtu_specInfoId     int    `json:"dtu_specInfoId" gorm:"size:10;primary_key;AUTO_INCREMENT"`
	Dtu_id      string `json:"dtu_id" gorm:"size:20;primary_key;unique;not null;"`
	Pkg_id   string `json:"pkg_id" gorm:"size:20;"`
	Dtu_type   uint8    `json:"dtu_type" gorm:"Type：uint8"`
	Dtu_setupType   uint8    `json:"dtu_setupType" gorm:"Type：uint8"`
	//aliyun，pkg去查询是否有dtu绑定，有绑定判断是否在线，没有绑定就不在线
	Dtu_aliyunStatus uint8    `json:"dtu_aliyunStatus" gorm:"Type：uint8"`

	//Dtu_statusInfo
	Dtu_uptime time.Time  `json:"dtu_uptime"`
	Dtu_csq   uint8    `json:"dtu_csq" gorm:"Type：uint8"`
	Dtu_errNbr   uint8    `json:"dtu_errNbr" gorm:"Type：uint8"`

	DataScope  string `json:"dataScope" gorm:"-"`
	models.BaseModel

}
func (DtuListInfo) TableName() string {
	return "user_dtu_specinfo"
}
func (e *DtuListInfo) Getdtu_listinfo(pageSize int, pageIndex int) ([]DtuListInfo,int, error) {
	var doc []DtuListInfo

	table := orm.Eloquent.Table(e.TableName()).Select([]string{"user_dtu_specinfo.dtu_spec_info_id",
		"user_dtu_specinfo.dtu_id",
		"user_dtu_specinfo.pkg_id",
		"user_dtu_specinfo.dtu_type",
		"user_dtu_specinfo.dtu_setup_type",
		"user_dtu_specinfo.dtu_aliyun_status",

		"user_dtu_statusinfo.dtu_uptime",
		"user_dtu_statusinfo.dtu_csq",
		"user_dtu_statusinfo.dtu_err_nbr"})
	table = table.Joins("LEFT JOIN user_dtu_statusinfo ON user_dtu_specinfo.dtu_id=user_dtu_statusinfo.dtu_id")
	if e.Dtu_specInfoId != 0 {
		table = table.Where("dtu_spec_info_id = ?", e.Dtu_specInfoId)
	}
	if e.Dtu_id != "" {
		table = table.Where("dtu_id = ?", e.Dtu_id)
	}
	if e.Pkg_id != "" {
		table = table.Where("pkg_id = ?", e.Pkg_id)
	}
	if e.Dtu_type != 0 {
		table = table.Where("dtu_type = ?", e.Dtu_type)
	}
	if e.Dtu_setupType != 0 {
		table = table.Where("dtu_setup_type = ?", e.Dtu_setupType)
	}
	if e.Dtu_aliyunStatus != 0 {
		table = table.Where("dtu_aliyun_status = ?", e.Dtu_aliyunStatus)
	}
	if e.Dtu_csq != 0 {
		table = table.Where("dtu_csq > ?", e.Dtu_csq)
	}
	if e.Dtu_errNbr != 0 {
		table = table.Where("dtu_err_nbr >= ?", e.Dtu_errNbr)
	}

	// 数据权限控制
	dataPermission := new(models.DataPermission)
	dataPermission.UserId, _ = tools.StringToInt(e.DataScope)
	table, err := dataPermission.GetDataScope(e.TableName(), table)
	if err != nil {
		return nil, 0, err
	}
	var count int
	if err := table.Order("dtu_spec_info_id").Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
		return nil, 0, err
	}
	table.Where("`deleted_at` IS NULL").Count(&count)
	return doc, count, nil
}



type DtuDetailInfo struct {
	//Dtu_specInfo
	Dtu_specInfoId     int    `json:"dtu_specInfoId" gorm:"size:10;primary_key;AUTO_INCREMENT"`
	Dtu_id      string `json:"dtu_id" gorm:"size:20;primary_key;unique;not null;"`
	Pkg_id   string `json:"pkg_id" gorm:"size:20;"`
	Dtu_type   uint8    `json:"dtu_type" gorm:"Type：uint8"`
	Dtu_setupType   uint8    `json:"dtu_setupType" gorm:"Type：uint8"`
	Dtu_coreVer   uint16    `json:"dtu_coreVer" gorm:"Type：uint16"`
	Dtu_hardVer   uint8    `json:"dtu_hardVer" gorm:"Type：uint8"`
	Dtu_softVer   uint8    `json:"dtu_softVer" gorm:"Type：uint8"`
	Dtu_protocolVer   string    `json:"dtu_protocolVer" gorm:"Type：size:10"`
	Dtu_devID      string `json:"dtu_devID" gorm:"size:20;"`
	Dtu_simIccid      string `json:"dtu_simIccid" gorm:"size:20;"`
	Dtu_imei      string `json:"dtu_imei" gorm:"size:20;"`
	Dtu_bmsBindStatus   uint8    `json:"dtu_bmsBindStatus" gorm:"Type：uint8"`
	//aliyun，pkg去查询是否有dtu绑定，有绑定判断是否在线，没有绑定就不在线
	Dtu_aliyunStatus uint8    `json:"dtu_aliyunStatus" gorm:"Type：uint8"`

	//Dtu_statusInfo
	Dtu_uptime time.Time  `json:"dtu_uptime"`
	Dtu_csq   uint8    `json:"dtu_csq" gorm:"Type：uint8"`
	Dtu_locateMode   uint8    `json:"dtu_locateMode" gorm:"Type：uint8"`
	Dtu_gpsSateCnt   uint8    `json:"dtu_gpsSateCnt" gorm:"Type：uint8"`
	Dtu_speed   uint16    `json:"dtu_speed" gorm:"Type：uint16"`
	Dtu_altitude   uint16    `json:"dtu_altitude" gorm:"Type：uint16"`
	Dtu_pluginVoltage   uint8    `json:"dtu_pluginVoltage" gorm:"Type：uint8"`
	Dtu_selfInVoltage   uint8    `json:"dtu_selfInVoltage" gorm:"Type：uint8"`
	Dtu_errNbr   uint8    `json:"dtu_errNbr" gorm:"Type：uint8"`
	Dtu_errCode   uint16    `json:"dtu_errCode" gorm:"Type：uint16"`

	//Bms_specInfo
	Pkg_count   uint8    `json:"pkg_count" gorm:"Type：uint8"`
	Pkg_type   uint8    `json:"pkg_type" gorm:"Type：uint8"`
	Pkg_capacity   uint16    `json:"pkg_capacity" gorm:"Type：uint16"`
	Pkg_nominalVoltage   uint16    `json:"pkg_nominalVoltage" gorm:"Type：uint16"`

	//Bms_statusInfo
	Bms_chargeStatus      uint8 `json:"bms_chargeStatus" gorm:"Type：uint8"`
	Bms_soc   uint8    `json:"bms_soc" gorm:"Type：uint8"`
	Bms_errNbr   uint8    `json:"bms_errNbr" gorm:"Type：uint8"`
	Bms_errCode   uint32    `json:"bms_errCode" gorm:"Type：uint32"`
	Bms_voltage   uint16    `json:"bms_voltage" gorm:"Type：uint16"`
	Bms_current	  uint16  `json:"bms_current" gorm:"Type：uint16"`

	//Dtu_paraSetReg
	Dtu_pkgInfoReportPeriod   uint16    `json:"dtu_pkgInfoReportPeriod" gorm:"Type：uint16"`
	Dtu_remoteLockCar   uint16    `json:"dtu_remoteLockCar" gorm:"Type：uint16"`
	Dtu_voiceTipsOnOff   uint8    `json:"dtu_voiceTipsOnOff" gorm:"Type：uint8"`
	Dtu_voiceTipsThresholdValue   uint8    `json:"dtu_voiceTipsThresholdValue" gorm:"Type：uint8"`
	Dtu_voiceTipsDownBulk   uint8    `json:"dtu_voiceTipsDownBulk" gorm:"Type：uint8"`
	Dtu_otaIP      string `json:"dtu_otaIP" gorm:"size:20;"`

	DataScope  string `json:"dataScope" gorm:"-"`
	models.BaseModel

}
func (DtuDetailInfo) TableName() string {
	return "user_dtu_specinfo"
}

//电池列表
func (e *DtuDetailInfo) GetDtuDetailInfo() ([]DtuDetailInfo,int, error) {
	if e.Dtu_id == "" {
		return nil, 0, errors.New("no dtuid")
	}
	var doc []DtuDetailInfo
	table := orm.Eloquent.Table(e.TableName())
	var dtu_pkg_bind Dtu_specInfo
	if err:= orm.Eloquent.Table("user_dtu_specinfo").Where("user_dtu_specinfo.dtu_id = ?", e.Dtu_id).First(&dtu_pkg_bind).Error;err!=nil{
		return nil, 0, errors.New("no dtuid")
	}else {
		if dtu_pkg_bind.Pkg_id != "" {
			table = table.Select([]string{"user_dtu_specinfo.dtu_spec_info_id",
				"user_dtu_specinfo.dtu_id",
				"user_dtu_specinfo.pkg_id",
				"user_dtu_specinfo.dtu_type",
				"user_dtu_specinfo.dtu_setup_type",
				"user_dtu_specinfo.dtu_core_ver",
				"user_dtu_specinfo.dtu_hard_ver",
				"user_dtu_specinfo.dtu_soft_ver",
				"user_dtu_specinfo.dtu_protocol_ver",
				"user_dtu_specinfo.dtu_dev_id",
				"user_dtu_specinfo.dtu_sim_iccid",
				"user_dtu_specinfo.dtu_imei",
				"user_dtu_specinfo.dtu_bms_bind_status",
				"user_dtu_specinfo.dtu_aliyun_status",

				"user_dtu_statusinfo.dtu_uptime",
				"user_dtu_statusinfo.dtu_csq",
				"user_dtu_statusinfo.dtu_locate_mode",
				"user_dtu_statusinfo.dtu_gps_sate_cnt",
				"user_dtu_statusinfo.dtu_speed",
				"user_dtu_statusinfo.dtu_altitude",
				"user_dtu_statusinfo.dtu_plugin_voltage",
				"user_dtu_statusinfo.dtu_self_in_voltage",
				"user_dtu_statusinfo.dtu_err_nbr",
				"user_dtu_statusinfo.dtu_err_code",

				"user_dtu_paraSetReg.dtu_pkg_info_report_period",
				"user_dtu_paraSetReg.dtu_remote_lock_car",
				"user_dtu_paraSetReg.dtu_voice_tips_on_off",
				"user_dtu_paraSetReg.dtu_voice_tips_threshold_value",
				"user_dtu_paraSetReg.dtu_voice_tips_down_bulk",
				"user_dtu_paraSetReg.dtu_ota_iP",

				"user_bms_statusinfo.bms_charge_status",
				"user_bms_statusinfo.bms_soc",
				"user_bms_statusinfo.bms_err_nbr",
				"user_bms_statusinfo.bms_err_code",
				"user_bms_statusinfo.bms_voltage",
				"user_bms_statusinfo.bms_current"})
			table = table.Joins("LEFT JOIN user_dtu_statusinfo ON user_dtu_specinfo.dtu_id=user_dtu_statusinfo.dtu_id").
				Joins("LEFT JOIN user_dtu_parasetreg ON user_dtu_specinfo.dtu_id=user_dtu_parasetreg.dtu_id").
				Joins("LEFT JOIN user_bms_statusinfo ON user_dtu_specinfo.pkg_id=user_bms_statusinfo.pkg_id")
		}else {
			table = table.Select([]string{"user_dtu_specinfo.dtu_spec_info_id",
				"user_dtu_specinfo.dtu_id",
				"user_dtu_specinfo.pkg_id",
				"user_dtu_specinfo.dtu_type",
				"user_dtu_specinfo.dtu_setup_type",
				"user_dtu_specinfo.dtu_core_ver",
				"user_dtu_specinfo.dtu_hard_ver",
				"user_dtu_specinfo.dtu_soft_ver",
				"user_dtu_specinfo.dtu_protocol_ver",
				"user_dtu_specinfo.dtu_dev_id",
				"user_dtu_specinfo.dtu_sim_iccid",
				"user_dtu_specinfo.dtu_imei",
				"user_dtu_specinfo.dtu_bms_bind_status",
				"user_dtu_specinfo.dtu_aliyun_status",

				"user_dtu_statusinfo.dtu_uptime",
				"user_dtu_statusinfo.dtu_csq",
				"user_dtu_statusinfo.dtu_locate_mode",
				"user_dtu_statusinfo.dtu_gps_sate_cnt",
				"user_dtu_statusinfo.dtu_speed",
				"user_dtu_statusinfo.dtu_altitude",
				"user_dtu_statusinfo.dtu_plugin_voltage",
				"user_dtu_statusinfo.dtu_self_in_voltage",
				"user_dtu_statusinfo.dtu_err_nbr",
				"user_dtu_statusinfo.dtu_err_code",

				"user_dtu_paraSetReg.dtu_pkg_info_report_period",
				"user_dtu_paraSetReg.dtu_remote_lock_car",
				"user_dtu_paraSetReg.dtu_voice_tips_on_off",
				"user_dtu_paraSetReg.dtu_voice_tips_threshold_value",
				"user_dtu_paraSetReg.dtu_voice_tips_down_bulk",
				"user_dtu_paraSetReg.dtu_ota_iP",})
			table = table.Joins("LEFT JOIN user_dtu_statusinfo ON user_dtu_specinfo.dtu_id=user_dtu_statusinfo.dtu_id").
				Joins("LEFT JOIN user_dtu_parasetreg ON user_dtu_specinfo.dtu_id=user_dtu_parasetreg.dtu_id")
		}
	}

	// 数据权限控制
	dataPermission := new(models.DataPermission)
	dataPermission.UserId, _ = tools.StringToInt(e.DataScope)
	table, err := dataPermission.GetDataScope(e.TableName(), table)
	if err != nil {
		return nil, 0, err
	}
	var count int
	table = table.Find(&doc)
	if table.Error!= nil {
		return nil, 0, err
	}
	if e.Dtu_specInfoId != 0 {
		table = table.Where("dtu_spec_info_id = ?", e.Dtu_specInfoId)
	}

	table = table.Where("user_dtu_specinfo.dtu_id = ?", e.Dtu_id).Where("user_dtu_specinfo.deleted_at IS NULL")
	if err := table.First(&doc).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	return doc, count, nil
}
