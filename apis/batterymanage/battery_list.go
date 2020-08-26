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
// @Tags 电池列表/BatteryList
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
func GetBatteryList(c *gin.Context) {
	var data batterymanage.BatteryListInfo
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
	id := c.Request.FormValue("bms_specInfoId")
	data.Bms_specInfoId, _ = tools.StringToInt(id)

	data.Pkg_id = c.Request.FormValue("pkg_id")

	pkg_type := c.Request.FormValue("pkg_type")
	temp1, _ := tools.StringToInt(pkg_type)
	data.Pkg_type = uint8(temp1)

	pkg_capacity:= c.Request.FormValue("pkg_capacity")
	temp2, _ := tools.StringToInt(pkg_capacity)
	data.Pkg_capacity = uint16(temp2)

	bms_chargeStatus:= c.Request.FormValue("bms_chargeStatus")
	temp3, _ := tools.StringToInt(bms_chargeStatus)
	data.Bms_chargeStatus = uint8(temp3)

	bms_soc:= c.Request.FormValue("bms_soc")
	temp4, _ := tools.StringToInt(bms_soc)
	data.Bms_soc = uint8(temp4)

	bms_errNbr:= c.Request.FormValue("bms_errNbr")
	temp5, _ := tools.StringToInt(bms_errNbr)
	data.Bms_soc = uint8(temp5)

	data.DataScope = tools.GetUserIdStr(c)
	result, count, err := data.GetBatteryListInfo(pageSize, pageIndex)
	tools.HasError(err, "", -1)
	app.PageOK(c, result, count, pageIndex, pageSize, "")
}
// @Summary 删除指定电池
// @Description 删除数据
// @Tags 电池
// @Param id path int true "id"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 500 {string} string	"{"code": 500, "message": "删除失败"}"
// @Router /api/bm1/battery/batterylist/{bms_specinfoId} [delete]
func DelOneBatteryList(c *gin.Context) {
	var data batterymanage.Bms_specInfo
	data.UpdateBy = tools.GetUserIdStr(c)
	ids := tools.IdsStrToIdsIntGroup("bms_specInfoId", c)
	result, err := data.BatchDelete(ids)
	tools.HasError(err, "删除失败", 500)
	app.OK(c, result, "删除成功")
}
