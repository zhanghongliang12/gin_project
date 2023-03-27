package demo

import (
	"fmt"
	"github.com/gin-contrib/sessions"
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

func CookiceHandler(self *gin.Context) {

	cookie, err := self.Cookie("key_cookice")
	if err != nil {
		cookie = "NOt set"
		self.SetCookie("key_cookice", "value", 60, "/", "localhost", false, true)

	}
	fmt.Println("cookie是:", cookie)
	self.JSONP(http.StatusOK, gin.H{
		"cookice": cookie,
	})
}

func SessionHandler(self *gin.Context) {
	session := sessions.Default(self)
	if session.Get("hello") != "1" {
		session.Set("hello", "1")
		session.Delete("1")
		session.Save()
	}
	self.JSONP(http.StatusOK, gin.H{
		"session": session.Get("hello"),
	})
}
