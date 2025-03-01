package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"ki9.com/gin_demo/internal/model/allmodel"
)

var PgDatabaseTasks []allmodel.Task

type JSONBody struct {
	Method string        `json:"method"`
	Task   allmodel.Task `json:"task"`
}

// @Tags appapi
// @Router /handle-task [post]
// @Summary add, modify or delete a task
// @Description when method is "create || read || update || delete", deal with task to sql
// @Accept json
// @Produce json
// @Param JSON_body body JSONBody true "请求体"
// @Success 200 {object} map[string]interface{} "成功响应"
// @Failure 400 {object} map[string]interface{} "请求错误"
// @Failure 404 {object} map[string]interface{} "服务不存在"
func HandleTask(ctx *gin.Context) {
	b := &JSONBody{}

	err := ctx.BindJSON(b)
	if err != nil {
		log.Println("func HandleTask: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//仅接受特定的mathod：CRUD，否则400
	switch b.Method {
	case "create":
		// do sth
		log.Println("1!!!!!!!!!!!!!!!!")

		ctx.JSON(http.StatusOK, gin.H{"msg": "success"})
		return

	case "read":
		//TODO
		log.Panicln("read not implemented!!")

		ctx.JSON(http.StatusOK, gin.H{"msg": "success"})
		return

	case "update":
		//TODO
		log.Panicln("update not implemented!!")

	case "delete":
		//TODO
		log.Panicln("delete not implemented!!")
	default:
		ctx.String(http.StatusBadRequest, "invaild method")
		return
	}

}
