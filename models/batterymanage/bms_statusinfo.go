package batterymanage

import (
"go-admin/models"
"time"
)
type Bms_statusinfo struct {
	Bms_statusinfoId     int    `json:"bms_statusinfoId" gorm:"size:10;primary_key;AUTO_INCREMENT"`
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
	models.BaseModel
}
func (Bms_statusinfo) TableName() string {
	return "user_bms_statusinfo"
}
