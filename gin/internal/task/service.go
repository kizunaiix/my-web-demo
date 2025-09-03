package task // TODO 改完service里的报错

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type TaskService interface {
	GetAllTasks() []*Task
	CreateTask(t Task) error
	GetTasksByUser(userid int) ([]*Task, error)
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
func (svc *taskService) taskIsNew(t *Task) (bool, error) {
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

func (svc *taskService) CreateTask(t *Task) error {

	//先检查task有没有id。没有的话是新的task，要给一个id
	if t.Id == "" {
		t.Id = uuid.New().String()
	}

	ok, err := svc.taskIsNew(t)
	if err != nil {
		return fmt.Errorf("failed to create task: %w", err)
	}
	if ok {
		svc.repo.CreateTask(t)
		if err != nil {
			return fmt.Errorf("failed to create task: %w", err)
		}
	}

	return nil
}

func (svc *taskService) GetTasksByUser(userid int) ([]*Task, error) {
	ts, err := svc.repo.GetTasksByUser(userid)
	return ts, fmt.Errorf("failed to get tasks by user: %w", err)
}

func (svc *taskService) UpdateTask(uid string) bool { //TODO 写到这里
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
