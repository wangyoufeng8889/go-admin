package router

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"go-admin/apis/batterymanage"
	"go-admin/middleware"
	"go-admin/pkg/jwtauth"
	jwt "go-admin/pkg/jwtauth"
)

// 路由示例
func InitExamplesRouter(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) *gin.Engine {

	// 无需认证的路由
	examplesNoCheckRoleRouter(r)
	// 需要认证的路由
	examplesCheckRoleRouter(r, authMiddleware)

	return r
}

// 无需认证的路由示例
func examplesNoCheckRoleRouter(r *gin.Engine) {
	// 可根据业务需求来设置接口版本
	v1 := r.Group("/api/v1")
	// 空接口防止v1定义无使用报错
	v1.GET("/nilcheckrole", nil)

	// {{无需认证路由自动补充在此处请勿删除}}
}

// 需要认证的路由示例
func examplesCheckRoleRouter(r *gin.Engine, authMiddleware *jwtauth.GinJWTMiddleware) {
	// 可根据业务需求来设置接口版本,bm1表示batterymanage1
	battery := r.Group("/api/bm1/battery")
	// 空接口防止v1定义无使用报错
	battery.GET("/checkrole", nil)

	// {{认证路由自动补充在此处请勿删除}}
	registerUserBatterylistRouter(battery, authMiddleware)
	registerUserBatterydetailRouter(battery, authMiddleware)
	registerserBatteryMoveRouter(battery, authMiddleware)
	registerUserDTUlistRouter(battery, authMiddleware)
	registerUserDTUdetailRouter(battery, authMiddleware)
	registerUserBatteryDashboardRouter(battery, authMiddleware)
	registerUserdtubmsbandRouter(battery, authMiddleware)
	// 可根据业务需求来设置接口版本,bm1表示batterymanage1
	otaupdate := r.Group("/api/bm1/otaupdate")
	// 空接口防止v1定义无使用报错
	otaupdate.GET("/checkrole", nil)
	registerUserFirmwareListRouter(otaupdate, authMiddleware)

}
func registerUserBatteryDashboardRouter(user *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	battertlist := user.Group("/dashboard").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		battertlist.GET("", batterymanage.GetBatteryDashboardInfo)//电池大屏信息
	}
}
func registerUserBatterylistRouter(user *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	battertlist := user.Group("/batterylist").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		battertlist.GET("", batterymanage.GetBatteryList)//电池列表
		battertlist.DELETE("/:bms_specInfoId", batterymanage.DelOneBatteryList)//删除电池
	}
}
func registerUserBatterydetailRouter(user *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	battertlist := user.Group("/batterydetail").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		battertlist.GET("", batterymanage.GetBatteryDetail)//电池详情
		battertlist.GET("/batterysoc", batterymanage.GetBatterySOC)//电池SOC
		battertlist.GET("/batterycell", batterymanage.GetBatteryCell)//电池单体
		battertlist.GET("/batterytemper", batterymanage.GetBatteryTemper)//电池温度
	}
}
func registerserBatteryMoveRouter(user *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	battertlist := user.Group("/batterymove").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		battertlist.GET("", batterymanage.GetBatteryMove)//电池轨迹
		battertlist.GET("/location", batterymanage.GetBatteryLocation)//电池位置
	}
}
func registerUserDTUlistRouter(user *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	battertlist := user.Group("/dtulist").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		battertlist.GET("", batterymanage.GetDtuList)//dtu列表
		battertlist.DELETE("/:dtu_specInfoId", batterymanage.DelOneDtuList)//删除dtu
	}
}
func registerUserDTUdetailRouter(user *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	battertlist := user.Group("/dtudetail").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		battertlist.GET("", batterymanage.GetDtuDetail)//dtu详情
		battertlist.GET("/dtucsq", batterymanage.GetDtuCSQ)//电池SOC
		battertlist.POST("/dtulock", batterymanage.SetDtuLock)//车辆锁控制

	}
}
func registerUserFirmwareListRouter(user *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	battertlist := user.Group("/firmwarelist").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		battertlist.GET("", batterymanage.GetFirmwareList)//固件列表
		battertlist.POST("", batterymanage.InsertFirmware)//固件列表
		battertlist.DELETE("/:ota_firmwareId", batterymanage.DelOneFirmwareList)//删除dtu
	}
}
func registerUserdtubmsbandRouter(user *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	battertlist := user.Group("/dtubmsbandlog").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		battertlist.GET("", batterymanage.GetDtuBmsBandList)//电池列表
	}
}