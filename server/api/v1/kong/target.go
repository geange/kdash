package kong

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/common/response"

	"github.com/gin-gonic/gin"
	"github.com/kong/go-kong/kong"
	"go.uber.org/zap"
)

func NewTargetApi(admin *kong.Client) *TargetApi {
	return &TargetApi{admin: admin}
}

type TargetApi struct {
	admin *kong.Client
}

func (s *TargetApi) Create(ctx *gin.Context) {
	params := struct {
		Upstream string      `json:"upstream"`
		Target   kong.Target `json:"target"`
	}{}
	_ = ctx.ShouldBindJSON(&params)

	item, err := s.admin.Targets.Create(ctx, &params.Upstream, &params.Target)
	if err != nil {
		global.GVA_LOG.Error(CreateError, zap.Any("err", err))
		response.FailWithMessage(CreateError, ctx)
		return
	}
	response.OkWithData(item, ctx)
}

func (s *TargetApi) Delete(ctx *gin.Context) {
	params := struct {
		Upstream string `json:"upstream"`
		Target   string `json:"target"`
	}{}

	err := s.admin.Targets.Delete(ctx, &params.Upstream, &params.Target)
	if err != nil {
		global.GVA_LOG.Error(DeleteError, zap.Any("err", err))
		response.FailWithMessage(DeleteError, ctx)
		return
	}
	response.Ok(ctx)
}

func (s *TargetApi) List(ctx *gin.Context) {
	params := struct {
		ID           string    `form:"id"`
		Size         int       `form:"size,omitempty"`
		Offset       string    `form:"offset,omitempty"`
		Tags         []*string `form:"tags,omitempty"`
		MatchAllTags bool      `form:"match_all_tags"`
	}{}
	_ = ctx.ShouldBindQuery(&params)

	items, opts, err := s.admin.Targets.List(ctx, &params.ID, &kong.ListOpt{
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

func (s *TargetApi) ListAll(ctx *gin.Context) {
	params := struct {
		ID string `json:"id"`
	}{}
	_ = ctx.ShouldBindJSON(&params)

	items, err := s.admin.Targets.ListAll(ctx, &params.ID)
	if err != nil {
		global.GVA_LOG.Error(ListError, zap.Any("err", err))
		response.FailWithMessage(ListError, ctx)
		return
	}
	response.OkWithData(map[string]interface{}{
		"list": items,
	}, ctx)
}

func (s *TargetApi) MarkHealthy(ctx *gin.Context) {
	params := struct {
		Upstream string      `json:"upstream"`
		Target   kong.Target `json:"target"`
	}{}

	err := s.admin.Targets.MarkHealthy(ctx, &params.Upstream, &params.Target)
	if err != nil {
		global.GVA_LOG.Error(DeleteError, zap.Any("err", err))
		response.FailWithMessage(DeleteError, ctx)
		return
	}
	response.Ok(ctx)
}

func (s *TargetApi) MarkUnhealthy(ctx *gin.Context) {
	params := struct {
		Upstream string      `json:"upstream"`
		Target   kong.Target `json:"target"`
	}{}

	err := s.admin.Targets.MarkUnhealthy(ctx, &params.Upstream, &params.Target)
	if err != nil {
		global.GVA_LOG.Error(DeleteError, zap.Any("err", err))
		response.FailWithMessage(DeleteError, ctx)
		return
	}
	response.Ok(ctx)
}
