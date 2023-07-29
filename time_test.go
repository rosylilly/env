package env_test

import (
	"testing"
	"time"

	"github.com/rosylilly/env"
)

func TestTime(t *testing.T) {
	now := time.Now()

	testCases := []struct {
		input        string
		defaultValue time.Time
		expected     time.Time
	}{
		{"", now, now},
		{"2006-01-02T15:04:05Z", now, time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC)},
		{"/", now, now},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.input, func(t *testing.T) {
			t.Setenv("TIME", tc.input)

			actual := env.Time("TIME", tc.defaultValue)
			if actual != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}

	defaultValue := env.Time("TIME", now)
	if defaultValue != now {
		t.Errorf("expected %v, got %v", now, defaultValue)
	}
}

func TestDuration(t *testing.T) {
	testCases := []struct {
		input        string
		defaultValue time.Duration
		expected     time.Duration
	}{
		{"", 0, 0},
		{"1s", 0, time.Second},
		{"1m", 0, time.Minute},
		{"1h", 0, time.Hour},
		{"/", 0, 0},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.input, func(t *testing.T) {
			t.Setenv("DURATION", tc.input)

			actual := env.Duration("DURATION", tc.defaultValue)
			if actual != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}

	defaultValue := env.Duration("DURATION", 0)
	if defaultValue != 0 {
		t.Errorf("expected %v, got %v", 0, defaultValue)
	}
}
