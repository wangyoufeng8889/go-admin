package batterymanage

import (
	"errors"
	orm "go-admin/global"
	"go-admin/models"
	"go-admin/tools"
	"time"
)
type Bms_specInfo struct {
	Bms_specInfoId     int    `json:"bms_specInfoId" gorm:"size:10;primary_key;AUTO_INCREMENT"`
	Dtu_uptime time.Time  `json:"dtu_uptime"`
	Pkg_id   string `json:"pkg_id" gorm:"size:20;primary_key;unique;not null;"`
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
func (Bms_specInfo) TableName() string {
	return "user_bms_specinfo"
}
type BatteryListInfo struct {
	//Bms_specInfo
	Bms_specInfoId     int    `json:"bms_specInfoId" gorm:"size:10;primary_key;AUTO_INCREMENT"`
	Pkg_id   string `json:"pkg_id" gorm:"size:20;primary_key;unique;not null;"`
	Pkg_count   uint8    `json:"pkg_count" gorm:"Type：uint8"`
	Pkg_type   uint8    `json:"pkg_type" gorm:"Type：uint8"`
	Pkg_capacity   uint16    `json:"pkg_capacity" gorm:"Type：uint16"`
	Pkg_nominalVoltage   uint16    `json:"pkg_nominalVoltage" gorm:"Type：uint16"`

	//Bms_statusInfo
	Dtu_uptime time.Time  `json:"dtu_uptime"`
	Bms_chargeStatus      uint8 `json:"bms_chargeStatus" gorm:"Type：uint8"`
	Bms_soc   uint8    `json:"bms_soc" gorm:"Type：uint8"`
	Bms_errNbr   uint8    `json:"bms_errNbr" gorm:"Type：uint8"`
	Bms_voltage   uint16    `json:"bms_voltage" gorm:"Type：uint16"`
	//无关数据库
	DataScope  string `json:"dataScope" gorm:"-"`
	models.BaseModel

}
func (BatteryListInfo) TableName() string {
	return "user_bms_specinfo"
}
//电池列表
func (e *BatteryListInfo) GetBatteryListInfo(pageSize int, pageIndex int) ([]BatteryListInfo,int, error) {
	var doc []BatteryListInfo

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
	if e.Bms_specInfoId != 0 {
		table = table.Where("bms_spec_info_id = ?", e.Bms_specInfoId)
	}
	if e.Pkg_id != "" {
		table = table.Where("pkg_id = ?", e.Pkg_id)
	}
	if e.Pkg_type != 0 {
		table = table.Where("pkg_type = ?", e.Pkg_type)
	}
	if e.Pkg_capacity != 0 {
		table = table.Where("pkg_capacity = ?", e.Pkg_capacity)
	}
	if e.Bms_chargeStatus != 0 {
		table = table.Where("bms_charge_status = ?", e.Bms_chargeStatus)
	}
	if e.Bms_soc != 0 {
		table = table.Where("bms_soc > ?", e.Bms_soc)
	}
	if e.Bms_errNbr != 0 {
		table = table.Where("bms_err_nbr > ?", e.Bms_errNbr)
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
func (e *Bms_specInfo) BatchDelete(id []int) (Result bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("bms_spec_info_id in (?)", id).Delete(&Bms_specInfo{}).Error; err != nil {
		return
	}
	Result = true
	return
}
type BatteryDetailInfo struct {
	//Bms_specInfo
	Bms_specInfoId     int    `json:"bms_specInfoId" gorm:"size:10;primary_key;AUTO_INCREMENT"`
	Pkg_id   string `json:"pkg_id" gorm:"size:20;primary_key;unique;not null;"`
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

	//Bms_statusInfo
	Dtu_uptime time.Time  `json:"dtu_uptime"`
	Bms_chargeStatus      uint8 `json:"bms_chargeStatus" gorm:"Type：uint8"`
	Bms_soc   uint8    `json:"bms_soc" gorm:"Type：uint8"`
	Bms_errStatus   uint8    `json:"bms_errStatus" gorm:"Type：uint8"`
	Bms_errNbr   uint8    `json:"bms_errNbr" gorm:"Type：uint8"`
	Bms_errCode   uint32    `json:"bms_errCode" gorm:"Type：uint32"`
	Bms_voltage   uint16    `json:"bms_voltage" gorm:"Type：uint16"`
	Bms_current	  uint16  `json:"bms_current" gorm:"Type：uint16"`
	Bms_maxCellVoltage   uint16    `json:"bms_maxCellVoltage" gorm:"Type：uint16"`
	Bms_minCellVoltage   uint16    `json:"bms_minCellVoltage" gorm:"Type：uint16"`
	Bms_averageCellVoltage   uint16    `json:"bms_averageCellVoltage" gorm:"Type：uint16"`
	Bms_maxTemperature   uint8    `json:"bms_maxTemperature" gorm:"Type：uint8"`
	Bms_minTemperature   uint8    `json:"bms_minTemperature" gorm:"Type：uint8"`
	Bms_mosTemperature   uint8    `json:"bms_mosTemperature" gorm:"Type：uint8"`
	Bms_balanceResistance   uint8    `json:"bms_balanceResistance" gorm:"Type：uint8"`
	Bms_chargeMosStatus   uint8    `json:"bms_chargeMosStatus" gorm:"Type：uint8"`
	Bms_dischargeMosStatus   uint8    `json:"bms_dischargeMosStatus" gorm:"Type：uint8"`
	Bms_otaBufStatus   uint8    `json:"bms_otaBufStatus" gorm:"Type：uint8"`
	Bms_magneticCheck   uint8   `json:"bms_magneticCheck" gorm:"Type：uint8"`

	//Bms_cellInfo
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

	//Bms_paraSetReg
	Bms_chargeMosCtr   uint8    `json:"bms_chargeMosCtr" gorm:"Type：uint8"`
	Bms_dischargeMosCtr   uint8    `json:"bms_dischargeMosCtr" gorm:"Type：uint8"`
	Bms_chargeHighTempProtect   uint8    `json:"bms_chargeHighTempProtect" gorm:"Type：uint8"`
	Bms_chargeHighTempRelease   uint8    `json:"bms_chargeHighTempRelease" gorm:"Type：uint8"`
	Bms_chargeLowTempProtect   uint8    `json:"bms_chargeLowTempProtect" gorm:"Type：uint8"`
	Bms_chargeLowTempRelease   uint8    `json:"bms_chargeLowTempRelease" gorm:"Type：uint8"`
	Bms_chargeHighTempDelay   uint8    `json:"bms_chargeHighTempDelay" gorm:"Type：uint8"`
	Bms_chargeLowTempDelay   uint8    `json:"bms_chargeLowTempDelay" gorm:"Type：uint8"`
	Bms_dischargeHighTempProtect   uint8    `json:"bms_dischargeHighTempProtect" gorm:"Type：uint8"`
	Bms_dischargeHighTempRelease   uint8    `json:"bms_dischargeHighTempRelease" gorm:"Type：uint8"`
	Bms_dischargeLowTempProtect   uint8    `json:"bms_dischargeLowTempProtect" gorm:"Type：uint8"`
	Bms_dischargeLowTempRelease   uint8    `json:"bms_dischargeLowTempRelease" gorm:"Type：uint8"`
	Bms_dischargeHighTempDelay   uint8    `json:"bms_dischargeHighTempDelay" gorm:"Type：uint8"`
	Bms_dischargeLowTempDelay   uint8    `json:"bms_dischargeLowTempDelay" gorm:"Type：uint8"`
	Bms_mosHighTempProtect   uint8    `json:"bms_mosHighTempProtect" gorm:"Type：uint8"`
	Bms_mosHighTempRelease   uint8    `json:"bms_mosHighTempRelease" gorm:"Type：uint8"`
	Bms_pkgOverVoltageProtect   uint16    `json:"bms_pkgOverVoltageProtect" gorm:"Type：uint16"`
	Bms_pkgOverVoltageRelease   uint16    `json:"bms_pkgOverVoltageRelease" gorm:"Type：uint16"`
	Bms_pkgUnderVoltageProtect   uint16    `json:"bms_pkgUnderVoltageProtect" gorm:"Type：uint16"`
	Bms_pkgUnderVoltageRelease   uint16    `json:"bms_pkgUnderVoltageRelease" gorm:"Type：uint16"`
	Bms_pkgUnderVoltageDelay   uint8    `json:"bms_pkgUnderVoltageDelay" gorm:"Type：uint8"`
	Bms_pkgOverVoltageDelay   uint8    `json:"bms_pkgOverVoltageDelay" gorm:"Type：uint8"`
	Bms_cellOverVoltageProtect   uint16    `json:"bms_cellOverVoltageProtect" gorm:"Type：uint16"`
	Bms_cellOverVoltageRelease   uint16    `json:"bms_cellOverVoltageRelease" gorm:"Type：uint16"`
	Bms_cellUnderVoltageProtect   uint16    `json:"bms_cellUnderVoltageProtect" gorm:"Type：uint16"`
	Bms_cellUnderVoltageRelease   uint16    `json:"bms_cellUnderVoltageRelease" gorm:"Type：uint16"`
	Bms_cellUnderVoltageDelay   uint8    `json:"bms_cellUnderVoltageDelay" gorm:"Type：uint8"`
	Bms_cellOverVoltageDelay   uint8    `json:"bms_cellOverVoltageDelay" gorm:"Type：uint8"`
	Bms_chargeOverCurrentProtect   uint16    `json:"bms_chargeOverCurrentProtect" gorm:"Type：uint16"`
	Bms_chargeOverCurrentDelay   uint8    `json:"bms_chargeOverCurrentDelay" gorm:"Type：uint8"`
	Bms_chargeOverCurrentRelease   uint8    `json:"bms_chargeOverCurrentRelease" gorm:"Type：uint8"`
	Bms_dischargeOverCurrentProtect   uint16    `json:"bms_dischargeOverCurrentProtect" gorm:"Type：uint16"`
	Bms_dischargeOverCurrentDelay   uint8    `json:"bms_dischargeOverCurrentDelay" gorm:"Type：uint8"`
	Bms_dischargeOverCurrentRelease   uint8    `json:"bms_dischargeOverCurrentRelease" gorm:"Type：uint8"`
	Bms_balanceOpenVoltage   uint16    `json:"bms_balanceOpenVoltage" gorm:"Type：uint16"`
	Bms_balanceVoltageDiff   uint16    `json:"bms_balanceVoltageDiff" gorm:"Type：uint16"`
	Bms_balanceTime   uint16    `json:"bms_balanceTime" gorm:"Type：uint16"`
	Bms_funcConfig   uint16    `json:"bms_funcConfig" gorm:"Type：uint16"`
	Bms_hardCellOverVoltage   uint16    `json:"bms_hardCellOverVoltage" gorm:"Type：uint16"`
	Bms_hardCellUnderVoltage   uint16    `json:"bms_hardCellUnderVoltage" gorm:"Type：uint16"`
	Bms_hardOverCurrentAndDelayTime   uint16    `json:"bms_hardOverCurrentAndDelayTime" gorm:"Type：uint16"`
	Bms_hardUnderVoltageAndOverCurrentDelayTime   uint16    `json:"bms_hardUnderVoltageAndOverCurrentDelayTime" gorm:"Type：uint16"`
	Bms_magneticCheckEnable   uint16    `json:"bms_magneticCheckEnable" gorm:"Type：uint16"`
	Bms_forceIntoStorageMode   uint16    `json:"bms_forceIntoStorageMode" gorm:"Type：uint16"`
	Bms_enableChargeStatus   uint16    `json:"bms_enableChargeStatus" gorm:"Type：uint16"`

	//Bms_temperatureInfo
	Bms_temperature1   uint8    `json:"bms_temperature1" gorm:"Type：uint8"`
	Bms_temperature2   uint8    `json:"bms_temperature2" gorm:"Type：uint8"`
	Bms_temperature3   uint8    `json:"bms_temperature3" gorm:"Type：uint8"`
	Bms_temperature4   uint8    `json:"bms_temperature4" gorm:"Type：uint8"`
	Bms_temperature5   uint8    `json:"bms_temperature5" gorm:"Type：uint8"`
	Bms_temperature6   uint8    `json:"bms_temperature6" gorm:"Type：uint8"`

	//Dtu_specInfo
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
	Dtu_aliyunStatus uint8    `json:"dtu_aliyunStatus" gorm:"Type：uint8"`


	//Dtu_statusInfo
	Dtu_latitudeType   string    `json:"dtu_latitudeType" gorm:"Type：size:2"`
	Dtu_longitudeType   string    `json:"dtu_longitudeType" gorm:"Type：size:2"`
	Dtu_latitude   string    `json:"dtu_latitude" gorm:"Type：size:20"`
	Dtu_longitude   string    `json:"dtu_longitude" gorm:"Type：size:20"`
	Dtu_csq   uint8    `json:"dtu_csq" gorm:"Type：uint8"`
	Dtu_locateMode   uint8    `json:"dtu_locateMode" gorm:"Type：uint8"`
	Dtu_gpsSateCnt   uint8    `json:"dtu_gpsSateCnt" gorm:"Type：uint8"`
	//无关数据库
	DataScope  string `json:"dataScope" gorm:"-"`
	models.BaseModel
}


func (BatteryDetailInfo) TableName() string {
	return "user_bms_specinfo"
}
//电池列表
func (e *BatteryDetailInfo) GetBatteryDetailInfo() ([]BatteryDetailInfo,int, error) {
	if e.Pkg_id == "" {
		return nil, 0, errors.New("no pkgid")
	}
	var doc []BatteryDetailInfo
	table := orm.Eloquent.Table(e.TableName())
	var dtu_pkg_bind Dtu_specInfo
	if err:= orm.Eloquent.Table("user_dtu_specinfo").Where("user_dtu_specinfo.pkg_id = ?", e.Pkg_id).First(&dtu_pkg_bind).Error;err!=nil{
		//找不到对应的dtu
		table = table.Select([]string{"user_bms_specinfo.bms_spec_info_id",
			"user_bms_specinfo.pkg_id",
			"user_bms_specinfo.bms_id",
			"user_bms_specinfo.pkg_count",
			"user_bms_specinfo.pkg_type",
			"user_bms_specinfo.pkg_capacity",
			"user_bms_specinfo.pkg_nominal_voltage",
			"user_bms_specinfo.pkg_ntc_count",
			"user_bms_specinfo.pkg_manufacture_date",
			"user_bms_specinfo.bms_hard_ver",
			"user_bms_specinfo.bms_soft_ver",
			"user_bms_specinfo.bms_protocol_ver",
			"user_bms_statusinfo.dtu_uptime",
			"user_bms_statusinfo.bms_charge_status",
			"user_bms_statusinfo.bms_soc",
			"user_bms_statusinfo.bms_err_status",
			"user_bms_statusinfo.bms_err_nbr",
			"user_bms_statusinfo.bms_err_code",
			"user_bms_statusinfo.bms_voltage",
			"user_bms_statusinfo.bms_current",
			"user_bms_statusinfo.bms_max_cell_voltage",
			"user_bms_statusinfo.bms_min_cell_voltage",
			"user_bms_statusinfo.bms_average_cell_voltage",
			"user_bms_statusinfo.bms_max_temperature",
			"user_bms_statusinfo.bms_min_temperature",
			"user_bms_statusinfo.bms_mos_temperature",
			"user_bms_statusinfo.bms_balance_resistance",
			"user_bms_statusinfo.bms_charge_mos_status",
			"user_bms_statusinfo.bms_discharge_mos_status",
			"user_bms_statusinfo.bms_ota_buf_status",
			"user_bms_statusinfo.bms_magnetic_check",
			"user_bms_cellinfo.bms_cell_voltage1",
			"user_bms_cellinfo.bms_cell_voltage2",
			"user_bms_cellinfo.bms_cell_voltage3",
			"user_bms_cellinfo.bms_cell_voltage4",
			"user_bms_cellinfo.bms_cell_voltage5",
			"user_bms_cellinfo.bms_cell_voltage6",
			"user_bms_cellinfo.bms_cell_voltage7",
			"user_bms_cellinfo.bms_cell_voltage8",
			"user_bms_cellinfo.bms_cell_voltage9",
			"user_bms_cellinfo.bms_cell_voltage10",
			"user_bms_cellinfo.bms_cell_voltage11",
			"user_bms_cellinfo.bms_cell_voltage12",
			"user_bms_cellinfo.bms_cell_voltage13",
			"user_bms_cellinfo.bms_cell_voltage14",
			"user_bms_cellinfo.bms_cell_voltage15",
			"user_bms_cellinfo.bms_cell_voltage16",
			"user_bms_cellinfo.bms_cell_voltage17",
			"user_bms_cellinfo.bms_cell_voltage18",
			"user_bms_cellinfo.bms_cell_voltage19",
			"user_bms_cellinfo.bms_cell_voltage20",
			"user_bms_temperatureinfo.bms_temperature1",
			"user_bms_temperatureinfo.bms_temperature2",
			"user_bms_temperatureinfo.bms_temperature3",
			"user_bms_temperatureinfo.bms_temperature4",
			"user_bms_temperatureinfo.bms_temperature5",
			"user_bms_temperatureinfo.bms_temperature6",
			"user_bms_historyinfo.pkg_history_max_cell_voltage",
			"user_bms_historyinfo.pkg_history_min_cell_voltage",
			"user_bms_historyinfo.pkg_history_max_voltage_dif",
			"user_bms_historyinfo.pkg_history_max_temperature",
			"user_bms_historyinfo.pkg_history_min_temperature",
			"user_bms_historyinfo.pkg_history_max_discharge_current",
			"user_bms_historyinfo.pkg_history_max_charge_current",
			"user_bms_historyinfo.pkg_nbr_of_charge_discharge",
			"user_bms_historyinfo.pkg_nbrof_charging_cycle",
			"user_bms_parasetreg.bms_charge_mos_ctr",
			"user_bms_parasetreg.bms_discharge_mos_ctr",
			"user_bms_parasetreg.bms_charge_high_temp_protect",
			"user_bms_parasetreg.bms_charge_high_temp_release",
			"user_bms_parasetreg.bms_charge_low_temp_protect",
			"user_bms_parasetreg.bms_charge_low_temp_release",
			"user_bms_parasetreg.bms_charge_high_temp_delay",
			"user_bms_parasetreg.bms_charge_low_temp_delay",
			"user_bms_parasetreg.bms_discharge_high_temp_protect",
			"user_bms_parasetreg.bms_discharge_high_temp_release",
			"user_bms_parasetreg.bms_discharge_low_temp_protect",
			"user_bms_parasetreg.bms_discharge_low_temp_release",
			"user_bms_parasetreg.bms_discharge_high_temp_delay",
			"user_bms_parasetreg.bms_discharge_low_temp_delay",
			"user_bms_parasetreg.bms_mos_high_temp_protect",
			"user_bms_parasetreg.bms_mos_high_temp_release",
			"user_bms_parasetreg.bms_pkg_over_voltage_protect",
			"user_bms_parasetreg.bms_pkg_over_voltage_release",
			"user_bms_parasetreg.bms_pkg_under_voltage_protect",
			"user_bms_parasetreg.bms_pkg_under_voltage_release",
			"user_bms_parasetreg.bms_pkg_under_voltage_delay",
			"user_bms_parasetreg.bms_pkg_over_voltage_delay",
			"user_bms_parasetreg.bms_cell_over_voltage_protect",
			"user_bms_parasetreg.bms_cell_over_voltage_release",
			"user_bms_parasetreg.bms_cell_under_voltage_protect",
			"user_bms_parasetreg.bms_cell_under_voltage_release",
			"user_bms_parasetreg.bms_cell_under_voltage_delay",
			"user_bms_parasetreg.bms_cell_over_voltage_delay",
			"user_bms_parasetreg.bms_charge_over_current_protect",
			"user_bms_parasetreg.bms_charge_over_current_delay",
			"user_bms_parasetreg.bms_charge_over_current_release",
			"user_bms_parasetreg.bms_discharge_over_current_protect",
			"user_bms_parasetreg.bms_discharge_over_current_delay",
			"user_bms_parasetreg.bms_discharge_over_current_release",
			"user_bms_parasetreg.bms_balance_open_voltage",
			"user_bms_parasetreg.bms_balance_voltage_diff",
			"user_bms_parasetreg.bms_balance_time",
			"user_bms_parasetreg.bms_func_config",
			"user_bms_parasetreg.bms_hard_cell_over_voltage",
			"user_bms_parasetreg.bms_hard_cell_under_voltage",
			"user_bms_parasetreg.bms_hard_over_current_and_delay_time",
			"user_bms_parasetreg.bms_hard_under_voltage_and_over_current_delay_time",
			"user_bms_parasetreg.bms_magnetic_check_enable",
			"user_bms_parasetreg.bms_force_into_storage_mode",
			"user_bms_parasetreg.bms_enable_charge_status"})
		table = table.Joins("LEFT JOIN user_bms_cellinfo ON user_bms_specinfo.pkg_id=user_bms_cellinfo.pkg_id").
			Joins("LEFT JOIN user_bms_historyinfo ON user_bms_specinfo.pkg_id=user_bms_historyinfo.pkg_id").
			Joins("LEFT JOIN user_bms_parasetreg ON user_bms_specinfo.pkg_id=user_bms_parasetreg.pkg_id").
			Joins("LEFT JOIN user_bms_statusinfo ON user_bms_specinfo.pkg_id=user_bms_statusinfo.pkg_id").
			Joins("LEFT JOIN user_bms_temperatureinfo ON user_bms_specinfo.pkg_id=user_bms_temperatureinfo.pkg_id")
	}else {
		table = table.Select([]string{"user_bms_specinfo.bms_spec_info_id",
			"user_bms_specinfo.pkg_id",
			"user_bms_specinfo.dtu_id",
			"user_bms_specinfo.bms_id",
			"user_bms_specinfo.pkg_count",
			"user_bms_specinfo.pkg_type",
			"user_bms_specinfo.pkg_capacity",
			"user_bms_specinfo.pkg_nominal_voltage",
			"user_bms_specinfo.pkg_ntc_count",
			"user_bms_specinfo.pkg_manufacture_date",
			"user_bms_specinfo.bms_hard_ver",
			"user_bms_specinfo.bms_soft_ver",
			"user_bms_specinfo.bms_protocol_ver",
			"user_bms_statusinfo.dtu_uptime",
			"user_bms_statusinfo.bms_charge_status",
			"user_bms_statusinfo.bms_soc",
			"user_bms_statusinfo.bms_err_status",
			"user_bms_statusinfo.bms_err_nbr",
			"user_bms_statusinfo.bms_err_code",
			"user_bms_statusinfo.bms_voltage",
			"user_bms_statusinfo.bms_current",
			"user_bms_statusinfo.bms_max_cell_voltage",
			"user_bms_statusinfo.bms_min_cell_voltage",
			"user_bms_statusinfo.bms_average_cell_voltage",
			"user_bms_statusinfo.bms_max_temperature",
			"user_bms_statusinfo.bms_min_temperature",
			"user_bms_statusinfo.bms_mos_temperature",
			"user_bms_statusinfo.bms_balance_resistance",
			"user_bms_statusinfo.bms_charge_mos_status",
			"user_bms_statusinfo.bms_discharge_mos_status",
			"user_bms_statusinfo.bms_ota_buf_status",
			"user_bms_statusinfo.bms_magnetic_check",
			"user_bms_cellinfo.bms_cell_voltage1",
			"user_bms_cellinfo.bms_cell_voltage2",
			"user_bms_cellinfo.bms_cell_voltage3",
			"user_bms_cellinfo.bms_cell_voltage4",
			"user_bms_cellinfo.bms_cell_voltage5",
			"user_bms_cellinfo.bms_cell_voltage6",
			"user_bms_cellinfo.bms_cell_voltage7",
			"user_bms_cellinfo.bms_cell_voltage8",
			"user_bms_cellinfo.bms_cell_voltage9",
			"user_bms_cellinfo.bms_cell_voltage10",
			"user_bms_cellinfo.bms_cell_voltage11",
			"user_bms_cellinfo.bms_cell_voltage12",
			"user_bms_cellinfo.bms_cell_voltage13",
			"user_bms_cellinfo.bms_cell_voltage14",
			"user_bms_cellinfo.bms_cell_voltage15",
			"user_bms_cellinfo.bms_cell_voltage16",
			"user_bms_cellinfo.bms_cell_voltage17",
			"user_bms_cellinfo.bms_cell_voltage18",
			"user_bms_cellinfo.bms_cell_voltage19",
			"user_bms_cellinfo.bms_cell_voltage20",
			"user_bms_temperatureinfo.bms_temperature1",
			"user_bms_temperatureinfo.bms_temperature2",
			"user_bms_temperatureinfo.bms_temperature3",
			"user_bms_temperatureinfo.bms_temperature4",
			"user_bms_temperatureinfo.bms_temperature5",
			"user_bms_temperatureinfo.bms_temperature6",
			"user_bms_historyinfo.pkg_history_max_cell_voltage",
			"user_bms_historyinfo.pkg_history_min_cell_voltage",
			"user_bms_historyinfo.pkg_history_max_voltage_dif",
			"user_bms_historyinfo.pkg_history_max_temperature",
			"user_bms_historyinfo.pkg_history_min_temperature",
			"user_bms_historyinfo.pkg_history_max_discharge_current",
			"user_bms_historyinfo.pkg_history_max_charge_current",
			"user_bms_historyinfo.pkg_nbr_of_charge_discharge",
			"user_bms_historyinfo.pkg_nbrof_charging_cycle",
			"user_bms_parasetreg.bms_charge_mos_ctr",
			"user_bms_parasetreg.bms_discharge_mos_ctr",
			"user_bms_parasetreg.bms_charge_high_temp_protect",
			"user_bms_parasetreg.bms_charge_high_temp_release",
			"user_bms_parasetreg.bms_charge_low_temp_protect",
			"user_bms_parasetreg.bms_charge_low_temp_release",
			"user_bms_parasetreg.bms_charge_high_temp_delay",
			"user_bms_parasetreg.bms_charge_low_temp_delay",
			"user_bms_parasetreg.bms_discharge_high_temp_protect",
			"user_bms_parasetreg.bms_discharge_high_temp_release",
			"user_bms_parasetreg.bms_discharge_low_temp_protect",
			"user_bms_parasetreg.bms_discharge_low_temp_release",
			"user_bms_parasetreg.bms_discharge_high_temp_delay",
			"user_bms_parasetreg.bms_discharge_low_temp_delay",
			"user_bms_parasetreg.bms_mos_high_temp_protect",
			"user_bms_parasetreg.bms_mos_high_temp_release",
			"user_bms_parasetreg.bms_pkg_over_voltage_protect",
			"user_bms_parasetreg.bms_pkg_over_voltage_release",
			"user_bms_parasetreg.bms_pkg_under_voltage_protect",
			"user_bms_parasetreg.bms_pkg_under_voltage_release",
			"user_bms_parasetreg.bms_pkg_under_voltage_delay",
			"user_bms_parasetreg.bms_pkg_over_voltage_delay",
			"user_bms_parasetreg.bms_cell_over_voltage_protect",
			"user_bms_parasetreg.bms_cell_over_voltage_release",
			"user_bms_parasetreg.bms_cell_under_voltage_protect",
			"user_bms_parasetreg.bms_cell_under_voltage_release",
			"user_bms_parasetreg.bms_cell_under_voltage_delay",
			"user_bms_parasetreg.bms_cell_over_voltage_delay",
			"user_bms_parasetreg.bms_charge_over_current_protect",
			"user_bms_parasetreg.bms_charge_over_current_delay",
			"user_bms_parasetreg.bms_charge_over_current_release",
			"user_bms_parasetreg.bms_discharge_over_current_protect",
			"user_bms_parasetreg.bms_discharge_over_current_delay",
			"user_bms_parasetreg.bms_discharge_over_current_release",
			"user_bms_parasetreg.bms_balance_open_voltage",
			"user_bms_parasetreg.bms_balance_voltage_diff",
			"user_bms_parasetreg.bms_balance_time",
			"user_bms_parasetreg.bms_func_config",
			"user_bms_parasetreg.bms_hard_cell_over_voltage",
			"user_bms_parasetreg.bms_hard_cell_under_voltage",
			"user_bms_parasetreg.bms_hard_over_current_and_delay_time",
			"user_bms_parasetreg.bms_hard_under_voltage_and_over_current_delay_time",
			"user_bms_parasetreg.bms_magnetic_check_enable",
			"user_bms_parasetreg.bms_force_into_storage_mode",
			"user_bms_parasetreg.bms_enable_charge_status",

			"user_dtu_specinfo.dtu_type",
			"user_dtu_specinfo.dtu_setup_type",
			"user_dtu_specinfo.dtu_core_ver",
			"user_dtu_specinfo.dtu_hard_ver",
			"user_dtu_specinfo.dtu_soft_ver",
			"user_dtu_specinfo.dtu_protocol_ver",
			"user_dtu_specinfo.dtu_sim_iccid",
			"user_dtu_specinfo.dtu_imei",
			"user_dtu_specinfo.dtu_bms_bind_status",

			"user_dtu_specinfo.dtu_aliyun_status",

			"user_dtu_statusinfo.dtu_latitude_type",
			"user_dtu_statusinfo.dtu_longitude_type",
			"user_dtu_statusinfo.dtu_latitude",
			"user_dtu_statusinfo.dtu_longitude",
			"user_dtu_statusinfo.dtu_csq",
			"user_dtu_statusinfo.dtu_locate_mode",
			"user_dtu_statusinfo.dtu_gps_sate_cnt"})
		table = table.Joins("LEFT JOIN user_bms_cellinfo ON user_bms_specinfo.pkg_id=user_bms_cellinfo.pkg_id").
			Joins("LEFT JOIN user_bms_historyinfo ON user_bms_specinfo.pkg_id=user_bms_historyinfo.pkg_id").
			Joins("LEFT JOIN user_bms_parasetreg ON user_bms_specinfo.pkg_id=user_bms_parasetreg.pkg_id").
			Joins("LEFT JOIN user_bms_statusinfo ON user_bms_specinfo.pkg_id=user_bms_statusinfo.pkg_id").
			Joins("LEFT JOIN user_bms_temperatureinfo ON user_bms_specinfo.pkg_id=user_bms_temperatureinfo.pkg_id").
			Joins("LEFT JOIN user_dtu_specinfo ON user_bms_specinfo.dtu_id=user_dtu_specinfo.dtu_id").
			Joins("LEFT JOIN user_dtu_statusinfo ON user_bms_specinfo.dtu_id=user_dtu_statusinfo.dtu_id")
	}


	// 数据权限控制
	dataPermission := new(models.DataPermission)
	dataPermission.UserId, _ = tools.StringToInt(e.DataScope)
	table, err := dataPermission.GetDataScope(e.TableName(), table)
	if err != nil {
		return nil, 0, err
	}
	var count int
	table = table.Find(&doc)
	if table.Error!= nil {
		return nil, 0, table.Error
	}
	if e.Bms_specInfoId != 0 {
		table = table.Where("bms_spec_info_id = ?", e.Bms_specInfoId)
	}

	table = table.Where("user_bms_specinfo.pkg_id = ?", e.Pkg_id).Where("user_bms_specinfo.deleted_at IS NULL")
	if err := table.First(&doc).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	return doc, count, nil
}




