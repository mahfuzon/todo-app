package model

import "time"

type Todo struct {
	Id        int
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
