package errhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"ki9.com/gin_demo/internal/dto"
)

// TODO logger中间件应该记录c.Errors错误

func ErrorHandlerMiddleware(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		err := c.Errors.Last().Err

		if be, ok := err.(BizError); ok {
			c.JSON(
				be.StatusCode(),
				dto.UniResponseBody{
					Msg:  be.Error(),
					Code: be.BizCode(),
					Data: nil,
				},
			)
			return
		}
		// 未知错误
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    9999,
			"message": "internal server error",
			"data":    nil,
		})
	}

}
