package err

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"ki9.com/gin_demo/internal/dto"
)

var be BizError

func ErrorHandlerMiddleware(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		err := c.Errors.Last().Err

		if ok := errors.As(err, &be); ok { //这里会遇到可能是包装过的err，所以要用errors.As()。记得先定义var be BizError
			c.JSON(
				be.StatusCode(),
				dto.UniResponseBody{
					Msg:  be.Error(),
					Code: be.BizCode(),
					Data: nil,
				},
			)
			c.MustGet("logger").(*zap.Logger).Error("Business error", zap.Error(err))
			return
		}
		// 未知错误
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    9999,
			"message": "internal server error",
			"data":    nil,
		})
		c.MustGet("logger").(*zap.Logger).Error("not Business error", zap.Error(err))

		return
	}

}
