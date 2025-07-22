package dto

// uni-response for rest API
type UniResponse struct {
	Code int         `json:"code" example:"200"`
	Msg  string      `json:"msg" example:"success"`
	Data interface{} `json:"data" `
}
