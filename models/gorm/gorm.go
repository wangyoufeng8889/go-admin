package gorm

import (
	"github.com/jinzhu/gorm"
	"go-admin/models"
	"go-admin/models/batterymanage"
	"go-admin/models/tools"
)

func AutoMigrate(db *gorm.DB) error {
	db.SingularTable(true)
	err := db.AutoMigrate(new(models.CasbinRule)).Error
	if err != nil {
		return err
	}
	err = db.AutoMigrate(new(models.SysDept)).Error
	if err != nil {
		return err
	}
	err = db.AutoMigrate(new(models.SysConfig)).Error
	if err != nil {
		return err
	}
	err = db.AutoMigrate(new(tools.SysTables)).Error
	if err != nil {
		return err
	}
	err = db.AutoMigrate(new(tools.SysColumns)).Error
	if err != nil {
		return err
	}
	err = db.AutoMigrate(new(models.Menu)).Error
	if err != nil {
		return err
	}
	err = db.AutoMigrate(new(models.LoginLog)).Error
	if err != nil {
		return err
	}
	err = db.AutoMigrate(new(models.SysOperLog)).Error
	if err != nil {
		return err
	}
	err = db.AutoMigrate(new(models.RoleMenu)).Error
	if err != nil {
		return err
	}
	err = db.AutoMigrate(new(models.SysRoleDept)).Error
	if err != nil {
		return err
	}
	err = db.AutoMigrate(new(models.SysUser)).Error
	if err != nil {
		return err
	}
	err = db.AutoMigrate(new(models.SysRole)).Error
	if err != nil {
		return err
	}
	err = db.AutoMigrate(new(models.Post)).Error
	if err != nil {
		return err
	}
	err = db.AutoMigrate(new(models.DictData)).Error
	if err != nil {
		return err
	}
	err = db.AutoMigrate(new(models.DictType)).Error
	if err != nil {
		return err
	}
	err = db.AutoMigrate(new(models.SysJob)).Error
	if err != nil {
		return err
	}
	err = db.AutoMigrate(new(models.SysConfig)).Error
	if err != nil {
		return err
	}
	err = db.AutoMigrate(new(models.SysSetting)).Error
	if err != nil {
		return err
	}

	//增加电池数据库迁移
	err = db.AutoMigrate(new(batterymanage.Bms_statusInfo)).Error
	if err != nil {
		return err
	}
	err = db.AutoMigrate(new(batterymanage.Bms_statusInfoLog)).Error
	if err != nil {
		return err
	}
	err = db.AutoMigrate(new(batterymanage.Bms_cellInfo)).Error
	if err != nil {
		return err
	}
	err = db.AutoMigrate(new(batterymanage.Bms_cellInfoLog)).Error
	if err != nil {
		return err
	}
	err = db.AutoMigrate(new(batterymanage.Bms_specInfo)).Error
	if err != nil {
		return err
	}
	err = db.AutoMigrate(new(batterymanage.Bms_historyInfo)).Error
	if err != nil {
		return err
	}
	err = db.AutoMigrate(new(batterymanage.Bms_paraSetReg)).Error
	if err != nil {
		return err
	}
	err = db.AutoMigrate(new(batterymanage.Bms_temperatureInfo)).Error
	if err != nil {
		return err
	}
	err = db.AutoMigrate(new(batterymanage.Bms_temperatureInfoLog)).Error
	if err != nil {
		return err
	}
	err = db.AutoMigrate(new(batterymanage.Dtu_paraSetReg)).Error
	if err != nil {
		return err
	}
	err = db.AutoMigrate(new(batterymanage.Dtu_specInfo)).Error
	if err != nil {
		return err
	}
	err = db.AutoMigrate(new(batterymanage.Dtu_statusInfo)).Error
	if err != nil {
		return err
	}
	err = db.AutoMigrate(new(batterymanage.Dtu_statusInfoLog)).Error
	if err != nil {
		return err
	}
	models.DataInit()
	return err
}

func CustomMigrate(db *gorm.DB, table interface{}) error {
	db.SingularTable(true)
	return db.AutoMigrate(&table).Error
}
