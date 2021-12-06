package tests

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aasimsajjad22/go-todo-backend/datasources/mysql/todo_db"
	"github.com/aasimsajjad22/go-todo-backend/domain/todo"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	todo_db.SetClient(db)
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return db, mock
}

func TestGetAllApiEndPointSuccess(t *testing.T) {
	db, mock := NewMock()
	defer db.Close()
	rs := sqlmock.NewRows([]string{"id", "description"}).FromCSVString("1, Test Todo")
	mock.ExpectQuery("SELECT (.*) FROM todo").WillReturnRows(rs)

	response, err := http.Get("http://localhost:8081/todos")
	bytes, _ := ioutil.ReadAll(response.Body)
	var todos = make([]todo.Todo, 0)
	err = json.Unmarshal(bytes, &todos)

	assert.Nil(t, err)
	assert.NotNil(t, todos)
	assert.EqualValues(t, "Test Todo", todos[0].Description)
	assert.EqualValues(t, http.StatusOK, response.StatusCode)
}

func TestCreateApiEndPointSuccess(t *testing.T) {
	db, mock := NewMock()
	defer db.Close()
	prep := mock.ExpectPrepare("INSERT INTO todo (.+) VALUES (.+)")
	prep.ExpectExec().WithArgs("Test Todo").WillReturnResult(sqlmock.NewResult(1, 1))

	var reqBody = []byte(`{"description": "Test Todo"}`)
	response, err := http.Post("http://localhost:8081/todo", "application/json", bytes.NewBuffer(reqBody))
	respBytes, _ := ioutil.ReadAll(response.Body)

	var todo todo.Todo
	err = json.Unmarshal(respBytes, &todo)

	assert.Nil(t, err)
	assert.NotNil(t, todo)
	assert.EqualValues(t, "Test Todo", todo.Description)
	assert.EqualValues(t, http.StatusCreated, response.StatusCode)
}
