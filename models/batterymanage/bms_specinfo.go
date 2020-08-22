package batterymanage


import (
	orm "go-admin/global"
	"go-admin/models"
	"go-admin/tools"
	"time"
)
type Bms_specInfo struct {
	Bms_specInfoId     int    `json:"bms_specInfoId" gorm:"size:10;primary_key;AUTO_INCREMENT"`
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
	Bms_protocolVer   string    `json:"bms_protocolVer" gorm:"Type：size:10"`
	DataScope  string `json:"dataScope" gorm:"-"`
	UpdateBy  string `gorm:"size:128;" json:"updateBy"`
	models.BaseModel
}
func (Bms_specInfo) TableName() string {
	return "user_bms_specinfo"
}
//电池列表
func (e *Bms_specInfo) GetBms_specinfo(pageSize int, pageIndex int,is_oneList string) ([]Bms_specInfo,int, error) {
	var doc []Bms_specInfo

	table := orm.Eloquent.Select("*").Table(e.TableName())
	if e.Bms_specInfoId != 0 {
		table = table.Where("bms_spec_info_id = ?", e.Bms_specInfoId)
	}
	if e.Pkg_id != "" {
		table = table.Where("pkg_id = ?", e.Pkg_id)
	}else {
		table = table.Not("pkg_id = ?", "")
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
		if err := table.Order("bms_spec_info_id").Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
			return nil, 0, err
		}
		table.Where("`deleted_at` IS NULL").Count(&count)
	}
	return doc, count, nil
}
func (e *Bms_specInfo) BatchDelete(id []int) (Result bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("bms_spec_info_id in (?)", id).Delete(&Bms_specInfo{}).Error; err != nil {
		return
	}
	Result = true
	return
}
