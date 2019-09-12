package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	dm := DemoMessage{
		Message: "Hello World",
		Status:  200,
	}
	r.GET("/demo", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg":     "hi",
			"demomsg": dm,
		})
	})

	r.GET("/proxy", func(c *gin.Context) {

		c.Data(http.StatusOK, "text/html", ProxyCall())
	})
	r.Run()
}
