package todo

import (
	"github.com/aasimsajjad22/go-todo-backend/datasources/mysql/todo_db"
	"github.com/aasimsajjad22/go-todo-backend/utils/errors"
)

const (
	querySaveTodo = "INSERT INTO todo (`description`) VALUES (?);"
)

func (t *Todo) Save() *errors.RestErr {
	stmt, err := todo_db.Client.Prepare(querySaveTodo)
	if err != nil {
		return errors.NewInternalServerError("error when trying to prepare statement")
	}
	defer stmt.Close()
	insertRes, insertErr := stmt.Exec(t.Description)
	if insertErr != nil {
		return errors.NewInternalServerError("error when trying to save todo")
	}
	todoId, err := insertRes.LastInsertId()
	if err != nil || todoId <= 0 {
		return errors.NewInternalServerError("error when trying to save todo")
	}
	t.Id = todoId
	return nil
}
