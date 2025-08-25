package task

import (
	"time"

	"ki9.com/gin_demo/internal/user"
)

var TaskSlice []Task

type Task struct {
	Id             string    `json:"id" example:"d29799b5-c693-42ce-8270-0c2b8ccd8300"`
	CreateTime     time.Time `json:"createTime" example:"2023-01-02T18:07:05Z"`
	UpdateTime     time.Time `json:"updateTime" example:"2023-01-02T20:23:21+01:00"`
	PlanToDoneTime time.Time `json:"planToDoneTime" example:"2025-03-15T14:58:45+08:00"`
	ActualDoneTime time.Time `json:"actualDoneTime" example:"2025-05-01T09:31:27+01:00"`
	Creater        user.User `json:"creater"`
	Description    string    `json:"description" example:"Rule the world"`
	Status         string    `json:"status" example:"pending"`
}
