package kong

import (
	kc "gin-vue-admin/api/v1/kong"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
	"github.com/kong/go-kong/kong"
)

type ConsumerRouter struct {
}

func (s *ConsumerRouter) InitConsumerRouter(admin *kong.Client, Router *gin.RouterGroup) {
	r := Router.Group("consumer").Use(middleware.OperationRecord())
	api := kc.NewConsumerApi(admin)

	{
		r.POST("create", api.Create)          // 创建
		r.PUT("update", api.Update)           // 更新
		r.DELETE("delete", api.Delete)        // 删除
		r.GET("get", api.Get)                 // 获取单一
		r.GET("custom_id", api.GetByCustomID) //
		r.GET("list", api.List)               // 查询列表
		r.GET("all", api.ListAll)             // 获取全部
	}
}
