package pomodoro_test

import (
	"testing"

	"github.com/bedminer1/pomo/pomodoro"
	"github.com/bedminer1/pomo/pomodoro/repository"
)

func getRepo(t *testing.T) (pomodoro.Repository, func()) {
	t.Helper()
	// inmemory repo doesnt require a cleanup function so
	// empty func is returned
	return repository.NewInMemoryRepo(), func() {}
}
