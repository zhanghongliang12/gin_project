package demo

import "github.com/gin-gonic/gin"

func DemoRouter(r *gin.Engine) {
	r.Group("/demo")
	{
		r.POST("/form", FromHandler)
		r.POST("/cookie", FromHandler)

	}

}
