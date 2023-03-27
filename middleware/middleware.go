package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func InItMiddleWareInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行了")
		c.Set("request", "中间件")
		// 执行函数
		c.Next()
		fmt.Println("请求的地址", c.Request.URL)
		fmt.Println("请求的ip", c.Request.Host)
		fmt.Println("请求的body", c.Request.Body)
		fmt.Println("请求的postFOrm", c.Request.PostForm)
		fmt.Println("请求的RequestURI", c.Request.RequestURI)

		status := c.Writer.Status()

		fmt.Println("中间件执行完毕", status)
		t1 := time.Since(t)

		fmt.Println("执行时间：", t1)

	}
}
