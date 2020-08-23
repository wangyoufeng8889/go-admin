package batterymanage

import (
	orm "go-admin/global"
	"go-admin/models"
	"go-admin/tools"
	"time"
)
//必须保持Dtu_statusInfo与Dtu_statusInfoLog结构体一致
type Dtu_statusInfoLog struct {
	Dtu_statusInfoLogId     int    `json:"dtu_statusInfoLogId" gorm:"size:10;primary_key;AUTO_INCREMENT"`
	Dtu_uptime time.Time  `json:"dtu_uptime"`
	Dtu_id      string `json:"dtu_id" gorm:"size:20;"`
	Pkg_id   string `json:"pkg_id" gorm:"size:20;"`
	Dtu_latitudeType   string    `json:"dtu_latitudeType" gorm:"Type：size:2"`
	Dtu_longitudeType   string    `json:"dtu_longitudeType" gorm:"Type：size:2"`
	Dtu_latitude   string    `json:"dtu_latitude" gorm:"Type：size:20"`
	Dtu_longitude   string    `json:"dtu_longitude" gorm:"Type：size:20"`
	Dtu_csq   uint8    `json:"dtu_csq" gorm:"Type：uint8"`
	Dtu_locateMode   uint8    `json:"dtu_locateMode" gorm:"Type：uint8"`
	Dtu_gpsSateCnt   uint8    `json:"dtu_gpsSateCnt" gorm:"Type：uint8"`
	Dtu_speed   uint16    `json:"dtu_speed" gorm:"Type：uint16"`
	Dtu_altitude   uint16    `json:"dtu_altitude" gorm:"Type：uint16"`
	Dtu_pluginVoltage   uint8    `json:"dtu_pluginVoltage" gorm:"Type：uint8"`
	Dtu_selfInVoltage   uint8    `json:"dtu_selfInVoltage" gorm:"Type：uint8"`
	Dtu_errStatus   uint8    `json:"dtu_errStatus" gorm:"Type：uint8"`
	Dtu_errNbr   uint8    `json:"dtu_errNbr" gorm:"Type：uint8"`
	Dtu_errCode   uint16    `json:"dtu_errCode" gorm:"Type：uint16"`
	DataScope  string `json:"dataScope" gorm:"-"`
	UpdateBy  string `gorm:"size:128;" json:"updateBy"`
	models.BaseModel
}
func (Dtu_statusInfoLog) TableName() string {
	return "user_dtu_statusinfolog"
}
type BatteryMoveInfo struct {
	//Dtu_statusInfo
	Dtu_statusInfoLogId     int    `json:"dtu_statusInfoLogId" gorm:"size:10;primary_key;AUTO_INCREMENT"`
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
func (BatteryMoveInfo) TableName() string {
	return "user_dtu_statusinfolog"
}
func (e *BatteryMoveInfo) GetBatteryMoveInfo(starttime time.Time, endtime time.Time) ([]BatteryMoveInfo,int, error) {
	var doc []BatteryMoveInfo

	table := orm.Eloquent.Table(e.TableName()).Select([]string{"user_dtu_statusinfolog.dtu_status_info_log_id",
		"user_dtu_statusinfolog.dtu_uptime",
		"user_dtu_statusinfolog.dtu_id",
		"user_dtu_statusinfolog.pkg_id",
		"user_dtu_statusinfolog.dtu_latitude_type",
		"user_dtu_statusinfolog.dtu_longitude_type",
		"user_dtu_statusinfolog.dtu_latitude",
		"user_dtu_statusinfolog.dtu_longitude"})

	// 数据权限控制
	dataPermission := new(models.DataPermission)
	dataPermission.UserId, _ = tools.StringToInt(e.DataScope)
	table, err := dataPermission.GetDataScope(e.TableName(), table)
	if err != nil {
		return nil, 0, err
	}
	var count int
	table = table.Order("dtu_status_info_log_id").Find(&doc)
	if table.Error != nil {
		return nil, 0, err
	}
	if e.Dtu_statusInfoLogId != 0 {
		table = table.Where("dtu_status_info_log_id = ?", e.Dtu_statusInfoLogId)
	}

	if e.Pkg_id != "" {
		table = table.Where("user_dtu_statusinfolog.pkg_id = ?", e.Pkg_id)
	}
	if e.Dtu_id != "" {
		table = table.Where("user_dtu_statusinfolog.dtu_id = ?", e.Dtu_id)
	}
	table = table.Where("user_dtu_statusinfolog.dtu_uptime BETWEEN ? AND ?",starttime,endtime)
	if err:=table.Where("`deleted_at` IS NULL").Find(&doc).Count(&count).Error;err!= nil{
		return nil, 0, err
	}
	return doc, count, nil
}
