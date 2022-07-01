package store

import "github.com/JohnNON/tasker/internal/model"

type Store interface {
	CreateTask(*model.Task) error
	CompleteTask(string) error
}
