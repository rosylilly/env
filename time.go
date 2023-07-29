package env

import (
	"os"
	"time"
)

func Time(key string, defaultValue time.Time) time.Time {
	v, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	if tv, err := time.Parse(time.RFC3339, v); err == nil {
		return tv
	}

	return defaultValue
}

func Duration(key string, defaultValue time.Duration) time.Duration {
	v, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	if dv, err := time.ParseDuration(v); err == nil {
		return dv
	}

	return defaultValue
}
