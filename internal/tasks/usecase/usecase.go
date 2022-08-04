package usecase

import "github.com/kapralovs/warehouse/internal/tasks"

type TasksUseCase struct {
	tasksRepo tasks.Repository
}
