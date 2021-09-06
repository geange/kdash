package kong

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model/common/response"
	"github.com/gin-gonic/gin"
	"github.com/kong/go-kong/kong"
	"go.uber.org/zap"
)

func NewServicesApi(admin *kong.Client) *ServiceApi {
	return &ServiceApi{admin: admin}
}

type ServiceApi struct {
	admin *kong.Client
}

func (s *ServiceApi) Create(ctx *gin.Context) {
	var params kong.Service
	_ = ctx.ShouldBindJSON(&params)

	service, err := s.admin.Services.Create(ctx, &params)
	if err != nil {
		global.GVA_LOG.Error(CreateServiceError, zap.Any("err", err))
		response.FailWithMessage(CreateServiceError, ctx)
		return
	}
	response.OkWithData(service, ctx)
}

func (s *ServiceApi) Get(ctx *gin.Context) {
	params := struct {
		ID string `form:"id"`
	}{}
	_ = ctx.ShouldBindQuery(&params)

	fmt.Println(params)

	service, err := s.admin.Services.Get(ctx, &params.ID)
	if err != nil {
		global.GVA_LOG.Error(GetServiceError, zap.Any("err", err))
		response.FailWithMessage(GetServiceError, ctx)
		return
	}
	response.OkWithData(service, ctx)
}

func (s *ServiceApi) GetForRoute(ctx *gin.Context) {
	params := struct {
		Route string `form:"route"`
	}{}

	_ = ctx.ShouldBindQuery(&params)

	service, err := s.admin.Services.GetForRoute(ctx, &params.Route)
	if err != nil {
		global.GVA_LOG.Error(GetServiceForRouteError, zap.Any("err", err))
		response.FailWithMessage(GetServiceForRouteError, ctx)
		return
	}
	response.OkWithData(service, ctx)
}

func (s *ServiceApi) Update(ctx *gin.Context) {
	var params kong.Service
	_ = ctx.ShouldBindJSON(&params)

	service, err := s.admin.Services.Update(ctx, &params)
	if err != nil {
		global.GVA_LOG.Error(UpdateServiceError, zap.Any("err", err))
		response.FailWithMessage(UpdateServiceError, ctx)
		return
	}
	response.OkWithData(service, ctx)
}

func (s *ServiceApi) Delete(ctx *gin.Context) {
	params := struct {
		ID string `json:"id"`
	}{}
	_ = ctx.ShouldBindJSON(&params)

	err := s.admin.Services.Delete(ctx, &params.ID)
	if err != nil {
		global.GVA_LOG.Error(DeleteServiceError, zap.Any("err", err))
		response.FailWithMessage(DeleteServiceError, ctx)
		return
	}
	response.Ok(ctx)
}

func (s *ServiceApi) List(ctx *gin.Context) {
	params := struct {
		Size         int       `form:"size,omitempty"`
		Offset       string    `form:"offset,omitempty"`
		Tags         []*string `form:"tags,omitempty"`
		MatchAllTags bool      `form:"match_all_tags"`
	}{}
	_ = ctx.ShouldBindQuery(&params)

	services, opts, err := s.admin.Services.List(ctx, &kong.ListOpt{
		Size:         params.Size,
		Offset:       params.Offset,
		Tags:         params.Tags,
		MatchAllTags: params.MatchAllTags,
	})
	if err != nil {
		global.GVA_LOG.Error(ListServiceError, zap.Any("err", err))
		response.FailWithMessage(ListServiceError, ctx)
		return
	}
	response.OkWithData(map[string]interface{}{
		"list": services,
		"opts": opts,
	}, ctx)
}

func (s *ServiceApi) ListAll(ctx *gin.Context) {
	items, err := s.admin.Services.ListAll(ctx)
	if err != nil {
		global.GVA_LOG.Error(ListServiceError, zap.Any("err", err))
		response.FailWithMessage(ListServiceError, ctx)
		return
	}
	response.OkWithData(map[string]interface{}{
		"list": items,
	}, ctx)
}
