package task // TODO 改完service里的报错

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"ki9.com/gin_demo/internal/dto"
)

type TaskService interface {
	GetAllTasks() []*Task
	CreateTask(t Task)
	ReadTask(uid string) []Task
	UpdateTask(uid string) bool
	DeleteTask(uids ...string) []Task
}

var _ TaskService = (*taskService)(nil)

type taskService struct {
	repo TaskRepository
}

func NewTaskService(repo TaskRepository) *taskService {
	return &taskService{repo: repo}
}

// 判断解析出的task是不是旧的,id不为空即为新的task
func (svc *taskService) taskIsNew(t *Task) (bool, error) { //TODO: 函数名不合适
	//新的task的id应该是空的
	if t.Id != "" {
		return false, errors.New("task id should be empty for new task")
	}

	//同用户创建的同内容的第二个Task无效
	for _, v := range svc.repo.GetAllTasks() {
		if v.Creater.Uid == t.Creater.Uid && v.Description == t.Description {
			return false, errors.New("task already exists for this user with the same description")
		}
	}

	return true, nil
}

func (svc *taskService) GetAllTasks() []*Task {
	return svc.repo.GetAllTasks()
}

func (svc *taskService) CreateTask(t *Task) { // TODO 写到这了，这个函数应该能返回err
	ok, err := svc.taskIsNew(t)
	if !ok {
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
