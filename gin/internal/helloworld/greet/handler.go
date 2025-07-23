package greet

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Tags Hello-world-test
// @Router /posttest [post]
// @Summary try a JSON-data to this api
// @Description post to it, when body is {"content": "hello api"}, you get 200 with answer
// @Accept json
// @Produce json
// @Param jsonbody body hellomessage true "请求体"
// @Success 200 json string "成功响应"
// @Failure 400 {object} map[string]interface{} "请求错误"
// @Failure 404 {object} map[string]interface{} "服务不存在"
func Greet(ctx *gin.Context) {

	req := reqBody{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Content == "hello api" {
		ctx.JSON(http.StatusOK, gin.H{"content": "hello my client!"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"msg": "not a helloto me, who r u?"})
	}
}
