package getWxFollowUrl

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const BatchUrl = "/get-wx-follow-url/batch"

func BatchEntry(c *gin.Context) {
	urls := c.PostForm("urls")
	if urls == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"message": "urls不能为空",
			"data": nil,
		})
		return
	}
	urlsArr := []string{urls}
	if strings.Contains(urls, ",") {
		urlsArr = strings.Split(urls, ",")
	}
	var urlResults = make(chan UrlResult, len(urlsArr))
	for _, val := range urlsArr {
		go grabUrl(val, urlResults)
	}
	for  {
		if len(urlResults) == len(urlsArr) {
			break
		}
	}
	var data []string
	for range urlsArr {
		data = append(data, (<- urlResults).Data)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"message": "",
		"data": data,
	})
}