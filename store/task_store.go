package store

import (
	"errors"
	"stability-test-task-api/models"
	"sync"
)

var (
	tasks  = make(map[int]models.Task)
	mu     sync.RWMutex
	nextID = 3
)

func init() {
	// Initialize with some default data
	tasks[1] = models.Task{ID: 1, Title: "Learn Go", Done: false}
	tasks[2] = models.Task{ID: 2, Title: "Build API", Done: false}
}

func GetAllTasks() []models.Task {
	mu.RLock()
	defer mu.RUnlock()

	taskList := make([]models.Task, 0, len(tasks))
	for _, t := range tasks {
		taskList = append(taskList, t)
	}
	return taskList
}

func GetTaskByID(id int) (*models.Task, error) {
	mu.RLock()
	defer mu.RUnlock()

	task, ok := tasks[id]
	if !ok {
		return nil, errors.New("task not found")
	}
	// Return a copy to avoid external modification of the store
	return &task, nil
}

func AddTask(task *models.Task) int {
	mu.Lock()
	defer mu.Unlock()

	task.ID = nextID
	tasks[task.ID] = *task
	nextID++
	return task.ID
}

func DeleteTask(id int) error {
	mu.Lock()
	defer mu.Unlock()

	if _, ok := tasks[id]; !ok {
		return errors.New("task not found")
	}

	delete(tasks, id)
	return nil
}
