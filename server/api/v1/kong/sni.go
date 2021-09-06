package kong

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/common/response"

	"github.com/gin-gonic/gin"
	"github.com/kong/go-kong/kong"
	"go.uber.org/zap"
)

func NewSniApi(admin *kong.Client) *SniApi {
	return &SniApi{admin: admin}
}

type SniApi struct {
	admin *kong.Client
}

func (s *SniApi) Create(ctx *gin.Context) {
	var params kong.SNI
	_ = ctx.ShouldBindJSON(&params)

	plugin, err := s.admin.SNIs.Create(ctx, &params)
	if err != nil {
		global.GVA_LOG.Error(CreateError, zap.Any("err", err))
		response.FailWithMessage(CreateError, ctx)
		return
	}
	response.OkWithData(plugin, ctx)
}

func (s *SniApi) Get(ctx *gin.Context) {
	params := struct {
		ID string `form:"id"`
	}{}
	_ = ctx.ShouldBindQuery(&params)

	item, err := s.admin.SNIs.Get(ctx, &params.ID)
	if err != nil {
		global.GVA_LOG.Error(GetError, zap.Any("err", err))
		response.FailWithMessage(GetError, ctx)
		return
	}
	response.OkWithData(item, ctx)
}

func (s *SniApi) Update(ctx *gin.Context) {
	var params kong.SNI
	_ = ctx.ShouldBindJSON(&params)

	plugin, err := s.admin.SNIs.Update(ctx, &params)
	if err != nil {
		global.GVA_LOG.Error(UpdateError, zap.Any("err", err))
		response.FailWithMessage(UpdateError, ctx)
		return
	}
	response.OkWithData(plugin, ctx)
}

func (s *SniApi) Delete(ctx *gin.Context) {
	params := struct {
		ID string `json:"id"`
	}{}

	err := s.admin.SNIs.Delete(ctx, &params.ID)
	if err != nil {
		global.GVA_LOG.Error(DeleteError, zap.Any("err", err))
		response.FailWithMessage(DeleteError, ctx)
		return
	}
	response.Ok(ctx)
}

func (s *SniApi) List(ctx *gin.Context) {
	params := struct {
		Size         int       `form:"size,omitempty"`
		Offset       string    `form:"offset,omitempty"`
		Tags         []*string `form:"tags,omitempty"`
		MatchAllTags bool      `form:"match_all_tags"`
	}{}
	_ = ctx.ShouldBindQuery(&params)

	items, opts, err := s.admin.SNIs.List(ctx, &kong.ListOpt{
		Size:         params.Size,
		Offset:       params.Offset,
		Tags:         params.Tags,
		MatchAllTags: params.MatchAllTags,
	})
	if err != nil {
		global.GVA_LOG.Error(ListError, zap.Any("err", err))
		response.FailWithMessage(ListError, ctx)
		return
	}
	response.OkWithData(map[string]interface{}{
		"list": items,
		"opts": opts,
	}, ctx)
}

func (s *SniApi) ListForCertificate(ctx *gin.Context) {
	params := struct {
		ID           string    `form:"id"`
		Size         int       `form:"size,omitempty"`
		Offset       string    `form:"offset,omitempty"`
		Tags         []*string `form:"tags,omitempty"`
		MatchAllTags bool      `form:"match_all_tags"`
	}{}
	_ = ctx.ShouldBindQuery(&params)

	items, opts, err := s.admin.SNIs.ListForCertificate(ctx, &params.ID, &kong.ListOpt{
		Size:         params.Size,
		Offset:       params.Offset,
		Tags:         params.Tags,
		MatchAllTags: params.MatchAllTags,
	})
	if err != nil {
		global.GVA_LOG.Error(ListError, zap.Any("err", err))
		response.FailWithMessage(ListError, ctx)
		return
	}
	response.OkWithData(map[string]interface{}{
		"list": items,
		"opts": opts,
	}, ctx)
}

func (s *SniApi) ListAll(ctx *gin.Context) {
	items, err := s.admin.SNIs.ListAll(ctx)
	if err != nil {
		global.GVA_LOG.Error(ListError, zap.Any("err", err))
		response.FailWithMessage(ListError, ctx)
		return
	}
	response.OkWithData(map[string]interface{}{
		"list": items,
	}, ctx)
}
