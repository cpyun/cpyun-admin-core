package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

type page struct {
	Total     int         `json:"total"`
	PageIndex int         `json:"page_index"`
	PageSize  int         `json:"page_size"`
	Data      interface{} `json:"data"`
}

const (
	ERROR   = 500
	SUCCESS = 200
)

func Success(c *gin.Context, msg string, data any) {

	c.AbortWithStatusJSON(http.StatusOK, response{
		Code: SUCCESS,
		Msg:  msg,
		Data: data,
	})
}

func Error(c *gin.Context, code int, msg string) {

	c.AbortWithStatusJSON(http.StatusOK, response{
		Code: code,
		Msg:  msg,
	})
}

// 自定义输出
func MeReturn(c *gin.Context, code int, msg string, data any) {

	c.AbortWithStatusJSON(http.StatusOK, response{
		Code: code,
		Msg:  msg,
		Data: data,
	})

}

// 自定义内容
func Custom(c *gin.Context, httpCode int, jsonObj any)  {
	c.AbortWithStatusJSON(httpCode, jsonObj)
}

func Page(c *gin.Context, result interface{}, total int, pageIndex int, pageSize int, msg string) {
	var page page
	page.Data = result
	page.Total = total
	page.PageIndex = pageIndex
	page.PageSize = pageSize
	Success(c, msg, page)
}

