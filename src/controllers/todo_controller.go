package controllers

import (
	"github.com/aasimsajjad22/go-todo-backend/domain/todo"
	"github.com/aasimsajjad22/go-todo-backend/services"
	"github.com/aasimsajjad22/go-todo-backend/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	TodoController = todoController{}
)

type todoController struct{}

func (tc *todoController) Create(c *gin.Context) {
	var t todo.Todo
	if err := c.ShouldBindJSON(&t); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.TodoService.Save(t)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func (tc *todoController) GetAll(c *gin.Context) {
	todos, err := services.TodoService.GetAll()
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, todos)
}
