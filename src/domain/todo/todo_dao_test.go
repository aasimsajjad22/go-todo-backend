package todo

import (
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aasimsajjad22/go-todo-backend/datasources/mysql/todo_db"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	todo_db.SetClient(db)
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return db, mock
}

/**
Mocking failed statement creation behaviour
*/
func TestTodoSaveStmtError(t *testing.T) {
	db, mock := NewMock()
	todo := Todo{}
	defer db.Close()
	mock.ExpectPrepare("INSERT INTO todo (.+) VALUES (.+)").WillReturnError(errors.New("database error"))

	saveErr := todo.Save()
	assert.NotNil(t, saveErr)
	assert.EqualValues(t, "error when trying to prepare statement", saveErr.Message)
	assert.EqualValues(t, http.StatusInternalServerError, saveErr.Status)
}

/**
Mocking failed insert query error behaviour
*/
func TestTodoSaveInsertError(t *testing.T) {
	db, mock := NewMock()
	todo := Todo{}
	defer db.Close()
	prep := mock.ExpectPrepare("INSERT INTO todo (.+) VALUES (.+)")
	prep.ExpectExec().WithArgs(todo.Description).WillReturnError(errors.New("database error"))
	saveErr := todo.Save()

	assert.NotNil(t, saveErr)
	assert.EqualValues(t, "error when trying to save todo", saveErr.Message)
	assert.EqualValues(t, http.StatusInternalServerError, saveErr.Status)
}

/**
Mocking failed LastInsertId behaviour
*/
func TestTodoSaveLastInsertIdError(t *testing.T) {
	db, mock := NewMock()
	todo := Todo{}
	defer db.Close()
	prep := mock.ExpectPrepare("INSERT INTO todo (.+) VALUES (.+)")
	prep.ExpectExec().WithArgs(todo.Description).WillReturnResult(sqlmock.NewResult(0, 1))
	saveErr := todo.Save()

	assert.NotNil(t, saveErr)
	assert.EqualValues(t, "error when trying to save todo", saveErr.Message)
	assert.EqualValues(t, http.StatusInternalServerError, saveErr.Status)
}

/**
Mocking success case
*/
func TestTodoSaveNoError(t *testing.T) {
	db, mock := NewMock()
	todo := Todo{}
	defer db.Close()
	prep := mock.ExpectPrepare("INSERT INTO todo (.+) VALUES (.+)")
	prep.ExpectExec().WithArgs(todo.Description).WillReturnResult(sqlmock.NewResult(1, 1))
	saveErr := todo.Save()

	assert.Nil(t, saveErr)
	assert.EqualValues(t, 1, todo.Id)
}

/**
TESTCASES FOR GETTING ALL TODOS
Mocking AllTodo query error case
*/
func TestGetAllTodoQueryError(t *testing.T) {
	db, mock := NewMock()
	todo := Todo{}
	defer db.Close()
	mock.ExpectQuery("SELECT (.*) FROM todo").WillReturnError(errors.New("query error"))
	todos, getErr := todo.GetAll()

	assert.Nil(t, todos)
	assert.NotNil(t, getErr)
	assert.EqualValues(t, "error when trying to get all todos", getErr.Message)
	assert.EqualValues(t, http.StatusInternalServerError, getErr.Status)
}

/**
Mocking getAllTodo result scan error case
*/
func TestGetAllTodoResultScanError(t *testing.T) {
	db, mock := NewMock()
	todo := Todo{}
	defer db.Close()
	rs := sqlmock.NewRows([]string{"id", "description"}).FromCSVString("5")
	mock.ExpectQuery("SELECT (.*) FROM todo").WillReturnRows(rs)
	todos, getErr := todo.GetAll()

	assert.Nil(t, todos)
	assert.NotNil(t, getErr)
	assert.EqualValues(t, "error when tying to scan todos", getErr.Message)
	assert.EqualValues(t, http.StatusInternalServerError, getErr.Status)
}

/**
Mocking getAllTodo success case
*/
func TestGetAllTodoNoError(t *testing.T) {
	db, mock := NewMock()
	todo := Todo{}
	defer db.Close()
	rs := sqlmock.NewRows([]string{"id", "description"}).FromCSVString("1, Test Todo")
	mock.ExpectQuery("SELECT (.*) FROM todo").WillReturnRows(rs)
	todos, getErr := todo.GetAll()

	assert.Nil(t, getErr)
	assert.NotNil(t, todos)
	assert.EqualValues(t, 1, len(todos))
	assert.EqualValues(t, 1, todos[0].Id)
	assert.EqualValues(t, "Test Todo", todos[0].Description)
}
