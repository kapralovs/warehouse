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
	Tasks map[string]*models.Task
}

func (ds *DataStorage) LoadTask(id string) (models.Task, error) {
	if task, ok := ds.Tasks[id]; ok {
		return task, nil
	}

	return nil, fmt.Errorf("task with id \"%v\" does not exist", id)
}

func (ds *DataStorage) DeleteTask(id string) error {
	if _, ok := ds.Tasks[id]; ok {
		delete(ds.Tasks, id)
		return nil
	}

	return errors.New("it is not possible to delete a task because it does not exist")
}

func (ds *DataStorage) SaveTask(t models.Task) error {
	if t != nil {
		if t.ID() == "" {
			return errors.New("can't save task with empty ID field")
		}

		ds.mu.Lock()
		defer ds.mu.Unlock()
		ds.Tasks[t.ID()] = t

		log.Printf("The task \"%s\" is saved.\n", ds.Products[t.ID()].ID)
		return nil
	}

	return errors.New("can't save current task, because task is nil")
}
