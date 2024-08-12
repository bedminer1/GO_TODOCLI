package todo

import (
	"fmt"
	"time"
)

// item is a TODO item
type item struct {
	Task string
	Done bool
	CreatedAt time.Time
	CompletedAt time.Time
}

// List represents a list of TODO items
type List []item

// Add creates a new todo item n appends to List
func (l *List) Add(task string) {
	t := item {
		Task: task,
		Done: false,
		CreatedAt: time.Now(),
		CompletedAt: time.Time{},
	}

	*l = append(*l, t)
}

// Complete marks a TODO item as complete
func (l *List) Complete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item %d does not exist", i)
	}

	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()

	return nil
}