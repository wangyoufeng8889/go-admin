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
	bm1 := r.Group("/api/bm1/battery")
	// 空接口防止v1定义无使用报错
	bm1.GET("/checkrole", nil)

	// {{认证路由自动补充在此处请勿删除}}
	registerUserBatterylistRouter(bm1, authMiddleware)
	registerUserBatterydetailRouter(bm1, authMiddleware)
	registerUserDTUlistRouter(bm1, authMiddleware)
	registerUserBDTUdetailRouter(bm1, authMiddleware)
}
func registerUserBatterylistRouter(user *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	battertlist := user.Group("/batterylist").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		battertlist.GET("", batterymanage.GetBatteryList)
		battertlist.DELETE("/:bms_specInfoId", batterymanage.DelOneBatteryList)
	}
}
func registerUserBatterydetailRouter(user *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	battertlist := user.Group("/batterydetail").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		battertlist.GET("/bms_statusinfo", batterymanage.GetBatteryDetail_bms_statusinfo)
	}
}
func registerUserDTUlistRouter(user *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	battertlist := user.Group("/dtulist").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		battertlist.GET("", batterymanage.GetDTUPKGList)
		battertlist.DELETE("/:dtuPkg_listId", batterymanage.DelOneDTUPKGList)
	}
}
func registerUserBDTUdetailRouter(user *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	battertlist := user.Group("/dtudetail").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		battertlist.GET("/dtu_statusinfo", batterymanage.GetDTUDetail_dtu_statusinfo)
		battertlist.GET("/dtu_specinfo", batterymanage.GetDTUDetail_dtu_specinfo)
		battertlist.GET("/dtu_aliyun", batterymanage.GetDTUDetail_dtu_aliyun)
	}
}
