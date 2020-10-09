package batterymanage

import (
	orm "go-admin/global"
	"go-admin/models"
	"go-admin/tools"
	"time"
)

type Bms_statusInfoLog struct {
	Bms_statusInfoLogId     int    `json:"bms_statusInfoLogId" gorm:"size:10;primary_key;AUTO_INCREMENT"`
	Dtu_uptime time.Time  `json:"dtu_uptime"`
	Pkg_id   string `json:"pkg_id" gorm:"size:20;"`
	Dtu_id      string `json:"dtu_id" gorm:"size:20;"`
	Bms_chargeStatus      uint8 `json:"bms_chargeStatus" gorm:"Type：uint8"`
	Bms_soc   uint8    `json:"bms_soc" gorm:"Type：uint8"`
	Bms_errStatus   uint8    `json:"bms_errStatus" gorm:"Type：uint8"`
	Bms_errNbr   uint8    `json:"bms_errNbr" gorm:"Type：uint8"`
	Bms_errCode   uint32    `json:"bms_errCode" gorm:"Type：uint32"`
	Bms_voltage   uint16    `json:"bms_voltage" gorm:"Type：uint16"`
	Bms_current	  uint16  `json:"bms_current" gorm:"Type：uint16"`
	Bms_maxCellVoltage   uint16    `json:"bms_maxCellVoltage" gorm:"Type：uint16"`
	Bms_minCellVoltage   uint16    `json:"bms_minCellVoltage" gorm:"Type：uint16"`
	Bms_averageCellVoltage   uint16    `json:"bms_averageCellVoltage" gorm:"Type：uint16"`
	Bms_maxTemperature   uint8    `json:"bms_maxTemperature" gorm:"Type：uint8"`
	Bms_minTemperature   uint8    `json:"bms_minTemperature" gorm:"Type：uint8"`
	Bms_mosTemperature   uint8    `json:"bms_mosTemperature" gorm:"Type：uint8"`
	Bms_balanceResistance   uint8    `json:"bms_balanceResistance" gorm:"Type：uint8"`
	Bms_chargeMosStatus   uint8    `json:"bms_chargeMosStatus" gorm:"Type：uint8"`
	Bms_dischargeMosStatus   uint8    `json:"bms_dischargeMosStatus" gorm:"Type：uint8"`
	Bms_otaBufStatus   uint8    `json:"bms_otaBufStatus" gorm:"Type：uint8"`
	Bms_magneticCheck   uint8   `json:"bms_magneticCheck" gorm:"Type：uint8"`
	DataScope  string `json:"dataScope" gorm:"-"`
	UpdateBy  string `gorm:"size:128;" json:"updateBy"`
	models.BaseModel
}
func (Bms_statusInfoLog) TableName() string {
	return "user_bms_statusinfolog"
}


type BatterySOCInfo struct {
	//Bms_statusInfoLog
	Bms_statusInfoLogId     int    `json:"bms_statusInfoLogId" gorm:"size:10;primary_key;AUTO_INCREMENT"`
	Dtu_uptime time.Time  `json:"dtu_uptime"`
	Pkg_id   string `json:"pkg_id" gorm:"size:20;"`
	Dtu_id      string `json:"dtu_id" gorm:"size:20;"`
	Bms_chargeStatus      uint8 `json:"bms_chargeStatus" gorm:"Type：uint8"`
	Bms_soc   uint8    `json:"bms_soc" gorm:"Type：uint8"`
	Bms_errNbr   uint8    `json:"bms_errNbr" gorm:"Type：uint8"`
	Bms_errCode   uint32    `json:"bms_errCode" gorm:"Type：uint32"`
	Bms_voltage   uint16    `json:"bms_voltage" gorm:"Type：uint16"`
	Bms_current	  uint16  `json:"bms_current" gorm:"Type：uint16"`
	//无关数据库
	DataScope  string `json:"dataScope" gorm:"-"`
	models.BaseModel
}
func (BatterySOCInfo) TableName() string {
	return "user_bms_statusinfolog"
}
func (e *BatterySOCInfo) GetBatterySOCInfo(starttime time.Time, endtime time.Time,dateflag int) ([]BatterySOCInfo,int, error) {
	var doc []BatterySOCInfo

	table := orm.Eloquent.Table(e.TableName()).Select([]string{"user_bms_statusinfolog.bms_status_info_log_id",
		"user_bms_statusinfolog.dtu_uptime",
		"user_bms_statusinfolog.pkg_id",
		"user_bms_statusinfolog.dtu_id",
		"user_bms_statusinfolog.bms_charge_status",
		"user_bms_statusinfolog.bms_soc",
		"user_bms_statusinfolog.bms_err_nbr",
		"user_bms_statusinfolog.bms_err_code",
		"user_bms_statusinfolog.bms_voltage",
		"user_bms_statusinfolog.bms_current"})

	// 数据权限控制
	dataPermission := new(models.DataPermission)
	dataPermission.UserId, _ = tools.StringToInt(e.DataScope)
	table, err := dataPermission.GetDataScope(e.TableName(), table)
	if err != nil {
		return nil, 0, err
	}
	var count int
	table = table.Order("user_bms_statusinfolog.bms_status_info_log_id").Find(&doc)
	if table.Error != nil {
		return nil, 0, err
	}
	if e.Bms_statusInfoLogId != 0 {
		table = table.Where("user_bms_statusinfolog.bms_status_info_log_id = ?", e.Bms_statusInfoLogId)
	}

	if e.Pkg_id != "" {
		table = table.Where("user_bms_statusinfolog.pkg_id = ?", e.Pkg_id)
	}
	if e.Dtu_id != "" {
		table = table.Where("user_bms_statusinfolog.dtu_id = ?", e.Dtu_id)
	}
	table_date := table.Where("user_bms_statusinfolog.dtu_uptime BETWEEN ? AND ?",starttime,endtime)
	if err:=table_date.Where("`deleted_at` IS NULL").Find(&doc).Count(&count).Error;err!= nil{
		return nil, 0, err
	}
	if count == 0 && dateflag != 1 {
		if err:=table.Where("`deleted_at` IS NULL").Last(&doc).Count(&count).Error;err!= nil{
			return nil, 0, err
		}
	}
	return doc, count, nil
}
