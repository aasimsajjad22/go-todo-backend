package services

import (
	"github.com/aasimsajjad22/go-todo-backend/domain/todo"
	"github.com/aasimsajjad22/go-todo-backend/utils/errors"
)

var TodoService todoServiceInterface

type todoService struct{}

type todoServiceInterface interface {
	Save(todo.Todo) (*todo.Todo, *errors.RestErr)
}

func init() {
	TodoService = &todoService{}
}

func (t *todoService) Save(todo todo.Todo) (*todo.Todo, *errors.RestErr) {
	return nil, nil
}
