package app

import "github.com/aasimsajjad22/go-todo-backend/controllers"

func mapUrls() {
	router.POST("/todo", controllers.TodoController.Create)
	router.GET("/todos", controllers.TodoController.GetAll)
}
