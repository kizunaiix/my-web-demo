package task

type taskService struct {
	repo taskRepository
}

func NewTaskService(repo taskRepository) *taskService {
	return &taskService{repo: repo}
}
