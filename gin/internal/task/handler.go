package task

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"ki9.com/gin_demo/internal/dto"
	"ki9.com/gin_demo/pkg/logger"
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
func (h *TaskHandler) TaskHandlerFunc(ctx *gin.Context) {
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

		h.svc.CreateTask(&b.Task)

		ctx.JSON(http.StatusOK, dto.UniResponseBody{Code: 200, Msg: "success", Data: b.Task})
		log.Printf("created Task: %v\n", b.Task)

	case "read":

		searchResults, err := h.svc.GetTasksByUser(b.Task.Creater.Uid)
		if err != nil {
			logger.Logger.Error("GetTasksByUser failed", zap.Error(err)) //TODO 这里直接把error给ctx由中间件统一打印.
			logger.Logger.Debug("test debug log")
			ctx.JSON(http.StatusInternalServerError, dto.UniResponseBody{Code: 500, Msg: "internal server error"}) //TODO 加个根据error写非200响应的中间件
			return
		}

		ctx.JSON(http.StatusOK, dto.UniResponseBody{Code: 200, Msg: "success", Data: searchResults})
		logger.Logger.Info(fmt.Sprintf("find tasks: %v", searchResults), zap.Int("num", len(searchResults)))

	case "update":

		if err = h.svc.UpdateTask(&b.Task); err != nil {
			logger.Logger.Error("UpdateTask failed", zap.Error(err))
			ctx.JSON(http.StatusOK, dto.UniResponseBody{Code: 404, Msg: "no update: task not found"})
			return
		}

		ctx.JSON(http.StatusOK, dto.UniResponseBody{Code: 200, Msg: "success", Data: b.Task})

	case "delete":
		if b.Task.Id == "" {
			log.Printf("invalid request: no task id")
			ctx.JSON(http.StatusBadRequest, dto.UniResponseBody{Code: 400, Msg: "invalid request: no task id"})
		} else {

			delatedTasks, err := h.svc.DeleteTasksById(b.Task.Id)
			if err != nil {
				logger.Logger.Error("DeleteTasksById failed", zap.Error(err))
				ctx.JSON(http.StatusOK, dto.UniResponseBody{Code: 404, Msg: "no delete: task not found"})
				return
			}

			ctx.JSON(http.StatusOK, dto.UniResponseBody{Code: 200, Msg: "Deleted task", Data: delatedTasks})
		}

	default:
		ctx.JSON(http.StatusBadRequest, dto.UniResponseBody{Code: 400, Msg: "invaild method: only accept \"create\", \"read\", \"update\", or \"delete\" as \"method\""})
	}

}
