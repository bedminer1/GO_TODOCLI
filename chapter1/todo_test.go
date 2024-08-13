package todo_test

import (
	"testing"
	"github.com/bedminer1/chapter1todo"
)

// TestAdd tests the Add method
func TestAdd(t *testing.T) {
	l := todo.List{}
	taskName := "New Task"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("expected %q, got %q instead", taskName, l[0].Task)
	}
}