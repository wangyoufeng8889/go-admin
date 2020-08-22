package batterymanage


import (
	orm "go-admin/global"
	"go-admin/models"
	"go-admin/tools"
	"time"
)
type Bms_specinfo struct {
	Bms_specinfoId     int    `json:"bms_specinfoId" gorm:"size:10;primary_key;AUTO_INCREMENT"`
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
func (Bms_specinfo) TableName() string {
	return "user_bms_specinfo"
}
func (e *Bms_specinfo) GetPage(pageSize int, pageIndex int) ([]Bms_specinfo, int, error) {
	var doc []Bms_specinfo

	table := orm.Eloquent.Select("*").Table(e.TableName())
	if e.Bms_specinfoId != 0 {
		//按照数据库格式
		table = table.Where("bms_specinfo_id = ?", e.Bms_specinfoId)
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

	if err := table.Order("bms_specinfo_id").Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
		return nil, 0, err
	}
	table.Where("`deleted_at` IS NULL").Count(&count)
	return doc, count, nil
}
func (e *Bms_specinfo) BatchDelete(id []int) (Result bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("bms_specinfo_id in (?)", id).Delete(&Bms_specinfo{}).Error; err != nil {
		return
	}
	Result = true
	return
}
