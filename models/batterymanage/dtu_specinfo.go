package batterymanage

import (
	orm "go-admin/global"
	"go-admin/models"
	"go-admin/tools"
	"time"
)
type Dtu_specInfo struct {
	Dtu_specInfoId     int    `json:"dtu_specInfoId" gorm:"size:10;primary_key;AUTO_INCREMENT"`
	Dtu_uptime time.Time  `json:"dtu_uptime"`
	Dtu_id      string `json:"dtu_id" gorm:"size:20;primary_key;unique;not null;"`
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
	DataScope  string `json:"dataScope" gorm:"-"`
	UpdateBy  string `gorm:"size:128;" json:"updateBy"`
	models.BaseModel
}
func (Dtu_specInfo) TableName() string {
	return "user_dtu_specinfo"
}
func (e *Dtu_specInfo) Getdtu_specinfo(is_oneList string) ([]Dtu_specInfo,int, error) {
	var doc []Dtu_specInfo

	table := orm.Eloquent.Select("*").Table(e.TableName())
	if e.Dtu_specInfoId != 0 {
		table = table.Where("dtu_spec_info_id = ?", e.Dtu_specInfoId)
	}
	if e.Pkg_id != "" {
		table = table.Where("pkg_id = ?", e.Pkg_id)
	}else {
		table = table.Not("pkg_id = ?", "0")
	}
	if e.Dtu_id != "" {
		table = table.Where("dtu_id = ?", e.Dtu_id)
	}
	// 数据权限控制
	dataPermission := new(models.DataPermission)
	dataPermission.UserId, _ = tools.StringToInt(e.DataScope)
	table, err := dataPermission.GetDataScope(e.TableName(), table)
	if err != nil {
		return nil, 0, err
	}
	var count int
	if is_oneList == "YES" {
		if err := table.Order("dtu_uptime desc").First(&doc).Error; err != nil {
			return nil, 0, err
		}
		table.Where("`deleted_at` IS NULL").Count(&count)
	}else{
		if err := table.Order("dtu_spec_info_id").Find(&doc).Error; err != nil {
			return nil, 0, err
		}
		table.Where("`deleted_at` IS NULL").Count(&count)
	}
	return doc, count, nil
}