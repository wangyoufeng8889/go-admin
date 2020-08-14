package batterymanage

import (
	"go-admin/models"
	"time"
)
type Dtu_statusinfo struct {
	Dtu_statusInfoId     int    `json:"dtu_statusInfoId" gorm:"size:10;primary_key;AUTO_INCREMENT"`
	Dtu_uptime time.Time  `json:"dtu_uptime"`
	Pkg_id   string `json:"pkg_id" gorm:"size:20;"`
	Dtu_id      string `json:"dtu_id" gorm:"size:20;"`
	Dtu_latitudeType   uint8    `json:"dtu_latitudeType" gorm:"Type：uint8"`
	Dtu_longitudeType   uint8    `json:"dtu_longitudeType" gorm:"Type：uint8"`
	Dtu_latitude   int    `json:"dtu_latitude" gorm:"Type：int"`
	Dtu_longitude   int    `json:"dtu_longitude" gorm:"Type：int"`
	Dtu_csq   uint8    `json:"dtu_csq" gorm:"Type：uint8"`
	Dtu_locateMode   uint8    `json:"dtu_locateMode" gorm:"Type：uint8"`
	Dtu_gpsSateCnt   uint8    `json:"dtu_gpsSateCnt" gorm:"Type：uint8"`
	Dtu_speed   uint16    `json:"dtu_speed" gorm:"Type：uint16"`
	Dtu_altitude   uint16    `json:"dtu_altitude" gorm:"Type：uint16"`
	Dtu_pluginVoltage   uint8    `json:"dtu_pluginVoltage" gorm:"Type：uint8"`
	Dtu_selfInVoltage   uint8    `json:"dtu_selfInVoltage" gorm:"Type：uint8"`
	Dtu_errStatus   uint8    `json:"dtu_errStatus" gorm:"Type：uint8"`
	Dtu_errNbr   uint8    `json:"dtu_errNbr" gorm:"Type：uint8"`
	Dtu_errCode   uint16    `json:"dtu_errCode" gorm:"Type：uint16"`
	models.BaseModel
}
func (Dtu_statusinfo) TableName() string {
	return "user_dtu_statusinfo"
}
