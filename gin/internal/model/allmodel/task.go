package allmodel

import (
	"time"
)

// 用来代替数据库，暂时先用变量实现task的储存
var PgDatabaseTasks []Task

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

// 判断解析出的task是不是新的,id不为空即为新的task
func (t *Task) IsNew() bool {
	if t.Id != 0 {
		return false
	}

	//同用户创建的同内容的第二个Task无效
	for _, i := range PgDatabaseTasks {
		if i.Creater.Uid == t.Creater.Uid && i.Description == t.Description {
			return false
		}
	}

	return true
}
