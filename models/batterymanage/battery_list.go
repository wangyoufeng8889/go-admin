package batterymanage


import (
	orm "go-admin/global"
	"go-admin/models"
	"go-admin/tools"
	"time"
)
type Battery_list struct {
	Battery_listId     int    `json:"battery_listId" gorm:"size:10;primary_key;AUTO_INCREMENT"`
	Dtu_uptime time.Time  `json:"dtu_uptime"`
	Pkg_id   string `json:"pkg_id" gorm:"size:20;"`
	Dtu_id      string `json:"dtu_id" gorm:"size:20;"`
	Bms_chargeStatus   uint8    `json:"bms_chargeStatus" gorm:"Type：uint8"`
	Bms_soc   uint8    `json:"bms_soc" gorm:"Type：uint8"`
	Pkg_onOffLineStatus   uint8    `json:"pkg_onOffLineStatus" gorm:"Type：uint8"`
	DTU_onOffLineStatus   uint8    `json:"dtu_onOffLineStatus" gorm:"Type：uint8"`
	Pkg_errStatus   uint8    `json:"pkg_errStatus" gorm:"Type：uint8"`
	Pkg_abnormalStatus   uint8    `json:"pkg_abnormalStatus" gorm:"Type：uint8"`
	Pkg_usableStatus   uint8    `json:"pkg_usableStatus" gorm:"Type：uint8"`
	Pkg_runStatus   uint8    `json:"pkg_runStatus" gorm:"Type：uint8"`
	Pkg_count   uint8    `json:"pkg_count" gorm:"Type：uint8"`
	Pkg_type   uint8    `json:"pkg_type" gorm:"Type：uint8"`
	Pkg_capacity   uint16    `json:"pkg_capacity" gorm:"Type：uint16"`
	Pkg_nominalVoltage   uint16    `json:"pkg_nominalVoltage" gorm:"Type：uint16"`
	Dtu_type   uint8    `json:"dtu_type" gorm:"Type：uint8"`
	Dtu_setupType   uint8    `json:"dtu_setupType" gorm:"Type：uint8"`
	Dtu_latitudeType   string    `json:"dtu_latitudeType" gorm:"Type：size:2"`
	Dtu_longitudeType   string    `json:"dtu_longitudeType" gorm:"Type：size:2"`
	Dtu_latitude   string    `json:"dtu_latitude" gorm:"Type：size:20"`
	Dtu_longitude   string    `json:"dtu_longitude" gorm:"Type：size:20"`
	Dtu_csq   uint8    `json:"dtu_csq" gorm:"Type：uint8"`
	Dtu_locateMode   uint8    `json:"dtu_locateMode" gorm:"Type：uint8"`

	DataScope  string `json:"dataScope" gorm:"-"`
	UpdateBy  string `gorm:"size:128;" json:"updateBy"`
	models.BaseModel
}
func (Battery_list) TableName() string {
	return "user_battery_list"
}
func (e *Battery_list) GetPage(pageSize int, pageIndex int) ([]Battery_list, int, error) {
	var doc []Battery_list

	table := orm.Eloquent.Select("*").Table(e.TableName())
	if e.Battery_listId != 0 {
		table = table.Where("battery_list_id = ?", e.Battery_listId)
	}
	if e.Pkg_id != "" {
		table = table.Where("pkg_id = ?", e.Pkg_id)
	}else {
		table = table.Not("pkg_id = ?", "")
	}
	if e.Dtu_id != "" {
		table = table.Where("dtu_id = ?", e.Dtu_id)
	}
	if e.Bms_chargeStatus < 0xFF {
		table = table.Where("bms_charge_status = ?", e.Bms_chargeStatus)
	}
	if e.Pkg_onOffLineStatus < 0xFF {
		table = table.Where("pkg_on_off_line_status = ?", e.Pkg_onOffLineStatus)
	}


	// 数据权限控制
	dataPermission := new(models.DataPermission)
	dataPermission.UserId, _ = tools.StringToInt(e.DataScope)
	table, err := dataPermission.GetDataScope(e.TableName(), table)
	if err != nil {
		return nil, 0, err
	}
	var count int

	if err := table.Order("battery_list_id").Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
		return nil, 0, err
	}
	table.Where("`deleted_at` IS NULL").Count(&count)
	return doc, count, nil
}
func (e *Battery_list) BatchDelete(id []int) (Result bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("battery_list_id in (?)", id).Delete(&Battery_list{}).Error; err != nil {
		return
	}
	Result = true
	return
}
