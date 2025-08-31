package task

import "gorm.io/gorm"

type TaskRepository interface {
	// TODO 增删改查方法的抽象搬过来
	GetAllTasks() []*Task
	GetTaskById(id string) (*Task, error)
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

func (r *taskRepositoryMemSlice) GetTaskById(id string) (*Task, error) {
	for _, t := range r.tasks {
		if id == t.Id {
			return t, nil
		}
	} // TODO 写到这里/
	return nil, nil
}
