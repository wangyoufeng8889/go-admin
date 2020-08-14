package batterymanage

import (
	"github.com/gin-gonic/gin"
	"go-admin/models/batterymanage"
	"go-admin/tools"
	"go-admin/tools/app"
)

// @Summary 电池列表数据
// @Description 获取JSON
// @Tags 岗位
// @Param Pkg_id query string false "Pkg_id"
// @Param Dtu_id query string false "Dtu_id"
// @Param Battery_listId query string false "Battery_listId"
// @Param Bms_chargeStatus query string false "Bms_chargeStatus"
// @Success 200 {object} app.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/post [get]
// @Security Bearer
func GetBatteryList(c *gin.Context) {
	var data batterymanage.Battery_list
	var err error
	var pageSize = 10
	var pageIndex = 1

	if size := c.Request.FormValue("pageSize"); size != "" {
		pageSize = tools.StrToInt(err, size)
	}

	if index := c.Request.FormValue("pageIndex"); index != "" {
		pageIndex = tools.StrToInt(err, index)
	}

	id := c.Request.FormValue("bttery_listId")
	data.Battery_listId, _ = tools.StringToInt(id)

	data.Dtu_id = c.Request.FormValue("dtu_id")
	data.Pkg_id = c.Request.FormValue("pkg_id")
	//data.Bms_chargeStatus = c.Request.FormValue("bms_chargeStatus")

	data.DataScope = tools.GetUserIdStr(c)
	result, count, err := data.GetPage(pageSize, pageIndex)
	tools.HasError(err, "", -1)
	app.PageOK(c, result, count, pageIndex, pageSize, "")
}
// @Summary 删除指定电池
// @Description 删除数据
// @Tags 岗位
// @Param id path int true "id"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 500 {string} string	"{"code": 500, "message": "删除失败"}"
// @Router /api/v1/post/{postId} [delete]
func DelOneBatteryList(c *gin.Context) {
	var data batterymanage.Battery_list
	data.UpdateBy = tools.GetUserIdStr(c)
	ids := tools.IdsStrToIdsIntGroup("battery_listId", c)
	result, err := data.BatchDelete(ids)
	tools.HasError(err, "删除失败", 500)
	app.OK(c, result, "删除成功")
}
