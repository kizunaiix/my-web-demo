package task

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"ki9.com/gin_demo/internal/dto"
)

// @Tags appapi
// @Router /handle-task [post]
// @Summary create, read, update, or delete a task based on method
// @Description Perform CRUD operations on tasks.
// @Description - "create": Add new task (auto-generate ID if empty). Returns created task.
// @Description - "read": Get all tasks by creator UID. Returns task array.
// @Description - "update": Modify existing task by ID. Returns updated task.
// @Description - "delete": Remove task by ID. Returns deleted tasks.
// @Accept json
// @Produce json
// @Param JSON_body body rest.HandleTaskBody true "Request body"
// @Success 200 {object} rest.UniResponse "操作结果"
// @Failure 400 {object} rest.UniResponse
// @Failure 404 {object} rest.UniResponse
// TODO 把api的描述再完善一下，并且重新跑通swagger!
func HandleTask(ctx *gin.Context) { //TODO:CRUD的逻辑应该拆出来放到单独的service.go中
	b := &reqBody{}

	err := ctx.BindJSON(b)
	if err != nil {
		log.Println("func HandleTask: ", err)
		ctx.JSON(http.StatusBadRequest, dto.UniResponse{Code: 400, Msg: err.Error()})
		return
	}

	//仅接受特定的mathod：CRUD，否则400
	switch b.Method {
	case "create":

		if b.Task.IsAlreadyExist() {
			ctx.JSON(http.StatusBadRequest, dto.UniResponse{Code: 400, Msg: "failed: task already exists"})
			return
		}

		b.Task.Id = uuid.New().String()
		PgDatabaseTasks = append(PgDatabaseTasks, b.Task)

		ctx.JSON(http.StatusOK, dto.UniResponse{Code: 200, Msg: "success", Data: b.Task})
		log.Printf("created Task: %v\n", b.Task)

	case "read":

		var searchResults []Task

		for _, v := range PgDatabaseTasks {

			if v.Creater.Uid == b.Task.Creater.Uid {
				searchResults = append(searchResults, v)
			}
		}

		ctx.JSON(http.StatusOK, dto.UniResponse{Code: 200, Msg: "success", Data: searchResults})
		log.Printf("find tasks: %v", searchResults)

	case "update":

		updated := false
		for i, v := range PgDatabaseTasks {
			if v.Id == b.Task.Id {
				PgDatabaseTasks[i] = b.Task
				updated = true
			}
		}

		if !updated {
			ctx.JSON(http.StatusOK, dto.UniResponse{Code: 404, Msg: "no update: task not found"})
		} else {
			ctx.JSON(http.StatusOK, dto.UniResponse{Code: 200, Msg: "success", Data: b.Task})
		}

	case "delete":
		if b.Task.Id == "" {
			log.Printf("invalid request: no task id")
			ctx.JSON(http.StatusBadRequest, dto.UniResponse{Code: 400, Msg: "invalid request: no task id"})
		} else {

			//遍历数据库并删除相应id的task
			updatedTasks := []Task{}
			delatedTasks := []Task{}
			for i, v := range PgDatabaseTasks {
				if v.Id != b.Task.Id {
					updatedTasks = append(updatedTasks, PgDatabaseTasks[i])
				} else {
					delatedTasks = append(delatedTasks, PgDatabaseTasks[i])
				}
			}
			PgDatabaseTasks = updatedTasks

			ctx.JSON(http.StatusOK, dto.UniResponse{Code: 200, Msg: "Deleted task", Data: delatedTasks})
		}

	default:
		ctx.JSON(http.StatusBadRequest, dto.UniResponse{Code: 400, Msg: "invaild method: only accept \"create\", \"read\", \"update\", or \"delete\" as \"method\""})
	}

}
