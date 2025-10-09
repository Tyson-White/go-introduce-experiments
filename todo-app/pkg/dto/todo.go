package dto

import "time"

type CreateTodo struct {
	Title       string    `json:"title"`
	Text        string    `json:"text"`
	ExpiresTime time.Time `json:"expires_time"`
	Category    *int      `json:"category"`
}

type MarkTodo struct {
	Id int `json:"id"`
}

type DeleteTodoBody struct {
	Todo int `json:"todo"`
}
