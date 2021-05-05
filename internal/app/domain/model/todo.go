package model

import "time"

type Todo struct {
	ID          int64
	Title       string
	Description string
	Deadline    time.Time
}
