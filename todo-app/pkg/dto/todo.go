package dto

import "time"

type CreateTodo struct {
	Title       string    `json:"title"`
	Text        string    `json:"text"`
	ExpiresTime time.Time `json:"expires_time"`
}

type MarkTodo struct {
	Id int `json:"id"`
}
