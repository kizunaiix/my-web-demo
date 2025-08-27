package task

type TaskService interface {
	IsAlreadyExist(r TaskRepository, t *Task) bool
}

type taskService struct {
	repo TaskRepository
}

func NewTaskService(repo TaskRepository) *taskService {
	return &taskService{repo: repo}
}

// 判断解析出的task是不是旧的,id不为空即为新的task
func (ts *taskService) IsAlreadyExist(r TaskRepository, t *Task) bool { //TODO: 函数名不合适
	if t.Id != "" {
		return true
	}

	//同用户创建的同内容的第二个Task无效
	for _, i := range r.GetAllTasks() {
		if i.Creater.Uid == t.Creater.Uid && i.Description == t.Description {
			return true
		}
	}

	return false
}
