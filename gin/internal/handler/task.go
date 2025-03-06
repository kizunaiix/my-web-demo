package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
			b.Task.Id = uuid.New().String()
			allmodel.PgDatabaseTasks = append(allmodel.PgDatabaseTasks, b.Task)
			ctx.JSON(http.StatusOK, b.Task)
			log.Printf("created Task: %v\n", b.Task)
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Task already exists"})
		}

	case "read":

		var findResults []allmodel.Task

		for _, v := range allmodel.PgDatabaseTasks {

			if v.Creater.Uid == b.Task.Creater.Uid {
				findResults = append(findResults, v)
			}
		}

		ctx.JSON(http.StatusOK, findResults)
		log.Printf("find tasks: %v", findResults)

	case "update":
		//TODO
		log.Panicln("update not implemented!!")

	case "delete":
		if b.Task.Id == "" {
			log.Printf("invalid request: no task id")
			ctx.JSON(http.StatusBadRequest, "invalid request: no task id")
		} else {

			//遍历数据库并删除相应id的task
			for i, v := range allmodel.PgDatabaseTasks {
				if v.Id == b.Task.Id {
					allmodel.PgDatabaseTasks = append(allmodel.PgDatabaseTasks[:i], allmodel.PgDatabaseTasks[i+1:]...)
				}
			}
			ctx.JSON(http.StatusOK, fmt.Sprintf("Delated task: %v", b.Task.Id))
		}

	default:
		ctx.String(http.StatusBadRequest, "invaild method")

	}

}
