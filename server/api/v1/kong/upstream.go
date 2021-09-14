package kong

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/common/response"

	"github.com/gin-gonic/gin"
	"github.com/kong/go-kong/kong"
	"go.uber.org/zap"
)

func NewUpstreamApi(admin *kong.Client) *UpstreamApi {
	return &UpstreamApi{admin: admin}
}

type UpstreamApi struct {
	admin *kong.Client
}

func (s *UpstreamApi) Create(ctx *gin.Context) {
	var params kong.Upstream
	_ = ctx.ShouldBindJSON(&params)

	plugin, err := s.admin.Upstreams.Create(ctx, &params)
	if err != nil {
		global.GVA_LOG.Error(CreateError, zap.Any("err", err))
		response.FailWithMessage(CreateError, ctx)
		return
	}
	response.OkWithData(plugin, ctx)
}

func (s *UpstreamApi) Get(ctx *gin.Context) {
	params := struct {
		ID string `form:"id"`
	}{}
	_ = ctx.ShouldBindQuery(&params)

	item, err := s.admin.Upstreams.Get(ctx, &params.ID)
	if err != nil {
		global.GVA_LOG.Error(GetError, zap.Any("err", err))
		response.FailWithMessage(GetError, ctx)
		return
	}
	response.OkWithData(item, ctx)
}

func (s *UpstreamApi) Update(ctx *gin.Context) {
	var params kong.Upstream
	_ = ctx.ShouldBindJSON(&params)

	plugin, err := s.admin.Upstreams.Update(ctx, &params)
	if err != nil {
		global.GVA_LOG.Error(UpdateError, zap.Any("err", err))
		response.FailWithMessage(UpdateError, ctx)
		return
	}
	response.OkWithData(plugin, ctx)
}

func (s *UpstreamApi) Delete(ctx *gin.Context) {
	params := struct {
		ID string `json:"id"`
	}{}

	err := s.admin.Upstreams.Delete(ctx, &params.ID)
	if err != nil {
		global.GVA_LOG.Error(DeleteError, zap.Any("err", err))
		response.FailWithMessage(DeleteError, ctx)
		return
	}
	response.Ok(ctx)
}

func (s *UpstreamApi) List(ctx *gin.Context) {
	params := struct {
		Size         int       `form:"size,omitempty"`
		Offset       string    `form:"offset,omitempty"`
		Tags         []*string `form:"tags,omitempty"`
		MatchAllTags bool      `form:"match_all_tags"`
	}{}
	_ = ctx.ShouldBindQuery(&params)

	items, opts, err := s.admin.Upstreams.List(ctx, &kong.ListOpt{
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

func (s *UpstreamApi) ListAll(ctx *gin.Context) {
	items, err := s.admin.Upstreams.ListAll(ctx)
	if err != nil {
		global.GVA_LOG.Error(ListError, zap.Any("err", err))
		response.FailWithMessage(ListError, ctx)
		return
	}
	response.OkWithData(map[string]interface{}{
		"list": items,
	}, ctx)
}
