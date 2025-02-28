package allmodel

import "time"

type Task struct {
	CreateTime     time.Time
	UpdateTime     time.Time
	PlanToDoneTime time.Time
	ActualDoneTime time.Time
	Creater        User
	Description    string
	Status         string
}
