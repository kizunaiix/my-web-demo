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
	Method string        `json:"method" example:"create"`
	Task   allmodel.Task `json:"task"`
}

// @Tags appapi
// @Router /handle-task [post]
// @Summary Create, read, update, or delete a task based on method
// @Description Perform CRUD operations on tasks.
// @Description - "create": Add new task (auto-generate ID if empty). Returns created task.
// @Description - "read": Get all tasks by creator UID. Returns task array.
// @Description - "update": Modify existing task by ID. Returns updated task.
// @Description - "delete": Remove task by ID. Returns deleted tasks.
// @Accept json
// @Produce json
// @Param JSON_body body JSONBody true "Request body"
// @Success 200 {object} map[string]interface{} "操作结果"
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// TODO 把api的描述再完善一下
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

		updated := false
		for i, v := range allmodel.PgDatabaseTasks {
			if v.Id == b.Task.Id {
				allmodel.PgDatabaseTasks[i] = b.Task
				updated = true
			}
		}

		if !updated {
			ctx.JSON(http.StatusOK, gin.H{"msg": "no update: task not found"})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"msg": "success", "updateTask": b.Task})
		}

	case "delete":
		if b.Task.Id == "" {
			log.Printf("invalid request: no task id")
			ctx.JSON(http.StatusBadRequest, "invalid request: no task id")
		} else {

			//遍历数据库并删除相应id的task
			updatedTasks := []allmodel.Task{}
			delatedTasks := []allmodel.Task{}
			for i, v := range allmodel.PgDatabaseTasks {
				if v.Id != b.Task.Id {
					updatedTasks = append(updatedTasks, allmodel.PgDatabaseTasks[i])
				} else {
					delatedTasks = append(delatedTasks, allmodel.PgDatabaseTasks[i])
				}
			}
			allmodel.PgDatabaseTasks = updatedTasks

			ctx.JSON(http.StatusOK, fmt.Sprintf("Deleted task: %v", delatedTasks))
		}

	default:
		ctx.String(http.StatusBadRequest, "invaild method")

	}

}
