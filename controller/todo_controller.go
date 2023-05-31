package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/mahfuzon/to-do-list-app/helper"
	"github.com/mahfuzon/to-do-list-app/request/todo_request"
	"github.com/mahfuzon/to-do-list-app/response"
	"github.com/mahfuzon/to-do-list-app/service"
)

type TodoController struct {
	todoService service.TodoService
}

func NewTodoController(todoService service.TodoService) *TodoController {
	return &TodoController{todoService: todoService}
}

func (todoController *TodoController) Create(ctx echo.Context) error {
	request := todo_request.Create{}

	err := ctx.Bind(&request)
	if err != nil {
		apiResponse := response.NewApiResponse("error", "failed create todo", err.Error())
		return ctx.JSON(500, apiResponse)
	}

	err = ctx.Validate(&request)
	if err != nil {
		errorMessage := helper.ConverseToErrorString(err.(validator.ValidationErrors))
		apiResponse := response.NewApiResponse("error", "failed create todo", errorMessage)
		return ctx.JSON(422, apiResponse)
	}

	res, err := todoController.todoService.Create(request)
	if err != nil {
		apiResponse := response.NewApiResponse("error", "failed create todo", err.Error())
		return ctx.JSON(400, apiResponse)
	}

	apiResponse := response.NewApiResponse("success", "success create todo", res)
	return ctx.JSON(201, apiResponse)
}

func (todoController *TodoController) Update(ctx echo.Context) error {
	request := todo_request.Update{}

	err := ctx.Bind(&request)
	if err != nil {
		apiResponse := response.NewApiResponse("error", "failed update todo", err.Error())
		return ctx.JSON(500, apiResponse)
	}

	err = ctx.Validate(&request)
	if err != nil {
		errorMessage := helper.ConverseToErrorString(err.(validator.ValidationErrors))
		apiResponse := response.NewApiResponse("error", "failed update todo", errorMessage)
		return ctx.JSON(422, apiResponse)
	}

	res, err := todoController.todoService.Update(request)
	if err != nil {
		apiResponse := response.NewApiResponse("error", "failed update todo", err.Error())
		return ctx.JSON(400, apiResponse)
	}

	apiResponse := response.NewApiResponse("success", "success update todo", res)
	return ctx.JSON(201, apiResponse)
}

func (todoController *TodoController) GetById(ctx echo.Context) error {
	request := todo_request.Find{}

	err := ctx.Bind(&request)
	if err != nil {
		apiResponse := response.NewApiResponse("error", "failed get detail todo", err.Error())
		return ctx.JSON(500, apiResponse)
	}

	err = ctx.Validate(&request)
	if err != nil {
		errorMessage := helper.ConverseToErrorString(err.(validator.ValidationErrors))
		apiResponse := response.NewApiResponse("error", "failed get detail todo", errorMessage)
		return ctx.JSON(422, apiResponse)
	}

	res, err := todoController.todoService.Read(request)
	if err != nil {
		apiResponse := response.NewApiResponse("error", "failed get detail todo", err.Error())
		return ctx.JSON(400, apiResponse)
	}

	apiResponse := response.NewApiResponse("success", "success get detail todo", res)
	return ctx.JSON(201, apiResponse)
}

func (todoController *TodoController) GetAll(ctx echo.Context) error {
	res, err := todoController.todoService.ReadAll()
	if err != nil {
		apiResponse := response.NewApiResponse("error", "failed get all todo", err.Error())
		return ctx.JSON(400, apiResponse)
	}

	apiResponse := response.NewApiResponse("success", "success get all todo", res)
	return ctx.JSON(201, apiResponse)
}

func (todoController *TodoController) Delete(ctx echo.Context) error {
	request := todo_request.Delete{}

	err := ctx.Bind(&request)
	if err != nil {
		apiResponse := response.NewApiResponse("error", "failed delete todo", err.Error())
		return ctx.JSON(500, apiResponse)
	}

	err = ctx.Validate(&request)
	if err != nil {
		errorMessage := helper.ConverseToErrorString(err.(validator.ValidationErrors))
		apiResponse := response.NewApiResponse("error", "failed delete todo", errorMessage)
		return ctx.JSON(422, apiResponse)
	}

	err = todoController.todoService.Delete(request)
	if err != nil {
		apiResponse := response.NewApiResponse("error", "failed delete todo", err.Error())
		return ctx.JSON(400, apiResponse)
	}

	apiResponse := response.NewApiResponse("success", "success delete todo", nil)
	return ctx.JSON(201, apiResponse)
}
