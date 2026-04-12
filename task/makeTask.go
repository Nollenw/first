package task

import "time"

type Task struct {
	Title         string
	Description   string
	TimeCreated   time.Time
	IsCompleted   bool
	timeCompleted *time.Time
}

func NewTask(title string, description string) *Task {
	return &Task{
		Title:       title,
		Description: description,
		TimeCreated: time.Now(),
		IsCompleted: false,
	}
}
