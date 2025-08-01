package welcome

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Tags Hello-world-test
// @Router /welcome [get]
// @Summary welcome api
// @Description this api is like a "hello world"
// @Accept */*
// @Produce plain
// @Param service path string true "服务名称"
// @Param method query string false "操作方法（create/update/delete）"
// @Success 200 {object} map[string]interface{} "成功响应"
// @Failure 400 {object} map[string]interface{} "请求错误"
// @Failure 404 {object} map[string]interface{} "服务不存在"
func Welcome(c *gin.Context) {
	c.String(http.StatusOK, "welcome to api")
}
