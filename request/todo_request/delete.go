package todo_request

type Delete struct {
	Id int `param:"id" validate:"required"`
}
