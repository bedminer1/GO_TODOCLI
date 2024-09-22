package cmd

import (
	"github.com/bedminer1/pomo/pomodoro"
	"github.com/bedminer1/pomo/pomodoro/repository"
)

func getRepo() (pomodoro.Repository, error) {
	return repository.NewInMemoryRepo(), nil
}
