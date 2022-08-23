package internal

import (
	"fmt"
	"github.com/cpyun/cpyun-admin-core/logger"
	"github.com/cpyun/cpyun-admin-core/sdk/pkg"
	"github.com/cpyun/cpyun-admin-core/sdk/pkg/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type BaseController struct {
	Context *gin.Context
	Logger  *logger.Helper
	Orm     *gorm.DB
	Errors  error
	Msg     string
}

func (e *BaseController) AddError(err error) error {
	if e.Error == nil {
		e.Errors = err
	} else if err != nil {
		e.Errors = fmt.Errorf("%v; %w", e.Error, err)
	}
	return e.Errors
}

func (e *BaseController) MakeContext(ctx *gin.Context) *BaseController {
	e.Context = ctx
	return e
}

// GetOrm 获取Orm DB
func (e *BaseController) GetOrm() *gorm.DB {
	db, err := pkg.GetOrm(e.Context)
	if err != nil {
		e.Logger.Error(http.StatusInternalServerError, err, "数据库连接获取失败")
		e.AddError(err)
		//return nil
	}
	e.Orm = db
	return e.Orm
}

// Error 通常错误数据处理
func (e BaseController) Error(code int, msg string) {
	response.Error(e.Context, code, msg)
}

//
func (e BaseController) Success(msg string, data any) {
	response.Success(e.Context, msg, data)
}

// 自定义内容
func (e BaseController) Custom(httpCode int, jsonObj any)  {
	response.Custom(e.Context, httpCode, jsonObj)
}