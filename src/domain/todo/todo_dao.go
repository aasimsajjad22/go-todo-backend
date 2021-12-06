package todo

import (
	"github.com/aasimsajjad22/go-todo-backend/datasources/mysql/todo_db"
	"github.com/aasimsajjad22/go-todo-backend/utils/errors"
)

const (
	querySaveTodo = "INSERT INTO todo (`description`) VALUES (?);"
	queryGetTodos = "SELECT * FROM todo ORDER BY id DESC ;"
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

func (t *Todo) GetAll() ([]Todo, *errors.RestErr) {
	rows, err := todo_db.Client.Query(queryGetTodos)
	if err != nil {
		return nil, errors.NewInternalServerError("error when trying to get all todos")
	}
	defer rows.Close()
	todos := make([]Todo, 0)
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.Id, &todo.Description); err != nil {
			return nil, errors.NewInternalServerError("error when tying to scan todos")
		}
		todos = append(todos, todo)
	}
	return todos, nil
}
