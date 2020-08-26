package batterymanage

import (
	"github.com/gin-gonic/gin"
	"go-admin/models/batterymanage"
	"go-admin/tools"
	"go-admin/tools/app"
)
//电池列表
// @Summary 电池大屏信息
// @Description Get JSON
// @Tags 电池列表/BatteryDashboard
// @Success 200 {object} app.Response "{"code": 200, "data": [...]}"
// @Router /api/bm1/battery/dashboard [get]
// @john wang
func GetBatteryDashboardInfo(c *gin.Context) {
	var data batterymanage.DashboardInfo
	var err error
	data.DataScope = tools.GetUserIdStr(c)
	result, err := data.GetDashboardInfo()
	tools.HasError(err, "", -1)
	app.OK(c, result, "")
}
