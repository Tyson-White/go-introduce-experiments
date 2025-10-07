package models

import "time"

type Todo struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	Text       string    `json:"text"`
	Time       time.Time `json:"time"`
	Completed  bool      `json:"completed"`
	CategoryId *int      `json:"category_id" db:"category_id"`
}
