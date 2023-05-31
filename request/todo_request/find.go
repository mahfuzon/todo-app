package todo_request

type Find struct {
	Id int `param:"id" validate:"required"`
}
