package batterymanage

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-admin/models/batterymanage"
	"go-admin/tools"
	"go-admin/tools/app"
	"time"
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
	result, err := data.GetDtuDetailInfo()
	tools.HasError(err, "", -1)
	app.OK(c, result, "")
}

//电池soc
// @Summary 电池SOC数据
// @Description Get JSON
// @Tags 电池SOC/BatterySOC
// @Param pkg_id query string false "pkg_id"
// @Param dtu_id query string false "dtu_id"
// @Param startTime query Time false "开始时间"
// @Param endTime query Time false "结束时间"
// @Success 200 {object} app.Response "{"code": 200, "data": [...]}"
// @Router /api/bm1/battery/batterysoc[get]
// @john wang
func GetDtuCSQ(c *gin.Context) {
	var data batterymanage.DtuCSQInfo
	var err error
	var dateflag int = 0
	var starttime time.Time = time.Now().AddDate(0,0,-1)
	var endtime time.Time = time.Now()
	if date := c.Request.FormValue("startTime"); date != "" {
		l,err := time.LoadLocation("Local")
		if err != nil {
			fmt.Println(err)
		}
		starttime,err = time.ParseInLocation(TIME_LAYOUT, date, l)
		if err != nil {
			fmt.Println(err)
		}
		dateflag += 1
	}

	if date := c.Request.FormValue("endTime"); date != "" {
		l,err := time.LoadLocation("Local")
		if err != nil {
			fmt.Println(err)
		}
		endtime,err = time.ParseInLocation(TIME_LAYOUT, date, l)
		if err != nil {
			fmt.Println(err)
		}
		dateflag += 1
	}
	if dateflag == 2{
		dateflag = 1
	}
	//按照json格式
	data.Pkg_id = c.Request.FormValue("pkg_id")
	data.Dtu_id = c.Request.FormValue("dtu_id")

	data.DataScope = tools.GetUserIdStr(c)
	result, _, err := data.GetDtuCSQInfo(starttime, endtime,dateflag)
	tools.HasError(err, "", -1)
	app.OK(c, result, "")
}
