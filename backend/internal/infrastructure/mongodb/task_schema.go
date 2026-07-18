package mongodb

import (

	"github.com/AoiSatonaka/todo-app/backend/internal/domain/task"
)

type taskSchema struct {
	Uid  string
	Task task.Task `bson:"inline"`
}

func newTaskSchema(at task.Task, au string) (*taskSchema, error) {
	return &taskSchema{Task: at, Uid: au}, nil
}
