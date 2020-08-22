package batterymanage

import (
	"go-admin/models"
	"time"
)
type Bms_paraSetReg struct {
	Bms_paraSetRegId     int    `json:"bms_paraSetRegId" gorm:"size:10;primary_key;AUTO_INCREMENT"`
	Dtu_uptime time.Time  `json:"dtu_uptime"`
	Pkg_id   string `json:"pkg_id" gorm:"size:20;"`
	Dtu_id      string `json:"dtu_id" gorm:"size:20;"`
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
	DataScope  string `json:"dataScope" gorm:"-"`
	UpdateBy  string `gorm:"size:128;" json:"updateBy"`
	models.BaseModel
}
func (Bms_paraSetReg) TableName() string {
	return "user_bms_paraSetReg"
}

