package env

import (
	"encoding/json"
	"os"
)

func JSON[T any](key string, defaultValue T) T {
	v, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	var tv T
	if err := json.Unmarshal([]byte(v), &tv); err == nil {
		return tv
	}

	return defaultValue
}
