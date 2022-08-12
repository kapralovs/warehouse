package tasks

type Repository interface {
	CreateTask() error
	EditTask() error
	GetTasks() error
	DeleteTask() error
}

type UseCase interface {
	Start()
	Complete()
}
