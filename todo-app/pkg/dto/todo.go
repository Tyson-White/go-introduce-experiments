package dto

type CreateTodo struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type MarkTodo struct {
	Id int `json:id""`
}
