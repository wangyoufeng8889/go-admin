package batterymanage

import (
	"go-admin/models"
	"time"
)
type Dtu_specinfo struct {
	Dtu_specinfoId     int    `json:"dtu_specinfoId" gorm:"size:10;primary_key;AUTO_INCREMENT"`
	Dtu_uptime time.Time  `json:"dtu_uptime"`
	Dtu_id      string `json:"dtu_id" gorm:"size:20;"`
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
	models.BaseModel
}
func (Dtu_specinfo) TableName() string {
	return "user_dtu_specinfo"
}