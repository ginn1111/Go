package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {

	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {

		return fmt.Sprintf("%s - [%d] [%s]:%s %s \n", params.ClientIP, params.StatusCode, params.Method, params.Path, params.Latency)
	})

}
