package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"super-tool/getWxFollowUrl"
)

func main()  {
	r := gin.Default()
	r.GET("/api/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "123",
		})
	})
	route := r.Group("/api")
	route.POST(getWxFollowUrl.Url, getWxFollowUrl.Entry)
	r.Run(":4396")
}