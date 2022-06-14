package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"super-tool-service/getWxFollowUrl"
	"super-tool-service/mdToHtml"
	"super-tool-service/sliceUpload"
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
	route.POST(getWxFollowUrl.BatchUrl, getWxFollowUrl.BatchEntry)
	route.POST(mdToHtml.Url, mdToHtml.Entry)
	route.POST(sliceUpload.Url, sliceUpload.Entry)
	route.POST(sliceUpload.MergeUrl, sliceUpload.MergeEntry)
	r.Run(":4396")
}