package main

import (
	"github.com/joho/godotenv"
	"github.com/mahfuzon/to-do-list-app/controller"
	libraries "github.com/mahfuzon/to-do-list-app/library"
	"github.com/mahfuzon/to-do-list-app/repository"
	"github.com/mahfuzon/to-do-list-app/service"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	db := libraries.SetDb()
	todoRepository := repository.NewTodoRepository(db)
	todoService := service.NewTodoService(todoRepository)
	todoController := controller.NewTodoController(todoService)

	router := libraries.SetRouter()

	api := router.Group("/api")

	apiV1 := api.Group("/v1")

	apiV1Todo := apiV1.Group("/todo")
	apiV1Todo.POST("", todoController.Create)
	apiV1Todo.DELETE("/:id", todoController.Delete)
	apiV1Todo.GET("/:id", todoController.GetById)
	apiV1Todo.GET("", todoController.GetAll)
	apiV1Todo.PUT("/:id", todoController.Update)

	router.Logger.Fatal(router.Start(":8000"))
}