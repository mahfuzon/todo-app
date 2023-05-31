package service

import (
	"github.com/mahfuzon/to-do-list-app/model"
	"github.com/mahfuzon/to-do-list-app/repository"
	"github.com/mahfuzon/to-do-list-app/request/todo_request"
	"github.com/mahfuzon/to-do-list-app/response/todo_response"
)

type TodoService interface {
	Create(request todo_request.Create) (todo_response.TodoResponse, error)
	Read(request todo_request.Find) (todo_response.TodoResponse, error)
	ReadAll() ([]todo_response.TodoResponse, error)
	Update(request todo_request.Update) (todo_response.TodoResponse, error)
	Delete(request todo_request.Delete) error
}

type todoService struct {
	todoRepository repository.TodoRepository
}

func NewTodoService(todoRepository repository.TodoRepository) TodoService {
	return &todoService{todoRepository: todoRepository}
}

func (todoService *todoService) Create(request todo_request.Create) (todo_response.TodoResponse, error) {
	todo := model.Todo{
		Title: request.Title,
	}

	todo, err := todoService.todoRepository.Create(todo)
	if err != nil {
		return todo_response.TodoResponse{}, err
	}

	return todo_response.NewTodoResponse(todo), nil
}

func (todoService *todoService) Read(request todo_request.Find) (todo_response.TodoResponse, error) {
	todo, err := todoService.todoRepository.Find(request.Id)
	if err != nil {
		return todo_response.TodoResponse{}, err
	}

	return todo_response.NewTodoResponse(todo), nil
}

func (todoService *todoService) ReadAll() ([]todo_response.TodoResponse, error) {
	todos, err := todoService.todoRepository.All()
	if err != nil {
		return []todo_response.TodoResponse{}, err
	}

	return todo_response.NewListTodoResponse(todos), nil
}

func (todoService *todoService) Update(request todo_request.Update) (todo_response.TodoResponse, error) {
	todo, err := todoService.todoRepository.Find(request.Id)
	if err != nil {
		return todo_response.TodoResponse{}, err
	}

	todo.Title = request.Title

	todo, err = todoService.todoRepository.Update(todo)
	if err != nil {
		return todo_response.TodoResponse{}, err
	}

	return todo_response.NewTodoResponse(todo), nil
}

func (todoService *todoService) Delete(request todo_request.Delete) error {
	todo, err := todoService.todoRepository.Find(request.Id)
	if err != nil {
		return err
	}

	err = todoService.todoRepository.Delete(todo)
	if err != nil {
		return err
	}

	return nil
}
