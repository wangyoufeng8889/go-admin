package batterymanage

import (
	"github.com/gin-gonic/gin"
	"go-admin/models/batterymanage"
	"go-admin/tools"
	"go-admin/tools/app"
)
//固件列表
// @Summary 固件列表数据
// @Description Get JSON
// @Tags 固件列表/FirmwareList
// @Param firmwareName query string false "firmwareName"
// @Param firmwareVer query uint8 false "firmwareVer"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} app.Response "{"code": 200, "data": [...]}"
// @Router /api/bm1/battery/firmwarelist [get]
// @john wang
func GetFirmwareList(c *gin.Context) {
	var data batterymanage.Ota_firmware
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
	id := c.Request.FormValue("ota_firmwareId")
	data.Ota_firmwareId, _ = tools.StringToInt(id)

	data.FirmwareName = c.Request.FormValue("firmwareName")
	data.FirmwareVer = c.Request.FormValue("firmwareVer")


	data.DataScope = tools.GetUserIdStr(c)
	result, count, err := data.GetFirmwareListInfo(pageSize, pageIndex)
	tools.HasError(err, "", -1)
	app.PageOK(c, result, count, pageIndex, pageSize, "")
}
// @Summary 删除指定固件
// @Description 删除数据
// @Tags 固件
// @Param id path int true "id"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 500 {string} string	"{"code": 500, "message": "删除失败"}"
// @Router /api/bm1/battery/batterylist/{bms_specinfoId} [delete]
func DelOneFirmwareList(c *gin.Context) {
	var data batterymanage.Ota_firmware
	data.UpdateBy = tools.GetUserIdStr(c)
	ids := tools.IdsStrToIdsIntGroup("ota_firmwareId", c)
	result, err := data.BatchDelete(ids)
	tools.HasError(err, "删除失败", 500)
	app.OK(c, result, "删除成功")
}
