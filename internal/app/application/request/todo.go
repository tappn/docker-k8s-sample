package request

import "time"

type Todo struct {
	Title       string    `json:"title" validate:"min=1,max=50"`
	Description string    `json:"description" validate:"min=1,max=250"`
	Deadline    time.Time `json:"deadline"`
}
