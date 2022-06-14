package sliceUpload

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

const Url = "/file-upload"
const assetsSlicePath = "./assets/file-slice/"
const assetsPath = "./assets/file/"

func PathIsExist(path string) (bool, error)  {
	_, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return true, nil
}

func Entry(c *gin.Context)  {
	filename := c.PostForm("filename")
	size := c.PostForm("size")
	f, _ := c.FormFile("data")
	fns := strings.Split(filename, "-")
	hash, name, index := fns[0], fns[1], fns[2]
	strs := []string{name, index, size}
	path := "./assets/file-slice/"+hash+"/"
	exist, err := PathIsExist(path)
	if !exist {
		os.MkdirAll(path, os.ModePerm)
	}
	err = c.SaveUploadedFile(f, path+strings.Join(strs, "-"))
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "200",
			"data": "",
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"data": "",
		"msg": "success",
	})
}