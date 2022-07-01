package store

import "ex.ex/ex/internal/model"

type Store interface {
	CreateTask(*model.Task) error
	CompleteTask(string) error
}
