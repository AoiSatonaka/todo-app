package task

import (
	"time"

	"github.com/AoiSatonaka/todo-app/backend/internal/domain/task/priority"
	"github.com/AoiSatonaka/todo-app/backend/internal/domain/task/status"
)

type Task struct {
	Limit       time.Time
	Id          string `bson:"_id,omitempty"`
	Title       string
	Description string
	Labels      []string
	Priority    priority.Priority
	Status      status.Status
}

func New(i, t, d string, li time.Time, p int32, s int32, la []string) (*Task, error) {
	// cast
	cp, err := priority.New(p)
	if err != nil {
		return nil, err
	}

	cs, err := status.New(s)
	if err != nil {
		return nil, err
	}

	return &Task{Id: i, Title: t, Description: d, Limit: li, Priority: cp, Status: cs, Labels: la}, nil
}
