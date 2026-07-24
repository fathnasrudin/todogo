package todo

import (
	"fmt"
	"log"

	"github.com/google/uuid"
)

type Task struct {
	Title string `json:"title"`
	ID    string `json:"id"`
}

var Tasks []Task

// interface

type ITaskService interface {
	List() ([]Task, error)
	Create(data CreateTaskInput) error 
	Update(taskId string, tData UpdateTaskInput) error 
	Delete(taskId string) error
}


func NewTaskService(r TodoRepository) *TaskService {
	return &TaskService{repo: r}
}


type TaskService struct {
	repo TodoRepository
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

func (s *TaskService) Create(t CreateTaskInput) error {
	IDByte, err := uuid.NewV7()

	if err != nil {
		log.Fatalln("failed to generate ID", err)
	}
	newTask := Task{
		Title: t.Title,
		ID:    IDByte.String(),
	}

	if err := s.repo.Create(newTask); err != nil {return err}
	return nil
}

func (s *TaskService) Delete(taskId string) error {
	err := s.repo.Delete(taskId);
	if err != nil {
		return err
	}	
	return nil;
}

func (s *TaskService) Update(taskId string, tData UpdateTaskInput) error {
	if err := s.repo.Update(taskId, Task{Title: tData.Title}); err != nil { return err}
	return nil
}

func (s *TaskService) List() ([]Task, error) {
	tasks, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}