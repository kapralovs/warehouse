package tasks

type TaskRepository interface {
	CreateTask() error
	EditTask() error
	GetTasks() error
	DeleteTask() error
}
