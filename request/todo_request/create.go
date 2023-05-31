package todo_request

type Create struct {
	Title string `json:"Title" validate:"required"`
}
