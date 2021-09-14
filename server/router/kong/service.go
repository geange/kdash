package kong

import (
	kc "gin-vue-admin/api/v1/kong"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
	"github.com/kong/go-kong/kong"
)

type ServiceRouter struct {
}

func (*ServiceRouter) InitServiceRouter(admin *kong.Client, Router *gin.RouterGroup) {
	r := Router.Group("service").Use(middleware.OperationRecord())
	api := kc.NewServicesApi(admin)

	{
		r.POST("create", api.Create)   // 创建服务
		r.PUT("update", api.Update)    // 更新服务
		r.DELETE("delete", api.Delete) // 删除服务
		r.GET("get", api.Get)          // 获取单一服务信息
		r.GET("list", api.List)        // 查询服务列表
		r.GET("all", api.ListAll)      // 获取全部服务
	}
}
