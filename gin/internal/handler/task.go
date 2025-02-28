package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"ki9.com/gin_demo/internal/model/allmodel"
)

// @Tags appapi
// @Router /handle-task [post]
// @Summary add, modify or delete a task
// @Description when method is "add", add a new task
// @Accept json
// @Produce json
// @Param jsonbody body hellomessage true "请求体"
// @Success 200 json string "成功响应"
// @Failure 400 {object} map[string]interface{} "请求错误"
// @Failure 404 {object} map[string]interface{} "服务不存在"
func HandleTask(ctx *gin.Context) {
	err := ctx.BindJSON(&struct { // <- 就是能怎么都返回200,为什么bind会成功？
		Method string        `json:"method"`
		Task   allmodel.Task `json:"task"`
	}{})
	if err != nil {
		log.Println("func HandleTask: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

}
