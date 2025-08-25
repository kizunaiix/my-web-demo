package task

type TaskRepository interface {
	// TODO 增删改查方法的抽象搬过来
}

// 用来代替数据库，暂时先用变量实现task的储存
type taskRepositoryMemSlice struct {
	tasks []Task
}

func NewTaskRepositoryMemSlice() *taskRepositoryMemSlice {
	return &taskRepositoryMemSlice{
		tasks: make([]Task, 0),
	}
}
