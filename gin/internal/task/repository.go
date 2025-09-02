package task

import (
	"errors"

	"gorm.io/gorm"
)

type TaskRepository interface {
	// TODO 增删改查方法的抽象搬过来
	GetAllTasks() []*Task
	GetTaskById(id string) (*Task, error)
	CreateTask(t *Task) error
	UpdateTask(t *Task) error
	DeleteTask(id string) error
}

// TODO Pgsql的Repository实现
type taskRepositoryPgsql struct {
	db *gorm.DB
}

// 用来代替数据库，暂时先用变量实现task的储存
type taskRepositoryMemSlice struct {
	tasks []*Task
}

func NewTaskRepositoryMemSlice() *taskRepositoryMemSlice {
	return &taskRepositoryMemSlice{
		tasks: make([]*Task, 0),
	}
}

func (r *taskRepositoryMemSlice) GetAllTasks() []*Task {
	return r.tasks
}

func (r *taskRepositoryMemSlice) GetTaskById(id string) (t *Task, e error) {
	for _, t := range r.tasks {
		if id == t.Id {
			return t, nil
		}
	}
	return &Task{}, errors.New("task not found")
}

func (r *taskRepositoryMemSlice) CreateTask(t *Task) error {
	r.tasks = append(r.tasks, t)
	return nil
}

func (r *taskRepositoryMemSlice) UpdateTask(t *Task) error {
	for i, v := range r.tasks {
		if v.Id == t.Id {
			r.tasks[i] = t
			return nil
		}
	}
	return errors.New("task not found")
}

func (r *taskRepositoryMemSlice) DeleteTask(id string) error {
	for i, v := range r.tasks {
		if v.Id == id {
			r.tasks = append(r.tasks[:i], r.tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}
