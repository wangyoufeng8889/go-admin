package batterymanage

import (
	orm "go-admin/global"
	"go-admin/models"
	"go-admin/tools"
	"time"
)
type Bms_statusInfo struct {
	Bms_statusInfoId     int    `json:"bms_statusInfoId" gorm:"size:10;primary_key;AUTO_INCREMENT"`
	Dtu_uptime time.Time  `json:"dtu_uptime"`
	Pkg_id   string `json:"pkg_id" gorm:"size:20;primary_key;unique;not null;"`
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
func (Bms_statusInfo) TableName() string {
	return "user_bms_statusinfo"
}
func (e *Bms_statusInfo) GetBms_statusinfo(startdate time.Time, enddate time.Time,is_oneList string) ([]Bms_statusInfo,int, error) {
	var doc []Bms_statusInfo

	table := orm.Eloquent.Select("*").Table(e.TableName())
	if e.Bms_statusInfoId != 0 {
		table = table.Where("bms_status_info_id = ?", e.Bms_statusInfoId)
	}
	if e.Pkg_id != "" {
		table = table.Where("pkg_id = ?", e.Pkg_id)
	}else {
		table = table.Not("pkg_id = ?", "0")
	}
	if e.Dtu_id != "" {
		table = table.Where("dtu_id = ?", e.Dtu_id)
	}
	table = table.Where("dtu_uptime BETWEEN ? AND ?",startdate,enddate)
	// 数据权限控制
	dataPermission := new(models.DataPermission)
	dataPermission.UserId, _ = tools.StringToInt(e.DataScope)
	table, err := dataPermission.GetDataScope(e.TableName(), table)
	if err != nil {
		return nil, 0, err
	}
	var count int
	if is_oneList == "YES" {
		if err := table.Order("dtu_uptime desc").First(&doc).Error; err != nil {
			return nil, 0, err
		}
		table.Where("`deleted_at` IS NULL").Count(&count)
	}else{
		if err := table.Order("bms_status_info_id").Find(&doc).Error; err != nil {
			return nil, 0, err
		}
		table.Where("`deleted_at` IS NULL").Count(&count)
	}
	return doc, count, nil
}