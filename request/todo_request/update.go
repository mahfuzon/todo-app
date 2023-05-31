package todo_request

type Update struct {
	Id    int    `param:"id" validate:"required"`
	Title string `json:"Title" validate:"required"`
}
