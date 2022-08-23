package transfer

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


// Handler http.Handler 转换成 gin.HandlerFunc
func HandlerToGinHandler(handler http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}
