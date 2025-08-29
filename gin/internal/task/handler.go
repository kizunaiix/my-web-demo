package task

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"ki9.com/gin_demo/internal/dto"
)

type TaskHandler struct {
	svc TaskService
}

func NewTaskHandler(svc TaskService) *TaskHandler {
	return &TaskHandler{svc: svc}
}

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
// @Param body body reqBody true "Request body"
// @Success 200 {object} dto.UniResponseBody "操作结果"
// @Failure 400 {object} dto.UniResponseBody
// @Failure 404 {object} dto.UniResponseBody
func (h *TaskHandler) TaskHandlerFunc(ctx *gin.Context) { //TODO:CRUD的逻辑应该拆出来放到单独的service.go中
	b := &reqBody{}

	err := ctx.BindJSON(b)
	if err != nil {
		log.Println("func HandleTask: ", err)
		ctx.JSON(http.StatusBadRequest, dto.UniResponseBody{Code: 400, Msg: err.Error()})
		return
	}

	//仅接受特定的mathod：CRUD，否则400
	switch b.Method {
	case "create":

		h.svc.CreateTask(b.Task)

		ctx.JSON(http.StatusOK, dto.UniResponseBody{Code: 200, Msg: "success", Data: b.Task})
		log.Printf("created Task: %v\n", b.Task)

	case "read":

		var searchResults = h.svc.ReadTask(b.Task.Id)

		ctx.JSON(http.StatusOK, dto.UniResponseBody{Code: 200, Msg: "success", Data: searchResults})
		log.Printf("find tasks: %v", searchResults)

	case "update":

		updated := h.svc.UpdateTask(b.Task.Id)

		if !updated {
			ctx.JSON(http.StatusOK, dto.UniResponseBody{Code: 404, Msg: "no update: task not found"})
		} else {
			ctx.JSON(http.StatusOK, dto.UniResponseBody{Code: 200, Msg: "success", Data: b.Task})
		}

	case "delete":
		if b.Task.Id == "" {
			log.Printf("invalid request: no task id")
			ctx.JSON(http.StatusBadRequest, dto.UniResponseBody{Code: 400, Msg: "invalid request: no task id"})
		} else {

			delatedTasks := h.svc.DeleteTask(b.Task.Id)

			ctx.JSON(http.StatusOK, dto.UniResponseBody{Code: 200, Msg: "Deleted task", Data: delatedTasks})
		}

	default:
		ctx.JSON(http.StatusBadRequest, dto.UniResponseBody{Code: 400, Msg: "invaild method: only accept \"create\", \"read\", \"update\", or \"delete\" as \"method\""})
	}

}
