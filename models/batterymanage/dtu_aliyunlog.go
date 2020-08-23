package batterymanage

import (
	"go-admin/models"
	"time"
)
//dtu为主，dtu pkg脱离后只保留 dtu
type Dtu_aliyunLog struct {
	Dtu_aliyunLogId     int    `json:"dtu_aliyunLogId" gorm:"size:10;primary_key;AUTO_INCREMENT"`
	Dtu_uptime time.Time  `json:"dtu_uptime"`
	Dtu_id      string `json:"dtu_id" gorm:"size:20;primary_key;unique;not null;"`
	Dtu_aliyunStatus uint8    `json:"dtu_aliyunStatus" gorm:"Type：uint8"`
	DataScope  string `json:"dataScope" gorm:"-"`
	UpdateBy  string `gorm:"size:128;" json:"updateBy"`
	models.BaseModel
}
func (Dtu_aliyunLog) TableName() string {
	return "user_dtu_aliyunlog"
}
