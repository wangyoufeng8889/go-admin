package batterymanage

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-admin/models/batterymanage"
	"go-admin/tools"
	"go-admin/tools/app"
	"time"
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
func GetDTUDetail_dtu_statusinfo(c *gin.Context) {
	var data batterymanage.Dtu_statusInfo
	var err error
	var startdate = time.Now().AddDate(0,0,-1)
	var enddate = time.Now()
	//2006-01-02 15:04:05.9999999 +0800
	if date := c.Request.FormValue("startTime"); date != "" {
		l,err := time.LoadLocation("Local")
		if err != nil {
			fmt.Println(err)
		}
		startdate,err = time.ParseInLocation(TIME_LAYOUT, date, l)
		if err != nil {
			fmt.Println(err)
		}
	}

	if date := c.Request.FormValue("endTime"); date != "" {
		l,err := time.LoadLocation("Local")
		if err != nil {
			fmt.Println(err)
		}
		enddate,err = time.ParseInLocation(TIME_LAYOUT, date, l)
		if err != nil {
			fmt.Println(err)
		}
	}

	id := c.Request.FormValue("dtu_statusInfoId")
	data.Dtu_statusInfoId, _ = tools.StringToInt(id)

	data.Dtu_id = c.Request.FormValue("dtu_id")
	data.Pkg_id = c.Request.FormValue("pkg_id")
	var is_oneList string = c.Request.FormValue("is_oneList")

	data.DataScope = tools.GetUserIdStr(c)
	result,count, err := data.Getdtu_statusinfo(startdate, enddate,is_oneList)
	fmt.Println(count)
	tools.HasError(err, "", -1)
	app.OK(c, result, "")
}

func GetDTUDetail_dtu_specinfo(c *gin.Context) {
	var data batterymanage.Dtu_specInfo
	var err error
	//2006-01-02 15:04:05.9999999 +0800
	id := c.Request.FormValue("dtu_statusInfoId")
	data.Dtu_specInfoId, _ = tools.StringToInt(id)

	data.Dtu_id = c.Request.FormValue("dtu_id")
	data.Pkg_id = c.Request.FormValue("pkg_id")
	var is_oneList string = c.Request.FormValue("is_oneList")

	data.DataScope = tools.GetUserIdStr(c)
	result,count, err := data.Getdtu_specinfo(is_oneList)
	fmt.Println(count)
	tools.HasError(err, "", -1)
	app.OK(c, result, "")
}
func GetDTUDetail_dtu_aliyun(c *gin.Context) {
	var data batterymanage.Dtu_aliyun
	var err error
	//2006-01-02 15:04:05.9999999 +0800
	id := c.Request.FormValue("dtu_aliyunId")
	data.Dtu_aliyunId, _ = tools.StringToInt(id)

	data.Dtu_id = c.Request.FormValue("dtu_id")
	var is_oneList string = c.Request.FormValue("is_oneList")

	data.DataScope = tools.GetUserIdStr(c)
	result,count, err := data.GetDtu_aliyun(is_oneList)
	fmt.Println(count)
	tools.HasError(err, "", -1)
	app.OK(c, result, "")
}
