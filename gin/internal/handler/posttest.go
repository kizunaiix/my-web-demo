package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type hellomessage struct {
	Content string `json:"content"`
}

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
func PostTest(ctx *gin.Context) {

	req := hellomessage{}

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
