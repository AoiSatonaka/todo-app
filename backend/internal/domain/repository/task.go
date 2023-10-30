package repository

import "github.com/AoiSatonaka/todo-app/backend/internal/domain/model/task"

type TaskRepository interface {
	List() *[]task.Task
	Get(i int32) *task.Task
	Create(t *task.Task) *[]task.Task
	Update(t *task.Task) *[]task.Task
	Delete(i int32) *[]task.Task
}
