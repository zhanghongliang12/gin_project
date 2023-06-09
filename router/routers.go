package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"go_project/gin_project/app/demo"
	"go_project/gin_project/middleware"
)

type Option func(engine *gin.Engine)

var options = []Option{}

func Include(opts ...Option) {
	options = append(options, opts...)
}

func Init() *gin.Engine {
	r := gin.Default()
	// 注册中间件
	r.Use(middleware.InItMiddleWareInfo())
	// 注册session 中间件
	store := cookie.NewStore([]byte("sessionid"))
	r.Use(sessions.Sessions("session_id", store))

	for _, opt := range options {
		opt(r)
	}
	return r
}

func SetupRouter() *gin.Engine {
	Include(demo.DemoRouter)
	r := Init()
	return r
}
