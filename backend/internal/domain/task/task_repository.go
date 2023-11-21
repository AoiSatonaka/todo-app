package task

type ITaskRepository interface {
	List() (*[]Task, error)
	FindById(int32) (*Task, error)
	Create(*Task) (*[]Task, error)
	Update(*Task) (*[]Task, error)
	Delete(int32) (*[]Task, error)
}
