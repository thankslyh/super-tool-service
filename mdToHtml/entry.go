package mdToHtml

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
	"io/ioutil"
	"net/http"
)

const (
	Url = "/upload"
)
func Entry(ctx *gin.Context)  {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	f, err := file.Open()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	content, err := ioutil.ReadAll(f)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	strHtml, _ := mdToHtml(content)
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": strHtml,
	})
}

func mdToHtml(bt []byte) (string, *goquery.Document) {
	unsafe := blackfriday.Run(bt)
	html := bluemonday.UGCPolicy().Sanitize(string(unsafe))
	reader := bytes.NewReader([]byte(html))
	doc, _ := goquery.NewDocumentFromReader(reader)
	return html, doc
}