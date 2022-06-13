package sliceUpload

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const Url = "/file-upload"

func Entry(c *gin.Context)  {
	fmt.Println(c.PostForm("filename"))
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"data": "",
		"msg": "success",
	})
}