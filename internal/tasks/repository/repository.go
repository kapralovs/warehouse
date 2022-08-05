package tasks

import (
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/kapralovs/warehouse/internal/models"
)

type TaskRepository struct {
	mu    *sync.Mutex
	Tasks map[string]models.Task
}

func New() *TaskRepository {
	return &TaskRepository{}
}

func (tr *TaskRepository) LoadTask(id string) (models.Task, error) {
	if task, ok := tr.Tasks[id]; ok {
		return task, nil
	}

	return nil, fmt.Errorf("task with id \"%v\" does not exist", id)
}

func (tr *TaskRepository) DeleteTask(id string) error {
	if _, ok := tr.Tasks[id]; ok {
		delete(tr.Tasks, id)
		return nil
	}

	return errors.New("it is not possible to delete a task because it does not exist")
}

func (tr *TaskRepository) SaveTask(t models.Task) error {
	if t != nil {
		if t.ID() == "" {
			return errors.New("can't save task with empty ID field")
		}

		tr.mu.Lock()
		defer tr.mu.Unlock()
		tr.Tasks[t.ID] = t

		log.Printf("The task \"%s\" is saved.\n", tr.Products[t.ID()].ID)
		return nil
	}

	return errors.New("can't save current task, because task is nil")
}
