package batterymanage

import (
	"github.com/gin-gonic/gin"
	"go-admin/models/batterymanage"
	"go-admin/tools"
	"go-admin/tools/app"
)

//电池列表
// @Summary DTU列表数据
// @Description Get JSON
// @Tags DTU列表/DtuList
// @Param dtu_id query string false "dtu_id"
// @Param pkg_id query string false "pkg_id"
// @Param dtu_type query uint8 false "dtu_type"
// @Param dtu_setupType query uint8 false "dtu_setupType"
// @Param dtu_aliyunStatus query uint8 false "dtu_aliyunStatus"
// @Param dtu_csq query uint8 false "dtu_csq"
// @Param dtu_errNbr query uint8 false "dtu_errNbr"

// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} app.Response "{"code": 200, "data": [...]}"
// @Router /api/bm1/battery/batterylist [get]
// @Security Bearer
func GetDtuList(c *gin.Context) {
	var data batterymanage.DtuListInfo
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
	id := c.Request.FormValue("dtu_specInfoId")
	data.Dtu_specInfoId, _ = tools.StringToInt(id)

	data.Dtu_id = c.Request.FormValue("pkg_id")
	data.Pkg_id = c.Request.FormValue("pkg_id")

	dtu_type := c.Request.FormValue("dtu_type")
	temp1, _ := tools.StringToInt(dtu_type)
	data.Dtu_type = uint8(temp1)

	dtu_setupType := c.Request.FormValue("dtu_setupType")
	temp2, _ := tools.StringToInt(dtu_setupType)
	data.Dtu_setupType = uint8(temp2)

	dtu_aliyunStatus := c.Request.FormValue("dtu_aliyunStatus")
	temp3, _ := tools.StringToInt(dtu_aliyunStatus)
	data.Dtu_aliyunStatus = uint8(temp3)

	dtu_csq := c.Request.FormValue("dtu_csq")
	temp4, _ := tools.StringToInt(dtu_csq)
	data.Dtu_csq = uint8(temp4)

	dtu_errNbr := c.Request.FormValue("dtu_errNbr")
	temp5, _ := tools.StringToInt(dtu_errNbr)
	data.Dtu_errNbr = uint8(temp5)

	data.DataScope = tools.GetUserIdStr(c)
	result, count, err := data.Getdtu_listinfo(pageSize, pageIndex)
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
func DelOneDtuList(c *gin.Context) {
	var data batterymanage.Bms_specInfo
	data.UpdateBy = tools.GetUserIdStr(c)
	ids := tools.IdsStrToIdsIntGroup("bms_specInfoId", c)
	result, err := data.BatchDelete(ids)
	tools.HasError(err, "删除失败", 500)
	app.OK(c, result, "删除成功")
}

