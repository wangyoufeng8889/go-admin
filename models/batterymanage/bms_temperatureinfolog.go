package batterymanage
import (
	orm "go-admin/global"
	"go-admin/models"
	"go-admin/tools"
	"time"
)
type Bms_temperatureInfoLog struct {
	Bms_temperatureInfoLogId     int    `json:"bms_temperatureInfoLogId" gorm:"size:10;primary_key;AUTO_INCREMENT"`
	Dtu_uptime time.Time  `json:"dtu_uptime"`
	Pkg_id   string `json:"pkg_id" gorm:"size:20;"`
	Dtu_id      string `json:"dtu_id" gorm:"size:20;"`
	Bms_temperature1   uint8    `json:"bms_temperature1" gorm:"Type：uint8"`
	Bms_temperature2   uint8    `json:"bms_temperature2" gorm:"Type：uint8"`
	Bms_temperature3   uint8    `json:"bms_temperature3" gorm:"Type：uint8"`
	Bms_temperature4   uint8    `json:"bms_temperature4" gorm:"Type：uint8"`
	Bms_temperature5   uint8    `json:"bms_temperature5" gorm:"Type：uint8"`
	Bms_temperature6   uint8    `json:"bms_temperature6" gorm:"Type：uint8"`
	DataScope  string `json:"dataScope" gorm:"-"`
	UpdateBy  string `gorm:"size:128;" json:"updateBy"`
	models.BaseModel
}
func (Bms_temperatureInfoLog) TableName() string {
	return "user_bms_temperatureinfolog"
}
func (e *Bms_temperatureInfoLog) GetBms_temperatureInfoLog(starttime time.Time, endtime time.Time) ([]Bms_temperatureInfoLog,int, error) {
	var doc []Bms_temperatureInfoLog

	table := orm.Eloquent.Select("*").Table(e.TableName())
	// 数据权限控制
	dataPermission := new(models.DataPermission)
	dataPermission.UserId, _ = tools.StringToInt(e.DataScope)
	table, err := dataPermission.GetDataScope(e.TableName(), table)
	if err != nil {
		return nil, 0, err
	}
	var count int
	table = table.Order("user_bms_temperatureinfolog.Bms_temperature_info_log_id").Find(&doc)
	if table.Error != nil {
		return nil, 0, err
	}
	if e.Bms_temperatureInfoLogId != 0 {
		table = table.Where("user_bms_temperatureinfolog.Bms_temperature_info_log_id = ?", e.Bms_temperatureInfoLogId)
	}
	if e.Pkg_id != "" {
		table = table.Where("user_bms_temperatureinfolog.pkg_id = ?", e.Pkg_id)
	}
	if e.Dtu_id != "" {
		table = table.Where("user_bms_temperatureinfolog.dtu_id = ?", e.Dtu_id)
	}
	table = table.Where("user_bms_temperatureinfolog.dtu_uptime BETWEEN ? AND ?",starttime,endtime)
	if err:=table.Where("`deleted_at` IS NULL").Find(&doc).Count(&count).Error;err!= nil{
		return nil, 0, err
	}
	return doc, count, nil
}