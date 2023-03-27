package demo

import "github.com/gin-gonic/gin"

func DemoRouter(r *gin.Engine) {
	demo := r.Group("/demo")
	{
		demo.POST("/form", FromHandler)
		demo.POST("/cookie", FromHandler)

	}

}
