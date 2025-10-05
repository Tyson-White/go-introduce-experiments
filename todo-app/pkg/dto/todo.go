package dto

type CreateTodo struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}
