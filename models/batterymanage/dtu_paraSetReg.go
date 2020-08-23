package batterymanage
import (
	"go-admin/models"
	"time"
)
type Dtu_paraSetReg struct {
	Dtu_paraSetRegId     int    `json:"dtu_paraSetRegId" gorm:"size:10;primary_key;AUTO_INCREMENT"`
	Dtu_uptime time.Time  `json:"dtu_uptime"`
	Dtu_id      string `json:"dtu_id" gorm:"size:20;primary_key;"`
	Pkg_id   string `json:"pkg_id" gorm:"size:20;"`
	Dtu_pkgInfoReportPeriod   uint16    `json:"dtu_pkgInfoReportPeriod" gorm:"Type：uint16"`
	Dtu_remoteLockCar   uint16    `json:"dtu_remoteLockCar" gorm:"Type：uint16"`
	Dtu_voiceTipsOnOff   uint8    `json:"dtu_voiceTipsOnOff" gorm:"Type：uint8"`
	Dtu_voiceTipsThresholdValue   uint8    `json:"dtu_voiceTipsThresholdValue" gorm:"Type：uint8"`
	Dtu_voiceTipsDownBulk   uint8    `json:"dtu_voiceTipsDownBulk" gorm:"Type：uint8"`
	Dtu_otaIP      string `json:"dtu_otaIP" gorm:"size:20;"`
	DataScope  string `json:"dataScope" gorm:"-"`
	UpdateBy  string `gorm:"size:128;" json:"updateBy"`
	models.BaseModel
}
func (Dtu_paraSetReg) TableName() string {
	return "user_dtu_paraSetReg"
}
