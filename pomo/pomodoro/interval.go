package pomodoro

import (
	"errors"
	"time"
)

// Category constants
const (
	CategoryPomodoro   = "Pomodoro"
	CategoryShortBreak = "ShortBreak"
	CategoryLongBreak  = "LongBreak"
)

// State constants
const (
	StateNotStarted = iota
	StateRunning
	StatePaused
	StateDone
	StateCancelled
)

// Interval represents a pomodoro interval
type Interval struct {
	ID              int64
	StartTime       time.Time
	PlannedDuration time.Duration
	ActualDuration  time.Duration
	Categroy        string
	State           int
}

type Repository interface {
	Create(i Interval) (int64, error)
	Update(i Interval) error
	ByID(id int64) (Interval, error)
	Last() (Interval, error)
	Breaks(n int) ([]Interval, error)
}

var (
	ErrNoIntervals        = errors.New("no intervals")
	ErrIntervalNotRunning = errors.New("interval not running")
	ErrIntervalCompleted  = errors.New("interval is completed or cancelled")
	ErrInvalidState       = errors.New("invalid state")
	ErrInvalidID          = errors.New("invalid ID")
)

type IntervalConfig struct {
	repo               Repository
	PomodoroDuration   time.Duration
	ShortBreakDuration time.Duration
	LongBreakDuration  time.Duration
}

func NewConfig(repo Repository, pomodoro, shortBreak, longBreak time.Duration) *IntervalConfig {
	c := &IntervalConfig{
		repo: repo,
		PomodoroDuration: 25 * time.Minute,
		ShortBreakDuration: 5 * time.Minute,
		LongBreakDuration: 15 * time.Minute,
	}

	if pomodoro > 0 {
		c.PomodoroDuration = pomodoro
	}

	if shortBreak > 0 {
		c.ShortBreakDuration = shortBreak
	}

	if longBreak > 0 {
		c.LongBreakDuration = longBreak
	}

	return c
}