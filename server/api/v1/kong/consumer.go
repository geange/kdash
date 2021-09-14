package kong

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/common/response"
	"github.com/gin-gonic/gin"
	"github.com/kong/go-kong/kong"
	"go.uber.org/zap"
)

func NewConsumerApi(admin *kong.Client) *ConsumerApi {
	return &ConsumerApi{admin: admin}
}

type ConsumerApi struct {
	admin *kong.Client
}

func (s *ConsumerApi) Create(ctx *gin.Context) {
	var params kong.Consumer
	_ = ctx.ShouldBindJSON(&params)

	consumer, err := s.admin.Consumers.Create(ctx, &params)
	if err != nil {
		global.GVA_LOG.Error(CreateConsumerError, zap.Any("err", err))
		response.FailWithMessage(CreateConsumerError, ctx)
		return
	}
	response.OkWithData(consumer, ctx)
}

func (s *ConsumerApi) Get(ctx *gin.Context) {
	params := struct {
		ID string `form:"id"`
	}{}
	_ = ctx.ShouldBindQuery(&params)

	consumer, err := s.admin.Consumers.Get(ctx, &params.ID)
	if err != nil {
		global.GVA_LOG.Error(GetConsumerError, zap.Any("err", err))
		response.FailWithMessage(GetConsumerError, ctx)
		return
	}
	response.OkWithData(consumer, ctx)
}

func (s *ConsumerApi) GetByCustomID(ctx *gin.Context) {
	params := struct {
		ID string `json:"id"`
	}{}
	_ = ctx.ShouldBindJSON(&params)

	consumer, err := s.admin.Consumers.GetByCustomID(ctx, &params.ID)
	if err != nil {
		global.GVA_LOG.Error(GetConsumerError, zap.Any("err", err))
		response.FailWithMessage(GetConsumerError, ctx)
		return
	}
	response.OkWithData(consumer, ctx)
}

func (s *ConsumerApi) Update(ctx *gin.Context) {
	var params kong.Consumer
	_ = ctx.ShouldBindJSON(&params)

	consumer, err := s.admin.Consumers.Update(ctx, &params)
	if err != nil {
		global.GVA_LOG.Error(UpdateConsumerError, zap.Any("err", err))
		response.FailWithMessage(UpdateConsumerError, ctx)
		return
	}
	response.OkWithData(consumer, ctx)
}

func (s *ConsumerApi) Delete(ctx *gin.Context) {
	params := struct {
		ID string `json:"id"`
	}{}
	_ = ctx.ShouldBindJSON(&params)

	err := s.admin.Consumers.Delete(ctx, &params.ID)
	if err != nil {
		global.GVA_LOG.Error(GetConsumerError, zap.Any("err", err))
		response.FailWithMessage(GetConsumerError, ctx)
		return
	}
	response.Ok(ctx)
}

func (s *ConsumerApi) List(ctx *gin.Context) {
	params := struct {
		Size         int       `form:"size,omitempty"`
		Offset       string    `form:"offset,omitempty"`
		Tags         []*string `form:"tags,omitempty"`
		MatchAllTags bool      `form:"match_all_tags"`
	}{}
	_ = ctx.ShouldBindQuery(&params)

	consumers, opts, err := s.admin.Consumers.List(ctx, &kong.ListOpt{
		Size:         params.Size,
		Offset:       params.Offset,
		Tags:         params.Tags,
		MatchAllTags: params.MatchAllTags,
	})
	if err != nil {
		global.GVA_LOG.Error(ListConsumerError, zap.Any("err", err))
		response.FailWithMessage(ListConsumerError, ctx)
		return
	}
	response.OkWithData(map[string]interface{}{
		"list": consumers,
		"opts": opts,
	}, ctx)
}

func (s *ConsumerApi) ListAll(ctx *gin.Context) {
	items, err := s.admin.Consumers.ListAll(ctx)
	if err != nil {
		global.GVA_LOG.Error(ListConsumerError, zap.Any("err", err))
		response.FailWithMessage(ListConsumerError, ctx)
		return
	}
	response.OkWithData(map[string]interface{}{
		"list": items,
	}, ctx)
}
