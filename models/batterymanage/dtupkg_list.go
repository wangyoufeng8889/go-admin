package batterymanage

import (
	orm "go-admin/global"
	"go-admin/models"
	"go-admin/tools"
	"time"
)
//dtu为主，dtu pkg脱离后只保留 dtu
type DtuPkg_list struct {
	DtuPkg_listId     int    `json:"dtuPkg_listId" gorm:"size:10;primary_key;AUTO_INCREMENT"`
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
//dtu列表
func (e *DtuPkg_list) GetBms_specinfo(pageSize int, pageIndex int,is_oneList string) ([]DtuPkg_list,int, error) {
	var doc []DtuPkg_list

	table := orm.Eloquent.Select("*").Table(e.TableName())
	if e.DtuPkg_listId != 0 {
		table = table.Where("dtu_pkg_list_id = ?", e.DtuPkg_listId)
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
		if err := table.Order("dtu_pkg_list_id").Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
			return nil, 0, err
		}
		table.Where("`deleted_at` IS NULL").Count(&count)
	}
	return doc, count, nil
}
func (e *DtuPkg_list) BatchDelete(id []int) (Result bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("dtu_pkg_list_id in (?)", id).Delete(&DtuPkg_list{}).Error; err != nil {
		return
	}
	Result = true
	return
}