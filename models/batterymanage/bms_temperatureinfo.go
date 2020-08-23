package batterymanage
import (
	"go-admin/models"
	"time"
)
type Bms_temperatureInfo struct {
	Bms_temperatureInfoId     int    `json:"bms_temperatureInfoId" gorm:"size:10;primary_key;AUTO_INCREMENT"`
	Dtu_uptime time.Time  `json:"dtu_uptime"`
	Pkg_id   string `json:"pkg_id" gorm:"size:20;primary_key;"`
	Dtu_id      string `json:"dtu_id" gorm:"size:20;"`
	Bms_temperature1   uint8    `json:"bms_temperature1" gorm:"Type：uint8"`
	Bms_temperature2   uint8    `json:"bms_temperature2" gorm:"Type：uint8"`
	Bms_temperature3   uint8    `json:"bms_temperature3" gorm:"Type：uint8"`
	Bms_temperature4   uint8    `json:"bms_temperature4" gorm:"Type：uint8"`
	Bms_temperature5   uint8    `json:"bms_temperature5" gorm:"Type：uint8"`
	Bms_temperature6   uint8    `json:"bms_temperature6" gorm:"Type：uint8"`
	DataScope  string `json:"dataScope" gorm:"-"`
	UpdateBy  string `gorm:"size:128;" json:"updateBy"`
	models.BaseModel
}
func (Bms_temperatureInfo) TableName() string {
	return "user_bms_temperatureinfo"
}
