package application

import (
	"github.com/AoiSatonaka/todo-app/backend/internal/domain/task"
)

type IToDoService interface {
	List(string) (*[]task.Task, error)
	Get(string, string) (*task.Task, error)
	Create(*task.Task, string) (*[]task.Task, error)
	Update(*task.Task, string) (*[]task.Task, error)
	Delete(string, string) (*[]task.Task, error)
}

type ToDoService struct {
	tr task.ITaskRepository
}

var _ IToDoService = &ToDoService{}

func NewToDoService(atr task.ITaskRepository) *ToDoService {
	return &ToDoService{tr: atr}
}

func (ts *ToDoService) List(uid string) (*[]task.Task, error) {
	tasks, err := ts.tr.List(uid)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (ts *ToDoService) Get(i, uid string) (*task.Task, error) {
	task, err := ts.tr.FindById(i, uid)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (ts *ToDoService) Create(t *task.Task, uid string) (*[]task.Task, error) {
	if err := ts.tr.Create(t, uid); err != nil {
		return nil, err
	}

	// get tasks
	tasks, err := ts.List(uid)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (ts *ToDoService) Update(t *task.Task, uid string) (*[]task.Task, error) {
	if err := ts.tr.Update(t, uid); err != nil {
		return nil, err
	}

	// get tasks
	tasks, err := ts.List(uid)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (ts *ToDoService) Delete(i, uid string) (*[]task.Task, error) {
	if err := ts.tr.Delete(i, uid); err != nil {
		return nil, err
	}

	// get tasks
	tasks, err := ts.List(uid)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
