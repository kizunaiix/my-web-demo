package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"ki9.com/gin_demo/internal/model/allmodel"
)

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
// TODO:把api的描述再完善一下
// TODO 学习一下GraphQL的api风格
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

		if b.Task.IsNew() {
			allmodel.PgDatabaseTasks = append(allmodel.PgDatabaseTasks, b.Task)
			ctx.JSON(http.StatusOK, b.Task)
			log.Printf("created Task: %v\n", b.Task)
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Task already exists"})

		}

	case "read":

		for _, i := range allmodel.PgDatabaseTasks {

			var findResults []allmodel.Task

			if i.Creater.Uid == b.Task.Creater.Uid {
				findResults = append(findResults, i)
			}
			ctx.JSON(http.StatusOK, findResults)
			log.Printf("find tasks: %v", findResults)
		}

	case "update":
		//TODO
		log.Panicln("update not implemented!!")

	case "delete":
		//TODO
		log.Panicln("delete not implemented!!")
	default:
		ctx.String(http.StatusBadRequest, "invaild method")

	}

}
