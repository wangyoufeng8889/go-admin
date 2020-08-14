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
	bm1 := r.Group("/api/bm1")
	// 空接口防止v1定义无使用报错
	bm1.GET("/checkrole", nil)

	// {{认证路由自动补充在此处请勿删除}}
	registerUserBatterylistRouter(bm1, authMiddleware)
}
func registerUserBatterylistRouter(user *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	battertlist := user.Group("/battery").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		battertlist.GET("/batterylist", batterymanage.GetBatteryList)
		battertlist.DELETE("/batterylist/:battery_listId", batterymanage.DelOneBatteryList)
	}
}