package demo

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HelloHandler(self *gin.Context) {
	self.JSONP(http.StatusOK, gin.H{
		"message": "hello",
	})

}

func FromHandler(self *gin.Context) {
	b := self.DefaultPostForm("b", "2")
	a := self.PostForm("a")
	fmt.Println(a)
	self.JSONP(http.StatusOK, gin.H{
		"a": a,
		"b": b,
	})
}
