package task

import (
	"fmt"

	"github.com/google/uuid"
	"ki9.com/gin_demo/internal/middleware/myerr"
)

type TaskService interface {
	GetAllTasks() []*Task
	CreateTask(t *Task) error
	GetTasksByUser(userid int) ([]*Task, error)
	UpdateTask(t *Task) error
	DeleteTasksById(id ...string) ([]*Task, error)
}

var _ TaskService = (*taskService)(nil)

type taskService struct {
	repo TaskRepository
}

func NewTaskService(repo TaskRepository) TaskService {
	return &taskService{repo: repo}
}

func (svc *taskService) GetAllTasks() []*Task {
	return svc.repo.GetAllTasks()
}

func (svc *taskService) CreateTask(t *Task) error {

	//先检查task有没有id。没有的话是新的task，要给一个id
	if t.Id == "" {
		t.Id = uuid.New().String()
	} else {
		return myerr.ErrServerError("failed to create task: task id should be empty for new task")
	}

	//同用户创建的同内容的第二个Task无效
	for _, v := range svc.repo.GetAllTasks() {
		if v.Creator.Uid == t.Creator.Uid && v.Description == t.Description {
			return myerr.ErrServerError("failed to create task: task already exists for this user with the same description")
		}
	}

	if err := svc.repo.CreateTask(t); err != nil {
		return fmt.Errorf("failed to create task: %w", err)
	}

	return nil
}

func (svc *taskService) GetTasksByUser(userid int) ([]*Task, error) {
	ts, err := svc.repo.GetTasksByUser(userid)
	if err != nil {
		return nil, fmt.Errorf("failed to get tasks by user: %w", err)
	}
	return ts, nil
}

func (svc *taskService) UpdateTask(t *Task) error {

	if err := svc.repo.UpdateTask(t); err != nil {
		return fmt.Errorf("failed to update task: %w", err)
	}

	return nil
}

func (svc *taskService) DeleteTasksById(id ...string) ([]*Task, error) {

	var delatedTasks []*Task
	for _, id := range id {
		t, err := svc.repo.GetTaskById(id)
		if err != nil {
			return nil, fmt.Errorf("failed to delete task: %w", err)
		}
		delatedTasks = append(delatedTasks, t)
		err = svc.repo.DeleteTask(id)
		if err != nil {
			return nil, fmt.Errorf("failed to delete task: %w", err)
		}

	}
	return delatedTasks, nil
}
