package demo

import "github.com/gin-gonic/gin"

func DemoRouter(r *gin.Engine) {

	r.GET("/", HelloHandler)
}
