package batterymanage

import (
	orm "go-admin/global"
	"go-admin/models"
	"go-admin/tools"
	"time"
)

type Bms_cellInfoLog struct {
	Bms_cellInfoLogId     int    `json:"bms_cellInfoLogId" gorm:"size:10;primary_key;AUTO_INCREMENT"`
	Dtu_uptime time.Time  `json:"dtu_uptime"`
	Pkg_id   string `json:"pkg_id" gorm:"size:20;"`
	Dtu_id      string `json:"dtu_id" gorm:"size:20;"`
	Bms_cellVoltage1   uint16    `json:"bms_cellVoltage1" gorm:"Type：uint16"`
	Bms_cellVoltage2   uint16    `json:"bms_cellVoltage2" gorm:"Type：uint16"`
	Bms_cellVoltage3   uint16    `json:"bms_cellVoltage3" gorm:"Type：uint16"`
	Bms_cellVoltage4   uint16    `json:"bms_cellVoltage4" gorm:"Type：uint16"`
	Bms_cellVoltage5   uint16    `json:"bms_cellVoltage5" gorm:"Type：uint16"`
	Bms_cellVoltage6   uint16    `json:"bms_cellVoltage6" gorm:"Type：uint16"`
	Bms_cellVoltage7   uint16    `json:"bms_cellVoltage7" gorm:"Type：uint16"`
	Bms_cellVoltage8   uint16    `json:"bms_cellVoltage8" gorm:"Type：uint16"`
	Bms_cellVoltage9   uint16    `json:"bms_cellVoltage9" gorm:"Type：uint16"`
	Bms_cellVoltage10   uint16    `json:"bms_cellVoltage10" gorm:"Type：uint16"`
	Bms_cellVoltage11   uint16    `json:"bms_cellVoltage11" gorm:"Type：uint16"`
	Bms_cellVoltage12   uint16    `json:"bms_cellVoltage12" gorm:"Type：uint16"`
	Bms_cellVoltage13   uint16    `json:"bms_cellVoltage13" gorm:"Type：uint16"`
	Bms_cellVoltage14   uint16    `json:"bms_cellVoltage14" gorm:"Type：uint16"`
	Bms_cellVoltage15   uint16    `json:"bms_cellVoltage15" gorm:"Type：uint16"`
	Bms_cellVoltage16   uint16    `json:"bms_cellVoltage16" gorm:"Type：uint16"`
	Bms_cellVoltage17   uint16    `json:"bms_cellVoltage17" gorm:"Type：uint16"`
	Bms_cellVoltage18   uint16    `json:"bms_cellVoltage18" gorm:"Type：uint16"`
	Bms_cellVoltage19   uint16    `json:"bms_cellVoltage19" gorm:"Type：uint16"`
	Bms_cellVoltage20   uint16    `json:"bms_cellVoltage20" gorm:"Type：uint16"`
	DataScope  string `json:"dataScope" gorm:"-"`
	UpdateBy  string `gorm:"size:128;" json:"updateBy"`
	models.BaseModel
}
func (Bms_cellInfoLog) TableName() string {
	return "user_bms_cellinfolog"
}
func (e *Bms_cellInfoLog) GetBms_cellInfoLog(starttime time.Time, endtime time.Time,dateflag int) ([]Bms_cellInfoLog,int, error) {
	var doc []Bms_cellInfoLog

	table := orm.Eloquent.Select("*").Table(e.TableName())
	// 数据权限控制
	dataPermission := new(models.DataPermission)
	dataPermission.UserId, _ = tools.StringToInt(e.DataScope)
	table, err := dataPermission.GetDataScope(e.TableName(), table)
	if err != nil {
		return nil, 0, err
	}
	var count int
	table = table.Order("user_bms_cellinfolog.bms_cell_info_log_id").Find(&doc)
	if table.Error != nil {
		return nil, 0, err
	}
	if e.Bms_cellInfoLogId != 0 {
		table = table.Where("user_bms_cellinfolog.bms_cell_info_log_id = ?", e.Bms_cellInfoLogId)
	}
	if e.Pkg_id != "" {
		table = table.Where("user_bms_cellinfolog.pkg_id = ?", e.Pkg_id)
	}
	if e.Dtu_id != "" {
		table = table.Where("user_bms_cellinfolog.dtu_id = ?", e.Dtu_id)
	}
	table_date := table.Where("user_bms_cellinfolog.dtu_uptime BETWEEN ? AND ?",starttime,endtime)
	if err:=table_date.Where("`deleted_at` IS NULL").Find(&doc).Count(&count).Error;err!= nil{
		return nil, 0, err
	}
	if count == 0 && dateflag != 1 {
		if err:=table.Where("`deleted_at` IS NULL").Limit(100).Find(&doc).Count(&count).Error;err!= nil{
			return nil, 0, err
		}
	}
	return doc, count, nil
}

