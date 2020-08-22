package batterymanage
import (
	"go-admin/models"
	"time"
)
//dtu为主，dtu pkg脱离后只保留 dtu
type Dtu_aliyun struct {
	Dtu_listId     int    `json:"dtu_listId" gorm:"size:10;primary_key;AUTO_INCREMENT"`
	Dtu_uptime time.Time  `json:"dtu_uptime"`
	Dtu_id      string `json:"dtu_id" gorm:"size:20;"`
	Dtu_aliyunStatus uint8    `json:"dtu_aliyunStatus" gorm:"Type：uint8"`
	DataScope  string `json:"dataScope" gorm:"-"`
	UpdateBy  string `gorm:"size:128;" json:"updateBy"`
	models.BaseModel
}
func (Dtu_aliyun) TableName() string {
	return "user_dtu_aliyun"
}