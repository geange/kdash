package kong

import (
	kc "gin-vue-admin/api/v1/kong"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
	"github.com/kong/go-kong/kong"
)

type TargetRouter struct {
}

func (*TargetRouter) InitTargetRouter(admin *kong.Client, Router *gin.RouterGroup) {
	r := Router.Group("target").Use(middleware.OperationRecord())
	api := kc.NewTargetApi(admin)

	{
		r.POST("create", api.Create)           // 创建
		r.DELETE("delete", api.Delete)         // 删除
		r.GET("list", api.List)                // 查询列表
		r.GET("all", api.ListAll)              // 获取全部
		r.POST("healthy", api.MarkHealthy)     //
		r.POST("unhealthy", api.MarkUnhealthy) //
	}
}
