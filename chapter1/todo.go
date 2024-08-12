package todo

import "time"

// item is a TODO item
type item struct {
	Task string
	Done bool
	CreatedAt time.Time
	CompletedAt time.Time
}

// List represents a list of TODO items
type List []item
