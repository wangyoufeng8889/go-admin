package batterymanage

import (
	orm "go-admin/global"
	"go-admin/models"
	"go-admin/tools"
	"time"
)

type DashboardInfo struct {
	BatteryTotalNbr int `json:"batteryTotalNbr" gorm:"-"`
	BatteryOnlineNbr int `json:"batteryOnlineNbr" gorm:"-"`
	BatteryOfflineNbr int `json:"batteryOfflineNbr" gorm:"-"`

	DtuTotalNbr int `json:"dtuTotalNbr" gorm:"-"`
	DtuOnlineNbr int `json:"dtuOnlineNbr" gorm:"-"`
	DtuOfflineNbr int `json:"dtuOfflineNbr" gorm:"-"`

	DtuLocation []LocalInfo

	DataScope  string `json:"dataScope" gorm:"-"`
}
type LocalInfo struct {
	//Dtu_statusInfo
	Dtu_statusInfoId     int    `json:"dtu_statusInfoLogId" gorm:"size:10;primary_key;AUTO_INCREMENT"`
	Dtu_uptime time.Time  `json:"dtu_uptime"`
	Dtu_id      string `json:"dtu_id" gorm:"size:20;"`
	Pkg_id   string `json:"pkg_id" gorm:"size:20;"`
	Dtu_latitudeType   string    `json:"dtu_latitudeType" gorm:"Type：size:2"`
	Dtu_longitudeType   string    `json:"dtu_longitudeType" gorm:"Type：size:2"`
	Dtu_latitude   string    `json:"dtu_latitude" gorm:"Type：size:20"`
	Dtu_longitude   string    `json:"dtu_longitude" gorm:"Type：size:20"`
	//无关数据库
	DataScope  string `json:"dataScope" gorm:"-"`
	models.BaseModel
}
func (e *DashboardInfo) GetDashboardInfo() (DashboardInfo, error) {
	var doc DashboardInfo

	//电池在线数据
	table := orm.Eloquent.Table("user_bms_specinfo")
	// 数据权限控制
	dataPermission := new(models.DataPermission)
	dataPermission.UserId, _ = tools.StringToInt(e.DataScope)
	table, err := dataPermission.GetDataScope("user_bms_specinfo", table)
	if err != nil {
		return doc, err
	}

	var count int
	table.Where("`deleted_at` IS NULL").Count(&count)
	if table.Error != nil {
		return doc, table.Error
	}
	doc.BatteryTotalNbr = count

	table = orm.Eloquent.Table("user_dtu_specinfo").Where("dtu_aliyun_status = ?", 1).
		Where("user_dtu_specinfo.pkg_id != ?", "").
		Where("`deleted_at` IS NULL").
		Count(&count)
	if table.Error != nil {
		return doc, table.Error
	}
	doc.BatteryOnlineNbr = count
	doc.BatteryOfflineNbr = doc.BatteryTotalNbr - doc.BatteryOnlineNbr

	//dtu在线数据
	table = orm.Eloquent.Table("user_dtu_specinfo").Where("`deleted_at` IS NULL").Count(&count)
	if table.Error != nil {
		return doc, table.Error
	}
	doc.DtuTotalNbr = count

	table = orm.Eloquent.Table("user_dtu_specinfo").Where("dtu_aliyun_status = ?", 1).
		Where("`deleted_at` IS NULL").
		Count(&count)
	if table.Error != nil {
		return doc, table.Error
	}
	doc.DtuOnlineNbr = count
	doc.DtuOfflineNbr = doc.DtuTotalNbr - doc.DtuOnlineNbr

	var docLocal []LocalInfo
	table = orm.Eloquent.Table("user_dtu_statusinfo").Select([]string{"user_dtu_statusinfo.dtu_status_info_id",
		"user_dtu_statusinfo.dtu_uptime",
		"user_dtu_statusinfo.dtu_id",
		"user_dtu_statusinfo.pkg_id",
		"user_dtu_statusinfo.dtu_latitude_type",
		"user_dtu_statusinfo.dtu_longitude_type",
		"user_dtu_statusinfo.dtu_latitude",
		"user_dtu_statusinfo.dtu_longitude"})

	// 数据权限控制
	dataPermission = new(models.DataPermission)
	dataPermission.UserId, _ = tools.StringToInt(e.DataScope)
	table, err = dataPermission.GetDataScope("user_dtu_statusinfo", table)
	if err != nil {
		return doc, err
	}
	table = table.Order("dtu_status_info_id").Find(&docLocal)
	if table.Error != nil {
		return doc, err
	}
	if err:=table.Where("`deleted_at` IS NULL").Where("user_dtu_statusinfo.dtu_latitude <> ?", "0").Find(&docLocal).Error;err!= nil{
		return doc, err
	}
	doc.DtuLocation = docLocal

	return doc, nil
}
