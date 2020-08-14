package batterymanage


import (
	"go-admin/models"
	"time"
)
type Bms_specinfo struct {
	Dtu_specinfoId     int    `json:"dtu_specinfoId" gorm:"size:10;primary_key;AUTO_INCREMENT"`
	Dtu_uptime time.Time  `json:"dtu_uptime"`
	Pkg_id   string `json:"pkg_id" gorm:"size:20;"`
	Dtu_id      string `json:"dtu_id" gorm:"size:20;"`
	Bms_id      string `json:"bms_id" gorm:"size:20;"`
	Pkg_count   uint8    `json:"pkg_count" gorm:"Type：uint8"`
	Pkg_type   uint8    `json:"pkg_type" gorm:"Type：uint8"`
	Pkg_capacity   uint16    `json:"pkg_capacity" gorm:"Type：uint16"`
	Pkg_nominalVoltage   uint16    `json:"pkg_nominalVoltage" gorm:"Type：uint16"`
	Pkg_ntcCount   uint8    `json:"pkg_ntcCount" gorm:"Type：uint8"`
	Pkg_manufactureDate time.Time  `json:"pkg_manufactureDate"`
	Bms_hardVer   uint8    `json:"bms_hardVer" gorm:"Type：uint8"`
	Bms_softVer   uint8    `json:"bms_softVer" gorm:"Type：uint8"`
	Bms_protocolVer   uint16    `json:"bms_protocolVer" gorm:"Type：uint16"`
	models.BaseModel
}
func (Bms_specinfo) TableName() string {
	return "user_bms_specinfo"
}
