package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
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

// Delete removes a TODO at specified index from List
func (l *List) Delete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item %d does not exist", i)
	}

	*l = append(ls[:i-1], ls[i:]...)

	return nil
}

// Save encodes as json and saves to file
func (l *List) Save(filename string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, js, 0644)
}

// Get opens and decodes json into a List
func (l *List) Get(filename string) error {
	f, err := os.ReadFile(filename) // os.ReadFile takes care of closing the file
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}

		return err
	}

	if len(f) == 0 {
		return nil
	}

	return json.Unmarshal(f, l)
}

// String prints out a formatted list; implements the fmt.Stringer interface
func (l *List) String() string {
	formatted := ""
	for k, t := range *l {
		prefix := "\t"
		if t.Done {
			prefix = "X  "
		}
		formatted += fmt.Sprintf("%s%d: %s\n", prefix, k+1, t.Task)
	}

	return formatted
}