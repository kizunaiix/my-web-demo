package task

type taskService struct {
	// PgDatabaseTasks is a slice that simulates a database for storing tasks.
	PgDatabaseTasks []Task
}
