package pomodoro_test

import (
	"testing"
	"time"

	"github.com/bedminer1/pomo/pomodoro"
)

func TestNewConfig(t *testing.T) {
	testCases := []struct {
		name string
		input [3]time.Duration
		expect pomodoro.IntervalConfig
	}{
		{
			name: "Default",
			expect: pomodoro.IntervalConfig{
				PomodoroDuration: 25 * time.Minute,
				ShortBreakDuration: 5 * time.Minute,
				LongBreakDuration: 15 * time.Minute,
			},
		},
		{
			name: "SingleInput",
			input: [3]time.Duration{
				20 * time.Minute,
			},
			expect: pomodoro.IntervalConfig{
				PomodoroDuration: 20 * time.Minute,
				ShortBreakDuration: 5 * time.Minute,
				LongBreakDuration: 15 * time.Minute,
			},
		},
		{
			name: "MultiInput",
			input: [3]time.Duration{
				20 * time.Minute,
				10 * time.Minute,
				12 * time.Minute,
			},
			expect: pomodoro.IntervalConfig{
				PomodoroDuration: 20 * time.Minute,
				ShortBreakDuration: 10 * time.Minute,
				LongBreakDuration: 12 * time.Minute,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var repo pomodoro.Repository
			config := pomodoro.NewConfig(
				repo,
				tc.input[0],
				tc.input[1],
				tc.input[2],
			)

			if config.PomodoroDuration != tc.expect.PomodoroDuration {
				t.Error("Unexpected pomodoro duration")
			}
			if config.ShortBreakDuration != tc.expect.ShortBreakDuration {
				t.Error("Unexpected short break duration")
			}
			if config.LongBreakDuration != tc.expect.LongBreakDuration {
				t.Error("Unexpected long break duration")
			}
		})
	}
}
