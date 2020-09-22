package batterymanage

import (
	"errors"
	orm "go-admin/global"
	"go-admin/models"
	"go-admin/tools"
)

type Ota_firmware struct {
	Ota_firmwareId int    `json:"ota_firmwareId" gorm:"primary_key;AUTO_INCREMENT"`
	FirmwareName   string `json:"firmwareName" gorm:"type:varchar(256);"`
	FirmwareVer      string `json:"firmwareVer" gorm:"type:varchar(10);"`
	CoreVer      string `json:"coreVer" gorm:"type:varchar(10);"`
	FileName      string `json:"fileName" gorm:"type:varchar(256);"`
	Remark      string `json:"remark" gorm:"type:varchar(256);"`


	DataScope  string `json:"dataScope" gorm:"-"`
	CreateBy  string `gorm:"size:128;" json:"createBy"`
	UpdateBy  string `gorm:"size:128;" json:"updateBy"`
	models.BaseModel
}
func (Ota_firmware) TableName() string {
	return "user_ota_firmware"
}
func (e *Ota_firmware) GetFirmwareListInfo(pageSize int, pageIndex int) ([]Ota_firmware,int, error) {
	var doc []Ota_firmware
	table := orm.Eloquent.Select("*").Table(e.TableName())
	if e.Ota_firmwareId != 0 {
		table = table.Where("ota_firmware_id = ?", e.Ota_firmwareId)
	}
	if e.FirmwareName != "" {
		table = table.Where("firmwareName = ?", e.FirmwareName)
	}
	if e.FirmwareVer != "" {
		table = table.Where("firmware_ver = ?", e.FirmwareVer)
	}
	// 数据权限控制
	dataPermission := new(models.DataPermission)
	dataPermission.UserId, _ = tools.StringToInt(e.DataScope)
	table, err := dataPermission.GetDataScope(e.TableName(), table)
	if err != nil {
		return nil, 0, err
	}
	var count int
	if err := table.Order("ota_firmware_id").Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
		return nil, 0, err
	}
	table.Where("user_ota_firmware.deleted_at IS NULL").Count(&count)
	return doc, count, nil
}
func (e *Ota_firmware) BatchDelete(id []int) (Result bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("ota_firmware_id in (?)", id).Delete(&Ota_firmware{}).Error; err != nil {
		return
	}
	Result = true
	return
}



func (e *Ota_firmware) Create() (Ota_firmware, error) {
	var doc Ota_firmware
	i := 0
	orm.Eloquent.Table(e.TableName()).Where("firmware_name=? and firmware_ver=? and core_ver = ?", e.FirmwareName, e.FirmwareVer,e.CoreVer).Count(&i)
	if i > 0 {
		return doc, errors.New("固件已经存在！")
	}

	result := orm.Eloquent.Table(e.TableName()).Create(&e)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *e
	return doc, nil
}
