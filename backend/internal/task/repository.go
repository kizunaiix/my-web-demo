package task

import (
	"gorm.io/gorm"
	"ki9.com/gin_demo/internal/middleware/myerr"
)

type TaskRepository interface {
	// TODO 增删改查方法的抽象搬过来
	GetAllTasks() []*Task
	GetTaskById(id string) (*Task, error)
	GetTasksByUser(userid int) ([]*Task, error)
	CreateTask(t *Task) error
	UpdateTask(t *Task) error
	DeleteTask(id string) error
}

var (
	_ TaskRepository = (*taskRepositoryMemSlice)(nil)
	_ TaskRepository = (*taskRepositoryPgsql)(nil)
)

// TODO Pgsql的Repository实现
type taskRepositoryPgsql struct {
	db *gorm.DB
}

func NewTaskRepositoryPgsql(db *gorm.DB) TaskRepository {
	return &taskRepositoryPgsql{db: db}
}

func (r *taskRepositoryPgsql) GetAllTasks() []*Task {
	panic("not implemented") // TODO: Implement
}

func (r *taskRepositoryPgsql) GetTaskById(id string) (*Task, error) {
	panic("not implemented") // TODO: Implement
}

func (r *taskRepositoryPgsql) GetTasksByUser(userid int) ([]*Task, error) {
	panic("not implemented") // TODO: Implement
}

func (r *taskRepositoryPgsql) CreateTask(t *Task) error {
	panic("not implemented") // TODO: Implement
}

func (r *taskRepositoryPgsql) UpdateTask(t *Task) error {
	panic("not implemented") // TODO: Implement
}

func (r *taskRepositoryPgsql) DeleteTask(id string) error {
	panic("not implemented") // TODO: Implement
}

// 用来代替数据库，暂时先用变量实现task的储存
type taskRepositoryMemSlice struct {
	tasks []*Task
}

func NewTaskRepositoryMemSlice() TaskRepository {
	return &taskRepositoryMemSlice{
		tasks: make([]*Task, 0),
	}
}

func (r *taskRepositoryMemSlice) GetAllTasks() []*Task {
	return r.tasks
}

func (r *taskRepositoryMemSlice) GetTaskById(id string) (*Task, error) {
	for _, t := range r.tasks {
		if id == t.Id {
			return t, nil
		}
	}
	return &Task{}, myerr.ErrNotFound("Task not found")
}

func (r *taskRepositoryMemSlice) GetTasksByUser(userid int) (ts []*Task, e error) {
	for _, t := range r.tasks {
		if userid == t.Creator.Uid {
			ts = append(ts, t)
		}
	}

	if len(ts) == 0 {
		return []*Task{}, myerr.ErrNotFound("no tasks found for this user")
	}
	return ts, nil
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
	return myerr.ErrNotFound("Task not found")
}

func (r *taskRepositoryMemSlice) DeleteTask(id string) error {
	for i, v := range r.tasks {
		if v.Id == id {
			r.tasks = append(r.tasks[:i], r.tasks[i+1:]...)
			return nil
		}
	}
	return myerr.ErrNotFound("Task not found")
}
