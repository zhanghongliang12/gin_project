package demo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HelloHandler(self *gin.Context) {
	self.JSONP(http.StatusOK, gin.H{
		"message": "hello",
	})
}
