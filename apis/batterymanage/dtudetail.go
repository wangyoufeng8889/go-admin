package batterymanage

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-admin/models/batterymanage"
	"go-admin/tools"
	"go-admin/tools/app"
)
// @Summary DTU详情数据
// @Description 获取JSON
// @Tags 电池详情
// @Param Pkg_id query string false "Pkg_id"
// @Success 200 {object} app.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/post [get]
// @john wang
func GetDtuDetail(c *gin.Context) {
	var data batterymanage.DtuDetailInfo
	var err error
	//2006-01-02 15:04:05.9999999 +0800
	id := c.Request.FormValue("dtu_specInfoId")
	data.Dtu_specInfoId, _ = tools.StringToInt(id)

	data.Dtu_id = c.Request.FormValue("dtu_id")

	data.DataScope = tools.GetUserIdStr(c)
	result,count, err := data.GetDtuDetailInfo()
	fmt.Println(count)
	tools.HasError(err, "", -1)
	app.OK(c, result, "")
}

