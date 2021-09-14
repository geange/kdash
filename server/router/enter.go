package router

import (
	"gin-vue-admin/router/autocode"
	"gin-vue-admin/router/example"
	"gin-vue-admin/router/kong"
	"gin-vue-admin/router/system"
)

type RouterGroup struct {
	System   system.RouterGroup
	Example  example.RouterGroup
	Autocode autocode.RouterGroup
	Kong     kong.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
