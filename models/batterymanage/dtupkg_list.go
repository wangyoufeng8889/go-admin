package batterymanage

import (
	"go-admin/models"
	"time"
)
//dtu为主，dtu pkg脱离后只保留 dtu
type DtuPkg_list struct {
	Dtu_listId     int    `json:"dtu_listId" gorm:"size:10;primary_key;AUTO_INCREMENT"`
	Bind_uptime time.Time  `json:"bind_uptime"`
	Dtu_id      string `json:"dtu_id" gorm:"size:20;"`
	Pkg_id   string `json:"pkg_id" gorm:"size:20;"`

	DataScope  string `json:"dataScope" gorm:"-"`
	UpdateBy  string `gorm:"size:128;" json:"updateBy"`
	models.BaseModel
}
func (DtuPkg_list) TableName() string {
	return "user_dtupkg_list"
}