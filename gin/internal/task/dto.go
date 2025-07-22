package task

// body for /api/handle-task
type reqBody struct {
	Method string `json:"method" example:"create"`
	Task   Task   `json:"task"`
}
