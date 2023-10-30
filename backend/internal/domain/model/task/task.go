package task

import (
	"time"

	"github.com/AoiSatonaka/todo-app/backend/internal/domain/model/task/priority"
	"github.com/AoiSatonaka/todo-app/backend/internal/domain/model/task/status"
)

type Task struct {
	id         string
	title      string
	decription string
	limit      time.Time
	priority   priority.Priority
	status     status.Status
	labels     []string
}

func NewTask(i, t, d string, li time.Time, p int32, s int32, la []string) (*Task, error) {
	// cast
	cp, err := priority.New(p)
	if err != nil {
		return nil, err
	}

	cs, err := status.New(s)
	if err != nil {
		return nil, err
	}

	return &Task{id: i, title: t, decription: d, limit: li, priority: cp, status: cs, labels: la}, nil
}
