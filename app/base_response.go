package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResponseErr(self *gin.Context, message string, data map[string]interface{}) {
	self.JSONP(http.StatusOK, gin.H{
		"status":  300,
		"message": message,
		"data":    data,
	})

}

func ResponseOK(self *gin.Context, data interface{}) {
	self.JSONP(http.StatusOK, gin.H{
		"status":  200,
		"message": "ok",
		"data":    data,
	})
}
