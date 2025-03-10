package rest

import "ki9.com/gin_demo/internal/model/appmodel"

// body for api/handle-task
type HandleTaskBody struct {
	Method string        `json:"method" example:"create"`
	Task   appmodel.Task `json:"task"`
}

// uni-response for rest API
type UniResponse struct {
	Code int         `json:"code" example:"200"`
	Msg  string      `json:"msg" example:"success"`
	Data interface{} `json:"data" `
}
