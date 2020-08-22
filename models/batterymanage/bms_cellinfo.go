package batterymanage
import (
	"go-admin/models"
	"time"
)
type Bms_cellinfo struct {
	Bms_cellinfoId     int    `json:"bms_cellinfoId" gorm:"size:10;primary_key;AUTO_INCREMENT"`
	Dtu_uptime time.Time  `json:"dtu_uptime"`
	Dtu_id      string `json:"dtu_id" gorm:"size:20;"`
	Pkg_id   string `json:"pkg_id" gorm:"size:20;"`
	Bms_cellVoltage1   uint16    `json:"bms_cellVoltage1" gorm:"Type：uint16"`
	Bms_cellVoltage2   uint16    `json:"bms_cellVoltage2" gorm:"Type：uint16"`
	Bms_cellVoltage3   uint16    `json:"bms_cellVoltage3" gorm:"Type：uint16"`
	Bms_cellVoltage4   uint16    `json:"bms_cellVoltage4" gorm:"Type：uint16"`
	Bms_cellVoltage5   uint16    `json:"bms_cellVoltage5" gorm:"Type：uint16"`
	Bms_cellVoltage6   uint16    `json:"bms_cellVoltage6" gorm:"Type：uint16"`
	Bms_cellVoltage7   uint16    `json:"bms_cellVoltage7" gorm:"Type：uint16"`
	Bms_cellVoltage8   uint16    `json:"bms_cellVoltage8" gorm:"Type：uint16"`
	Bms_cellVoltage9   uint16    `json:"bms_cellVoltage9" gorm:"Type：uint16"`
	Bms_cellVoltage10   uint16    `json:"bms_cellVoltage10" gorm:"Type：uint16"`
	Bms_cellVoltage11   uint16    `json:"bms_cellVoltage11" gorm:"Type：uint16"`
	Bms_cellVoltage12   uint16    `json:"bms_cellVoltage12" gorm:"Type：uint16"`
	Bms_cellVoltage13   uint16    `json:"bms_cellVoltage13" gorm:"Type：uint16"`
	Bms_cellVoltage14   uint16    `json:"bms_cellVoltage14" gorm:"Type：uint16"`
	Bms_cellVoltage15   uint16    `json:"bms_cellVoltage15" gorm:"Type：uint16"`
	Bms_cellVoltage16   uint16    `json:"bms_cellVoltage16" gorm:"Type：uint16"`
	Bms_cellVoltage17   uint16    `json:"bms_cellVoltage17" gorm:"Type：uint16"`
	Bms_cellVoltage18   uint16    `json:"bms_cellVoltage18" gorm:"Type：uint16"`
	Bms_cellVoltage19   uint16    `json:"bms_cellVoltage19" gorm:"Type：uint16"`
	Bms_cellVoltage20   uint16    `json:"bms_cellVoltage20" gorm:"Type：uint16"`
	DataScope  string `json:"dataScope" gorm:"-"`
	UpdateBy  string `gorm:"size:128;" json:"updateBy"`
	models.BaseModel
}
func (Bms_cellinfo) TableName() string {
	return "user_bms_cellinfo"
}