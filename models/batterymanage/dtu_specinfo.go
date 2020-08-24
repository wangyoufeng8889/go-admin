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

	//aliyun，pkg去查询是否有dtu绑定，有绑定判断是否在线，没有绑定就不在线
	Dtu_aliyunStatus uint8    `json:"dtu_aliyunStatus" gorm:"Type：uint8"`


	DataScope  string `json:"dataScope" gorm:"-"`
	UpdateBy  string `gorm:"size:128;" json:"updateBy"`
	models.BaseModel
}
func (Dtu_specInfo) TableName() string {
	return "user_dtu_specinfo"
}

type DtuListInfo struct {
	//Dtu_specInfo
	Dtu_specInfoId     int    `json:"dtu_specInfoId" gorm:"size:10;primary_key;AUTO_INCREMENT"`
	Dtu_id      string `json:"dtu_id" gorm:"size:20;primary_key;unique;not null;"`
	Pkg_id   string `json:"pkg_id" gorm:"size:20;"`
	Dtu_type   uint8    `json:"dtu_type" gorm:"Type：uint8"`
	Dtu_setupType   uint8    `json:"dtu_setupType" gorm:"Type：uint8"`
	//aliyun，pkg去查询是否有dtu绑定，有绑定判断是否在线，没有绑定就不在线
	Dtu_aliyunStatus uint8    `json:"dtu_aliyunStatus" gorm:"Type：uint8"`

	//Dtu_statusInfo
	Dtu_uptime time.Time  `json:"dtu_uptime"`
	Dtu_csq   uint8    `json:"dtu_csq" gorm:"Type：uint8"`

	DataScope  string `json:"dataScope" gorm:"-"`
	UpdateBy  string `gorm:"size:128;" json:"updateBy"`
	models.BaseModel

}
func (DtuListInfo) TableName() string {
	return "user_bms_specinfo"
}
func (e *DtuListInfo) Getdtu_listinfo(pageSize int, pageIndex int) ([]DtuListInfo,int, error) {
	var doc []DtuListInfo

	table := orm.Eloquent.Table(e.TableName()).Select([]string{"user_bms_specinfo.bms_spec_info_id",
		"user_bms_specinfo.pkg_id",
		"user_bms_specinfo.pkg_count",
		"user_bms_specinfo.pkg_type",
		"user_bms_specinfo.pkg_capacity",
		"user_bms_specinfo.pkg_nominal_voltage",

		"user_bms_statusinfo.dtu_uptime",
		"user_bms_statusinfo.bms_charge_status",
		"user_bms_statusinfo.bms_soc",
		"user_bms_statusinfo.bms_err_nbr",
		"user_bms_statusinfo.bms_voltage"})
	table = table.Joins("LEFT JOIN user_bms_statusinfo ON user_bms_specinfo.pkg_id=user_bms_statusinfo.pkg_id")
	if e.Dtu_specInfoId != 0 {
		table = table.Where("dtu_spec_info_id = ?", e.Dtu_specInfoId)
	}
	if e.Dtu_id != "" {
		table = table.Where("dtu_id = ?", e.Dtu_id)
	}
	if e.Pkg_id != "" {
		table = table.Where("pkg_id = ?", e.Pkg_id)
	}
	if e.Dtu_type != 0 {
		table = table.Where("pkg_type = ?", e.Dtu_type)
	}

	// 数据权限控制
	dataPermission := new(models.DataPermission)
	dataPermission.UserId, _ = tools.StringToInt(e.DataScope)
	table, err := dataPermission.GetDataScope(e.TableName(), table)
	if err != nil {
		return nil, 0, err
	}
	var count int
	if err := table.Order("bms_spec_info_id").Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
		return nil, 0, err
	}
	table.Where("`deleted_at` IS NULL").Count(&count)
	return doc, count, nil
}
