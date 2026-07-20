package todo

import (
	"fmt"
	"log"
	"slices"

	"github.com/google/uuid"
)

type Task struct {
	Title string `json:"title"`
	ID    string `json:"id"`
}

var Tasks []Task

func NewTask(t CreateTaskInput) Task {
	IDByte, err := uuid.NewV7()

	if err != nil {
		log.Fatalln("failed to generate ID", err)
	}
	return Task{
		Title: t.Title,
		ID:    IDByte.String(),
	}
}

func findTask(taskId string) (*Task, error) {
	var foundTask *Task

	// find task based on id
	for i := range Tasks {
		if Tasks[i].ID == taskId {
			foundTask = &Tasks[i]
			break
		}
	}

	if foundTask == nil {
		return nil, fmt.Errorf("Task with ID %q not found", taskId)
	}

	return foundTask, nil
}

func deleteTask(taskId string) error {
	// find task
	task, err := findTask(taskId)
	if err != nil {
		return err
	}

	// delete task
	for i := range Tasks {
		if Tasks[i].ID == task.ID {
			Tasks = slices.Delete(Tasks, i, i+1)
			break
		}
	}

	return nil
}

func updateTask(taskId string, tData UpdateTaskInput) error {
	var foundTask *Task

	// find task based on id
	for i := range Tasks {
		if Tasks[i].ID == taskId {
			foundTask = &Tasks[i]
			break
		}
	}

	if foundTask == nil {
		return fmt.Errorf("Task with ID %q not found", taskId)
	}

	// task found
	// update task
	foundTask.Title = tData.Title

	// no error mean success update
	return nil
}