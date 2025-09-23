package task

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"ki9.com/gin_demo/internal/dto"
	"ki9.com/gin_demo/internal/middleware/myerr"
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
	logger := ctx.MustGet("logger").(*zap.Logger)

	b := &reqBody{}

	// err := ctx.BindJSON(b)  //--> 这里不用should系列的话，会直接自动abort，无法用我的自定义err信息包装
	err := ctx.ShouldBindJSON(b)
	if err != nil {
		logger.Debug("BindJSON failed", zap.Error(err))
		ctx.Error(myerr.ErrBadRequest("body should be json format"))
		return
	}

	//仅接受特定的mathod：CRUD，否则400
	switch b.Method {
	case "create": //TODO 从这里开始整改error处理。

		if err := h.svc.CreateTask(&b.Task); err != nil {
			ctx.Error(err)
			return
		}
		// -----------------------看到这里-
		ctx.JSON(http.StatusOK, dto.UniResponseBody{Code: 200, Msg: "success", Data: b.Task}) // TODO这里不能直接返task，因为task有可能是不合格的，需要由函数返
		logger.Info(fmt.Sprintf("created Task: %v", b.Task))

	case "read":

		searchResults, err := h.svc.GetTasksByUser(b.Task.Creater.Uid)
		if err != nil {
			ctx.Error(err)
			// logger.Logger.Error("GetTasksByUser failed", zap.Error(err)) //TODO 这里直接把error给ctx由中间件统一打印.
			// logger.Logger.Debug("test debug log")
			// ctx.JSON(http.StatusInternalServerError, dto.UniResponseBody{Code: 500, Msg: "internal server error"}) //TODO 加个根据error写非200响应的中间件
			return
		}

		ctx.JSON(http.StatusOK, dto.UniResponseBody{Code: 200, Msg: "success", Data: searchResults})
		logger.Info(fmt.Sprintf("find tasks: %v", searchResults), zap.Int("num", len(searchResults)))

	case "update":

		if err = h.svc.UpdateTask(&b.Task); err != nil {
			logger.Error("UpdateTask failed", zap.Error(err)) //TODO 删掉，以及往下的部分都还没修改
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
				logger.Error("DeleteTasksById failed", zap.Error(err))
				ctx.JSON(http.StatusOK, dto.UniResponseBody{Code: 404, Msg: "no delete: task not found"})
				return
			}

			ctx.JSON(http.StatusOK, dto.UniResponseBody{Code: 200, Msg: "Deleted task", Data: delatedTasks})
		}

	default:
		ctx.JSON(http.StatusBadRequest, dto.UniResponseBody{Code: 400, Msg: "invaild method: only accept \"create\", \"read\", \"update\", or \"delete\" as \"method\""})
	}

}
