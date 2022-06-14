package sliceUpload

import (
	"errors"
	"github.com/gin-gonic/gin"
	"io/fs"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	MergeUrl = "/file-upload/merge"
)

func MergeEntry(c *gin.Context)  {
	hash := c.PostForm("hash")
	size := c.PostForm("size")
	intSize, _ := strconv.Atoi(size)
	fsf, err := CheckFile(hash, int64(intSize))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg": err.Error(),
			"data": nil,
		})
		return
	}
	if fsf.MergeChunk() != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg": err.Error(),
			"data": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg": "上传成功",
		"data": nil,
	})
}

func CheckFile(hash string, totalSize int64) (*fileSlices, error) {
	var total int
	fsf := NewFileSlices()
	filepath.WalkDir("./assets/file-slice/" + hash + "/", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		fns := strings.Split(d.Name(), "-")
		rName, index, size := fns[0], fns[1], fns[2]
		intIndex, _ := strconv.Atoi(index)
		intSize, _ := strconv.Atoi(size)
		data, _ := ioutil.ReadFile(path)
		total += intSize
		*fsf = append(*fsf, FileSlice{
			Filename: rName,
			FileHash: hash,
			Index: intIndex,
			Data: data,
			Size: intSize,
		})
		return nil
	})
	if int64(total) != totalSize {
		return nil, errors.New("文件大小错误")
	}
	return fsf, nil
}
