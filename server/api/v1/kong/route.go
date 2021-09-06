package kong

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/common/response"
	"github.com/gin-gonic/gin"
	"github.com/kong/go-kong/kong"
	"go.uber.org/zap"
)

func NewRouteApi(admin *kong.Client) *RouteApi {
	return &RouteApi{admin: admin}
}

type RouteApi struct {
	admin *kong.Client
}

func (s *RouteApi) Create(ctx *gin.Context) {
	var params kong.Route
	_ = ctx.ShouldBindJSON(&params)

	route, err := s.admin.Routes.Create(ctx, &params)
	if err != nil {
		global.GVA_LOG.Error(CreateRouteError, zap.Any("err", err))
		response.FailWithMessage(CreateRouteError, ctx)
		return
	}
	response.OkWithData(route, ctx)
}

func (s *RouteApi) CreateInService(ctx *gin.Context) {
	params := struct {
		SericeID string     `json:"serice_id"`
		Route    kong.Route `json:"route"`
	}{}
	_ = ctx.ShouldBindJSON(&params)

	route, err := s.admin.Routes.CreateInService(ctx, &params.SericeID, &params.Route)
	if err != nil {
		global.GVA_LOG.Error(CreateRouteError, zap.Any("err", err))
		response.FailWithMessage(CreateRouteError, ctx)
		return
	}
	response.OkWithData(route, ctx)
}

func (s *RouteApi) Get(ctx *gin.Context) {
	params := struct {
		ID string `form:"id"`
	}{}
	_ = ctx.ShouldBindQuery(&params)

	route, err := s.admin.Routes.Get(ctx, &params.ID)
	if err != nil {
		global.GVA_LOG.Error(GetRouteError, zap.Any("err", err))
		response.FailWithMessage(GetRouteError, ctx)
		return
	}
	response.OkWithData(route, ctx)
}

func (s *RouteApi) Update(ctx *gin.Context) {
	var params kong.Route
	_ = ctx.ShouldBindJSON(&params)

	route, err := s.admin.Routes.Update(ctx, &params)
	if err != nil {
		global.GVA_LOG.Error(UpdateRouteError, zap.Any("err", err))
		response.FailWithMessage(UpdateRouteError, ctx)
		return
	}
	response.OkWithData(route, ctx)
}

func (s *RouteApi) Delete(ctx *gin.Context) {
	params := struct {
		ID string `json:"id"`
	}{}
	_ = ctx.ShouldBindJSON(&params)

	err := s.admin.Routes.Delete(ctx, &params.ID)
	if err != nil {
		global.GVA_LOG.Error(DeleteRouteError, zap.Any("err", err))
		response.FailWithMessage(DeleteRouteError, ctx)
		return
	}
	response.Ok(ctx)
}

func (s *RouteApi) List(ctx *gin.Context) {
	params := struct {
		Size         int       `form:"size,omitempty"`
		Offset       string    `form:"offset,omitempty"`
		Tags         []*string `form:"tags,omitempty"`
		MatchAllTags bool      `form:"match_all_tags"`
	}{}
	_ = ctx.ShouldBindQuery(&params)

	routes, opts, err := s.admin.Routes.List(ctx, &kong.ListOpt{
		Size:         params.Size,
		Offset:       params.Offset,
		Tags:         params.Tags,
		MatchAllTags: params.MatchAllTags,
	})
	if err != nil {
		global.GVA_LOG.Error(ListRouteError, zap.Any("err", err))
		response.FailWithMessage(ListRouteError, ctx)
		return
	}
	response.OkWithData(map[string]interface{}{
		"list": routes,
		"opts": opts,
	}, ctx)
}

func (s *RouteApi) ListAll(ctx *gin.Context) {
	items, err := s.admin.Routes.ListAll(ctx)
	if err != nil {
		global.GVA_LOG.Error(ListRouteError, zap.Any("err", err))
		response.FailWithMessage(ListRouteError, ctx)
		return
	}
	response.OkWithData(map[string]interface{}{
		"list": items,
	}, ctx)
}

func (s *RouteApi) ListForService(ctx *gin.Context) {
	params := struct {
		ServiceID    string    `json:"service_id"`
		Size         int       `json:"size,omitempty"`
		Offset       string    `json:"offset,omitempty"`
		Tags         []*string `json:"tags,omitempty"`
		MatchAllTags bool      `json:"match_all_tags"`
	}{}
	_ = ctx.ShouldBindJSON(&params)

	routes, opts, err := s.admin.Routes.ListForService(ctx, &params.ServiceID, &kong.ListOpt{
		Size:         params.Size,
		Offset:       params.Offset,
		Tags:         params.Tags,
		MatchAllTags: params.MatchAllTags,
	})
	if err != nil {
		global.GVA_LOG.Error(ListRouteError, zap.Any("err", err))
		response.FailWithMessage(ListRouteError, ctx)
		return
	}
	response.OkWithData(map[string]interface{}{
		"list": routes,
		"opts": opts,
	}, ctx)
}
