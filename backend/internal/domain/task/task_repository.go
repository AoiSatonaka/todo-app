package task

type ITaskRepository interface {
	List(string) (*[]Task, error)
	FindById(string, string) (*Task, error)
	Create(*Task, string) error
	Update(*Task, string) error
	Delete(string, string) error
}
