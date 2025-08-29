package task // TODO 改完service里的报错

import (
	"net/http"

	"github.com/google/uuid"
	"ki9.com/gin_demo/internal/dto"
)

type TaskService interface {
	TaskIsAlreadyExist(t *Task) bool
	GetAllTasks() []Task
	CreateTask(t Task)
	ReadTask(uid string) []Task
	UpdateTask(uid string) bool
	DeleteTask(uids ...string) []Task
}

type taskService struct {
	repo TaskRepository
}

func NewTaskService(repo TaskRepository) *taskService {
	return &taskService{repo: repo}
}

// 判断解析出的task是不是旧的,id不为空即为新的task
func (svc *taskService) TaskIsAlreadyExist(t *Task) bool { //TODO: 函数名不合适
	if t.Id != "" {
		return true
	}

	//同用户创建的同内容的第二个Task无效
	for _, i := range svc.repo.GetAllTasks() {
		if i.Creater.Uid == t.Creater.Uid && i.Description == t.Description {
			return true
		}
	}

	return false
}

func (svc *taskService) GetAllTasks() []Task {
	return svc.repo.GetAllTasks()
}

func (svc *taskService) CreateTask(t Task) {
	if h.svc.TaskIsAlreadyExist(&b.Task) {
		ctx.JSON(http.StatusBadRequest, dto.UniResponseBody{Code: 400, Msg: "failed: task already exists"})
		return
	}

	b.Task.Id = uuid.New().String()

	allTasks = append(h.svc.GetAllTasks(), b.Task)

}

func (svc *taskService) ReadTask(uid string) []Task {
	for _, v := range h.svc.GetAllTasks() {

		if v.Creater.Uid == b.Task.Creater.Uid {
			searchResults = append(searchResults, v)
		}
	}
}

func (svc *taskService) UpdateTask(uid string) bool {
	updated := false
	for i, v := range h.svc.GetAllTasks() {
		if v.Id == b.Task.Id {
			h.svc.GetAllTasks()[i] = b.Task
			updated = true
		}
	}
}

func (svc *taskService) DeleteTask(uids ...string) []Task {
	//遍历数据库并删除相应id的task
	updatedTasks := []Task{}
	delatedTasks := []Task{}

	for i, v := range h.svc.GetAllTasks() {
		if v.Id != b.Task.Id {
			updatedTasks = append(updatedTasks, h.svc.GetAllTasks()[i])
		} else {
			delatedTasks = append(delatedTasks, h.svc.GetAllTasks()[i])
		}
	}
	allTasks = append(allTasks, updatedTasks...)
}
