package batterymanage

import (
	"github.com/gin-gonic/gin"
	"go-admin/models/batterymanage"
	"go-admin/tools"
	"go-admin/tools/app"
)

//电池列表
// @Summary 电池列表数据
// @Description Get JSON
// @Tags DTU BMS 绑定关系列表/GetDtuBmsBandList
// @Param pkg_id query string false "pkg_id"
// @Param pkg_type query uint8 false "pkg_type"
// @Param pkg_capacity query uint16 false "pkg_capacity"
// @Param bms_chargeStatus query uint8 false "bms_chargeStatus"
// @Param bms_soc query uint8 false "bms_soc"
// @Param bms_errNbr query uint8 false "bms_errNbr"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} app.Response "{"code": 200, "data": [...]}"
// @Router /api/bm1/battery/batterylist [get]
// @john wang
func GetDtuBmsBandList(c *gin.Context) {
	var data batterymanage.DtuBms_BandInfoLog
	var err error
	var pageSize = 10
	var pageIndex = 1

	if size := c.Request.FormValue("pageSize"); size != "" {
		pageSize = tools.StrToInt(err, size)
	}

	if index := c.Request.FormValue("pageIndex"); index != "" {
		pageIndex = tools.StrToInt(err, index)
	}
	//按照json格式
	id := c.Request.FormValue("dtuBms_BandInfoLogId")
	data.DtuBms_BandInfoLogId, _ = tools.StringToInt(id)

	data.Pkg_id = c.Request.FormValue("pkg_id")
	data.Dtu_id = c.Request.FormValue("dtu_id")

	data.DataScope = tools.GetUserIdStr(c)
	result, count, err := data.GetDtuBmsBandListInfo(pageSize, pageIndex)
	tools.HasError(err, "", -1)
	app.PageOK(c, result, count, pageIndex, pageSize, "")
}