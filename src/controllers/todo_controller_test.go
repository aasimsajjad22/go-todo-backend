package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/aasimsajjad22/go-todo-backend/domain/todo"
	"github.com/aasimsajjad22/go-todo-backend/services"
	"github.com/aasimsajjad22/go-todo-backend/utils/errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

var (
	saveFunc   func(t todo.Todo) (*todo.Todo, *errors.RestErr)
	getAllFunc func() ([]todo.Todo, *errors.RestErr)
)

type todoServiceMock struct{}

func (*todoServiceMock) Save(t todo.Todo) (*todo.Todo, *errors.RestErr) {
	return saveFunc(t)
}

func (*todoServiceMock) GetAll() ([]todo.Todo, *errors.RestErr) {
	return getAllFunc()
}

/**
Testcases for create endpoint
*/
func TestTodoControllerCreateBindJsonError(t *testing.T) {
	response := httptest.NewRecorder()
	todoController := todoController{}
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodPost, "", nil)

	todoController.Create(c)
	restErr := errors.RestErr{}
	_ = json.Unmarshal(response.Body.Bytes(), &restErr)

	assert.NotNil(t, restErr)
	assert.Equal(t, http.StatusBadRequest, response.Code)
	assert.EqualValues(t, "invalid json body", restErr.Message)
}

func TestTodoControllerCreateError(t *testing.T) {
	response := httptest.NewRecorder()
	todoController := todoController{}
	jsonBytes, _ := json.Marshal(map[string]interface{}{"description": ""})

	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodPost, "", nil)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonBytes))

	todoController.Create(c)
	restErr := errors.RestErr{}
	_ = json.Unmarshal(response.Body.Bytes(), &restErr)

	assert.NotNil(t, restErr)
	assert.EqualValues(t, "please add description", restErr.Message)
	assert.EqualValues(t, http.StatusBadRequest, restErr.Status)
}

func TestTodoControllerCreateSuccess(t *testing.T) {
	saveFunc = func(t todo.Todo) (*todo.Todo, *errors.RestErr) {
		return &todo.Todo{Id: 1, Description: "Test Description"}, nil
	}
	response := httptest.NewRecorder()
	todoController := todoController{}
	jsonBytes, _ := json.Marshal(map[string]interface{}{"description": ""})

	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodPost, "", nil)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonBytes))

	services.TodoService = &todoServiceMock{}

	todoController.Create(c)
	var todo todo.Todo
	_ = json.Unmarshal(response.Body.Bytes(), &todo)

	assert.NotNil(t, todo)
	assert.EqualValues(t, 1, todo.Id)
	assert.EqualValues(t, "Test Description", todo.Description)
	assert.EqualValues(t, http.StatusCreated, response.Code)
}

/**
Testcases for get endpoint
*/
func TestTodoControllerGetAllError(t *testing.T) {
	getAllFunc = func() ([]todo.Todo, *errors.RestErr) {
		return nil, &errors.RestErr{
			Message: "unable to get todos",
			Status:  500,
			Error:   "internal_server_error",
		}
	}
	response := httptest.NewRecorder()
	todoController := todoController{}
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodGet, "", nil)

	services.TodoService = &todoServiceMock{}

	todoController.GetAll(c)
	restErr := errors.RestErr{}
	_ = json.Unmarshal(response.Body.Bytes(), &restErr)

	assert.NotNil(t, restErr)
	assert.EqualValues(t, "unable to get todos", restErr.Message)
	assert.EqualValues(t, 500, restErr.Status)
}

func TestTodoControllerGetAllSuccess(t *testing.T) {
	getAllFunc = func() ([]todo.Todo, *errors.RestErr) {
		todos := make([]todo.Todo, 0)
		todos = append(todos, todo.Todo{Id: 1, Description: "Todo 1"})
		todos = append(todos, todo.Todo{Id: 2, Description: "Todo 2"})
		return todos, nil
	}
	response := httptest.NewRecorder()
	todoController := todoController{}
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodGet, "", nil)

	services.TodoService = &todoServiceMock{}

	todoController.GetAll(c)

	todos := make([]todo.Todo, 0)
	_ = json.Unmarshal(response.Body.Bytes(), &todos)

	assert.NotNil(t, todos)
	assert.EqualValues(t, 2, len(todos))
	assert.EqualValues(t, 1, todos[0].Id)
	assert.EqualValues(t, "Todo 1", todos[0].Description)
}
