package allmodel

import (
	"time"
)

type Task struct {
	Id             int       `json:"id"`
	CreateTime     time.Time `json:"createTime"`
	UpdateTime     time.Time `json:"updateTime"`
	PlanToDoneTime time.Time `json:"planToDoneTime"`
	ActualDoneTime time.Time `json:"actualDoneTime"`
	Creater        User      `json:"creater"`
	Description    string    `json:"description"`
	Status         string    `json:"status"`
}

// 判断解析出的task是不是有效,id不为空即为有效
func (t *Task) isValid() bool {
	return !(t.Id == 0)
}
