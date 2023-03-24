package router

import (
	"github.com/gin-gonic/gin"
	"go_project/gin_project/app/demo"
)

type Option func(engine *gin.Engine)

var options = []Option{}

func Include(opts ...Option) {
	options = append(options, opts...)
}

func Init() *gin.Engine {
	r := gin.Default()
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
