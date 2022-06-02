package getWxFollowUrl

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"strings"
)

const tempUrl = "https://mp.weixin.qq.com/mp/profile_ext?action=home&scene=126"
const Url = "/get-wx-follow-url"

type UrlResult struct {
	Code int
	Msg, Data string
}

func Entry(c *gin.Context)  {
	url := c.PostForm("url")
	if url == "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "urls必传且不能为空字符串",
		})
		return
	}
	resultChain := make(chan UrlResult, 1)
	go grabUrl(url, resultChain)
	res := <- resultChain
	c.JSON(http.StatusOK, gin.H{
		"code": res.Code,
		"message": res.Msg,
		"data": res.Data,
	})
}

func grabUrl(url string, result chan UrlResult) {
	const necessaryStr = "mp.weixin.qq.com"
	if !strings.Contains(url, necessaryStr) {
		res := UrlResult{
			Code: 400,
			Msg: "url不是微信文章的url",
			Data: "",
		}
		result <- res
		return
	}
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		res := UrlResult{
			Code: 400,
			Msg: "url请求出错",
			Data: "",
		}
		result <- res
		return
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		res := UrlResult{
			Code: 400,
			Msg: "html解析出错",
			Data: "",
		}
		result <- res
		return
	}
	parseHtml(doc, result)
}

func parseHtml(doc *goquery.Document, result chan UrlResult) {
	sel := doc.Find("meta[property='og:url']")
	attr, bool := sel.Attr("content")
	if !bool {
		res := UrlResult{
			Code: 400,
			Msg: "结果不存在",
			Data: "",
		}
		result <- res
		return
	}
	fmt.Println(attr)
	uri, err := url.Parse(attr)
	if err != nil {
		res := UrlResult{
			Code: 400,
			Msg: "url解析错误",
			Data: "",
		}
		result <- res
		return
	}
	vals := uri.Query()
	biz := vals.Get("__biz")
	res := UrlResult{
		Code: 200,
		Msg: "",
		Data: genUrl(biz),
	}
	result <- res
	return
}

func genUrl(biz string) string  {
	urls := []string{
		tempUrl,
		"__biz=" + biz,
	}
	return strings.Join(urls, "&") + "#wechat_redirect"
}