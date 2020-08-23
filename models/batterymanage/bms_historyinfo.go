package batterymanage
import (
	"go-admin/models"
	"time"
)
type Bms_historyInfo struct {
	Bms_historyInfoId     int    `json:"bms_historyInfoId" gorm:"size:10;primary_key;AUTO_INCREMENT"`
	Dtu_uptime time.Time  `json:"dtu_uptime"`
	Pkg_id   string `json:"pkg_id" gorm:"size:20;primary_key;unique;not null;"`
	Dtu_id      string `json:"dtu_id" gorm:"size:20;"`
	Pkg_historyMaxCellVoltage   uint16    `json:"pkg_historyMaxCellVoltage" gorm:"Type：uint16"`
	Pkg_historyMinCellVoltage   uint16    `json:"pkg_historyMinCellVoltage" gorm:"Type：uint16"`
	Pkg_historyMaxVoltageDif   uint16    `json:"pkg_historyMaxVoltageDif" gorm:"Type：uint16"`
	Pkg_historyMaxTemperature   uint8    `json:"pkg_historyMaxTemperature" gorm:"Type：uint8"`
	Pkg_historyMinTemperature   uint8    `json:"pkg_historyMinTemperature" gorm:"Type：uint8"`
	Pkg_historyMaxDischargeCurrent   uint16    `json:"pkg_historyMaxDischargeCurrent" gorm:"Type：uint16"`
	Pkg_historyMaxChargeCurrent   uint16    `json:"pkg_historyMaxChargeCurrent" gorm:"Type：uint16"`
	Pkg_nbrOfChargeDischarge   uint16    `json:"pkg_nbrOfChargeDischarge" gorm:"Type：uint16"`
	Pkg_nbrofChargingCycle   uint16    `json:"pkg_nbrofChargingCycle" gorm:"Type：uint16"`
	DataScope  string `json:"dataScope" gorm:"-"`
	UpdateBy  string `gorm:"size:128;" json:"updateBy"`
	models.BaseModel
}
func (Bms_historyInfo) TableName() string {
	return "user_bms_historyinfo"
}
