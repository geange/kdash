package kong

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/common/response"
	"github.com/gin-gonic/gin"
	"github.com/kong/go-kong/kong"
	"go.uber.org/zap"
)

func NewPluginApi(admin *kong.Client) *PluginApi {
	return &PluginApi{admin: admin}
}

type PluginApi struct {
	admin *kong.Client
}

func (s *PluginApi) Create(ctx *gin.Context) {
	var params kong.Plugin
	_ = ctx.ShouldBindJSON(&params)

	plugin, err := s.admin.Plugins.Create(ctx, &params)
	if err != nil {
		global.GVA_LOG.Error(CreateError, zap.Any("err", err))
		response.FailWithMessage(CreateError, ctx)
		return
	}
	response.OkWithData(plugin, ctx)
}

func (s *PluginApi) Get(ctx *gin.Context) {
	params := struct {
		ID string `form:"id"`
	}{}
	_ = ctx.ShouldBindQuery(&params)

	item, err := s.admin.Plugins.Get(ctx, &params.ID)
	if err != nil {
		global.GVA_LOG.Error(GetError, zap.Any("err", err))
		response.FailWithMessage(GetError, ctx)
		return
	}
	response.OkWithData(item, ctx)
}

func (s *PluginApi) Update(ctx *gin.Context) {
	var params kong.Plugin
	_ = ctx.ShouldBindJSON(&params)

	plugin, err := s.admin.Plugins.Update(ctx, &params)
	if err != nil {
		global.GVA_LOG.Error(UpdateError, zap.Any("err", err))
		response.FailWithMessage(UpdateError, ctx)
		return
	}
	response.OkWithData(plugin, ctx)
}

func (s *PluginApi) Delete(ctx *gin.Context) {
	params := struct {
		ID string `json:"id"`
	}{}

	err := s.admin.Plugins.Delete(ctx, &params.ID)
	if err != nil {
		global.GVA_LOG.Error(DeleteError, zap.Any("err", err))
		response.FailWithMessage(DeleteError, ctx)
		return
	}
	response.Ok(ctx)
}

func (s *PluginApi) List(ctx *gin.Context) {
	params := struct {
		Size         int       `form:"size,omitempty"`
		Offset       string    `form:"offset,omitempty"`
		Tags         []*string `form:"tags,omitempty"`
		MatchAllTags bool      `form:"match_all_tags"`
	}{}
	_ = ctx.ShouldBindQuery(&params)

	items, opts, err := s.admin.Plugins.List(ctx, &kong.ListOpt{
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

func (s *PluginApi) ListAll(ctx *gin.Context) {
	items, err := s.admin.Plugins.ListAll(ctx)
	if err != nil {
		global.GVA_LOG.Error(ListError, zap.Any("err", err))
		response.FailWithMessage(ListError, ctx)
		return
	}
	response.OkWithData(map[string]interface{}{
		"list": items,
	}, ctx)
}

func (s *PluginApi) ListAllForConsumer(ctx *gin.Context) {
	params := struct {
		ID string `form:"id"`
	}{}
	_ = ctx.ShouldBindQuery(&params)

	items, err := s.admin.Plugins.ListAllForConsumer(ctx, &params.ID)
	if err != nil {
		global.GVA_LOG.Error(ListError, zap.Any("err", err))
		response.FailWithMessage(ListError, ctx)
		return
	}
	response.OkWithData(map[string]interface{}{
		"list": items,
	}, ctx)
}

func (s *PluginApi) ListAllForService(ctx *gin.Context) {
	params := struct {
		ID string `form:"id"`
	}{}
	_ = ctx.ShouldBindQuery(&params)

	items, err := s.admin.Plugins.ListAllForService(ctx, &params.ID)
	if err != nil {
		global.GVA_LOG.Error(ListError, zap.Any("err", err))
		response.FailWithMessage(ListError, ctx)
		return
	}
	response.OkWithData(map[string]interface{}{
		"list": items,
	}, ctx)
}

func (s *PluginApi) ListAllForRoute(ctx *gin.Context) {
	params := struct {
		ID string `form:"id"`
	}{}
	_ = ctx.ShouldBindQuery(&params)

	items, err := s.admin.Plugins.ListAllForRoute(ctx, &params.ID)
	if err != nil {
		global.GVA_LOG.Error(ListError, zap.Any("err", err))
		response.FailWithMessage(ListError, ctx)
		return
	}
	response.OkWithData(map[string]interface{}{
		"list": items,
	}, ctx)
}

func (s *PluginApi) Validate(ctx *gin.Context) {
	var params kong.Plugin
	_ = ctx.ShouldBindJSON(&params)

	items, err := s.admin.Plugins.Validate(ctx, &params)
	if err != nil {
		global.GVA_LOG.Error(ListError, zap.Any("err", err))
		response.FailWithMessage(ListError, ctx)
		return
	}
	response.OkWithData(map[string]interface{}{
		"list": items,
	}, ctx)
}

func (s *PluginApi) GetSchema(ctx *gin.Context) {
	params := struct {
		Name string `form:"name"`
	}{}
	_ = ctx.ShouldBindQuery(&params)

	item, err := s.admin.Plugins.GetSchema(ctx, &params.Name)
	if err != nil {
		global.GVA_LOG.Error(GetError, zap.Any("err", err))
		response.FailWithMessage(GetError, ctx)
		return
	}
	response.OkWithData(item, ctx)
}
