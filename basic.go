package env

import (
	"os"
	"strconv"
)

func Bool(key string, defaultValue bool) bool {
	v, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	if bv, err := strconv.ParseBool(v); err == nil {
		return bv
	}

	return defaultValue
}

func String(key string, defaultValue string) string {
	v, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	return v
}

func Int(key string, defaultValue int) int {
	v, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	if iv, err := strconv.ParseInt(v, 10, 0); err == nil {
		return int(iv)
	}

	return defaultValue
}

func Int8(key string, defaultValue int8) int8 {
	v, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	if iv, err := strconv.ParseInt(v, 10, 8); err == nil {
		return int8(iv)
	}

	return defaultValue
}

func Int16(key string, defaultValue int16) int16 {
	v, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	if iv, err := strconv.ParseInt(v, 10, 16); err == nil {
		return int16(iv)
	}

	return defaultValue
}

func Int32(key string, defaultValue int32) int32 {
	v, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	if iv, err := strconv.ParseInt(v, 10, 32); err == nil {
		return int32(iv)
	}

	return defaultValue
}

func Int64(key string, defaultValue int64) int64 {
	v, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	if iv, err := strconv.ParseInt(v, 10, 64); err == nil {
		return iv
	}

	return defaultValue
}

func Uint(key string, defaultValue uint) uint {
	v, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	if iv, err := strconv.ParseUint(v, 10, 0); err == nil {
		return uint(iv)
	}

	return defaultValue
}

func Uint8(key string, defaultValue uint8) uint8 {
	v, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	if iv, err := strconv.ParseUint(v, 10, 8); err == nil {
		return uint8(iv)
	}

	return defaultValue
}

func Uint16(key string, defaultValue uint16) uint16 {
	v, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	if iv, err := strconv.ParseUint(v, 10, 16); err == nil {
		return uint16(iv)
	}

	return defaultValue
}

func Uint32(key string, defaultValue uint32) uint32 {
	v, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	if iv, err := strconv.ParseUint(v, 10, 32); err == nil {
		return uint32(iv)
	}

	return defaultValue
}

func Uint64(key string, defaultValue uint64) uint64 {
	v, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	if iv, err := strconv.ParseUint(v, 10, 64); err == nil {
		return iv
	}

	return defaultValue
}

func Float32(key string, defaultValue float32) float32 {
	v, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	if iv, err := strconv.ParseFloat(v, 32); err == nil {
		return float32(iv)
	}

	return defaultValue
}

func Float64(key string, defaultValue float64) float64 {
	v, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	if iv, err := strconv.ParseFloat(v, 64); err == nil {
		return iv
	}

	return defaultValue
}
