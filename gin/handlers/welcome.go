package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary 处理 RPC 风格 API
// @Description 通过 `method` 参数区分创建、更新、删除等操作
// @Tags RPC
// @Accept json
// @Produce json
// @Param service path string true "服务名称"
// @Param method query string false "操作方法（create/update/delete）"
// @Success 200 {object} map[string]interface{} "成功响应"
// @Failure 400 {object} map[string]interface{} "请求错误"
// @Failure 404 {object} map[string]interface{} "服务不存在"
// @Router /api [get]
func Welcome(c *gin.Context) {
	c.String(http.StatusOK, "welcome to api")
}
