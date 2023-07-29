package env

import (
	"encoding"
	"os"
)

func Text[T any](key string, defaultValue T) T {
	v, ok := os.LookupEnv(key)
	if !ok || v == "" {
		return defaultValue
	}

	var (
		tv interface{} = new(T)
	)
	if iv, ok := tv.(encoding.TextUnmarshaler); ok {
		if err := iv.UnmarshalText([]byte(v)); err == nil {
			return *tv.(*T)
		}
	}

	return defaultValue
}
