package data

import (
	"fmt"

	"github.com/kapralovs/warehouse/internal/tasks"
)

func (ds *DataStorage) CreateTask() error {
	return nil
}

func (ds *DataStorage) LoadTask(id string) (tasks.Task, error) {
	if task, ok := ds.Tasks[id]; ok {
		return task, nil
	}

	return nil, fmt.Errorf("task with id \"%v\" does not exist", id)
}

func (ds *DataStorage) DeleteTask(id string) error {
	return nil
}

func (ds *DataStorage) Save() error {
	return nil
}
