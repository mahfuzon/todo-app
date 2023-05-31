package todo_response

import "github.com/mahfuzon/to-do-list-app/model"

type TodoResponse struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

func NewTodoResponse(todo model.Todo) TodoResponse {
	return TodoResponse{
		Id:    todo.Id,
		Title: todo.Title,
	}
}

func NewListTodoResponse(todos []model.Todo) []TodoResponse {
	var listTodoResponse []TodoResponse
	if len(todos) > 0 {
		for _, todo := range todos {
			todoResponse := NewTodoResponse(todo)
			listTodoResponse = append(listTodoResponse, todoResponse)
		}
	}

	return listTodoResponse
}
