package batterymanage

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-admin/models/batterymanage"
	"go-admin/tools"
	"go-admin/tools/app"
	"time"
)

//电池列表
// @Summary 电池轨迹数据
// @Description Get JSON
// @Tags 电池轨迹/BatteryMove
// @Param pkg_id query string false "pkg_id"
// @Param dtu_id query string false "dtu_id"
// @Param startTime query Time false "开始时间"
// @Param endTime query Time false "结束时间"
// @Success 200 {object} app.Response "{"code": 200, "data": [...]}"
// @Router /api/bm1/battery/batterylist [get]
// @Security Bearer
func GetBatteryMove(c *gin.Context) {
	var data batterymanage.BatteryMoveInfo
	var err error
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
	}

	//按照json格式
	data.Pkg_id = c.Request.FormValue("pkg_id")
	data.Dtu_id = c.Request.FormValue("dtu_id")

	data.DataScope = tools.GetUserIdStr(c)
	result, _, err := data.GetBatteryMoveInfo(starttime, endtime)
	tools.HasError(err, "", -1)
	app.OK(c, result, "")
}
