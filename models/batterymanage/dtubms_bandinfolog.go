package batterymanage

import (
	orm "go-admin/global"
	"go-admin/models"
	"go-admin/tools"
	"time"
)

type DtuBms_BandInfoLog struct {
	DtuBms_BandInfoLogId     int    `json:"dtuBms_BandInfoLogId" gorm:"size:10;primary_key;AUTO_INCREMENT"`
	Dtu_uptime time.Time  `json:"dtu_uptime"`
	Dtu_id      string `json:"dtu_id" gorm:"size:20;not null;"`
	Pkg_id   string `json:"pkg_id" gorm:"size:20;"`
	DataScope  string `json:"dataScope" gorm:"-"`
	UpdateBy  string `gorm:"size:128;" json:"updateBy"`
	models.BaseModel
}
func (DtuBms_BandInfoLog) TableName() string {
	return "user_dtu_bms_band_info_log"
}


func (e *DtuBms_BandInfoLog) GetDtuBmsBandListInfo(pageSize int, pageIndex int) ([]DtuBms_BandInfoLog,int, error) {
	var doc []DtuBms_BandInfoLog

	table := orm.Eloquent.Table(e.TableName()).Select([]string{"user_dtu_bms_band_info_log.dtu_bms_band_info_log_id",
		"user_dtu_bms_band_info_log.dtu_uptime",
		"user_dtu_bms_band_info_log.pkg_id",
		"user_dtu_bms_band_info_log.dtu_id"})
	if e.DtuBms_BandInfoLogId != 0 {
		table = table.Where("user_dtu_bms_band_info_log.dtu_spec_info_id = ?", e.DtuBms_BandInfoLogId)
	}
	if e.Dtu_id != "" {
		table = table.Where("user_dtu_bms_band_info_log.dtu_id = ?", e.Dtu_id)
	}
	if e.Pkg_id != "" {
		table = table.Where("user_dtu_bms_band_info_log.pkg_id = ?", e.Pkg_id)
	}
	// 数据权限控制
	dataPermission := new(models.DataPermission)
	dataPermission.UserId, _ = tools.StringToInt(e.DataScope)
	table, err := dataPermission.GetDataScope(e.TableName(), table)
	if err != nil {
		return nil, 0, err
	}
	var count int
	if err := table.Order("user_dtu_bms_band_info_log.dtu_uptime desc").Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
		return nil, 0, err
	}
	table.Where("user_dtu_bms_band_info_log.deleted_at IS NULL").Count(&count)
	return doc, count, nil
}

