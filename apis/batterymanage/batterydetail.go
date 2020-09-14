package batterymanage

import (
	"github.com/gin-gonic/gin"
	"go-admin/models/batterymanage"
	"go-admin/tools"
	"go-admin/tools/app"
)
var TIME_LAYOUT = "2006-01-02 15:04:05"
// @Summary 电池详情数据
// @Description 获取JSON
// @Tags 电池详情
// @Param Pkg_id query string false "Pkg_id"
// @Success 200 {object} app.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/post [get]
// @john wang
func GetBatteryDetail(c *gin.Context) {
	var data batterymanage.BatteryDetailInfo
	var err error
	//2006-01-02 15:04:05.9999999 +0800
	id := c.Request.FormValue("bms_specInfoId")
	data.Bms_specInfoId, _ = tools.StringToInt(id)

	data.Pkg_id = c.Request.FormValue("pkg_id")

	data.DataScope = tools.GetUserIdStr(c)
	result, err := data.GetBatteryDetailInfo()
	tools.HasError(err, "", -1)
	app.OK(c, result, "")
}
