package env_test

import (
	"fmt"
	"testing"

	"github.com/rosylilly/env"
)

func TestBool(t *testing.T) {
	testCases := []struct {
		input        string
		defaultValue bool
		expected     bool
	}{
		{"", false, false},
		{"0", false, false},
		{"1", false, true},
		{"false", false, false},
		{"true", false, true},
		{"False", false, false},
		{"True", false, true},
		{"FALSE", false, false},
		{"TRUE", false, true},
		{"f", false, false},
		{"t", false, true},
		{"F", false, false},
		{"T", false, true},
	}

	for i, tc := range testCases {
		tc := tc

		t.Run(fmt.Sprintf("#%02d", i), func(t *testing.T) {
			t.Setenv("TEST_BOOL", tc.input)

			if actual := env.Bool("TEST_BOOL", tc.defaultValue); actual != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}

	defaultValue := env.Bool("TEST_BOOL", true)
	if defaultValue != true {
		t.Errorf("expected %v, got %v", true, defaultValue)
	}
}

func TestString(t *testing.T) {
	empty := ""
	value := "value"

	testCases := []struct {
		input        *string
		defaultValue string
		expected     string
	}{
		{&empty, "", ""},
		{&empty, "default", ""},
		{nil, "default", "default"},
		{&value, "", "value"},
	}

	for i, tc := range testCases {
		tc := tc

		t.Run(fmt.Sprintf("#%02d", i), func(t *testing.T) {
			if tc.input != nil {
				t.Setenv("TEST_STRING", *tc.input)
			}

			if actual := env.String("TEST_STRING", tc.defaultValue); actual != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func TestInt(t *testing.T) {
	testCases := []struct {
		input        string
		defaultValue int
		expected     int
	}{
		{"", 0, 0},
		{"", 1, 1},
		{"0", 0, 0},
		{"0", 1, 0},
		{"1", 0, 1},
		{"1", 2, 1},
		{"-1", 0, -1},
		{"-1", 2, -1},
		{"a", 0, 0},
		{"a", 2, 2},
	}

	for i, tc := range testCases {
		tc := tc

		t.Run(fmt.Sprintf("#%02d", i), func(t *testing.T) {
			t.Setenv("TEST_INT", tc.input)

			if actual := env.Int("TEST_INT", tc.defaultValue); actual != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}

	defaultValue := env.Int("TEST_INT", 1)
	if defaultValue != 1 {
		t.Errorf("expected %v, got %v", 1, defaultValue)
	}
}

func TestInt8(t *testing.T) {
	testCases := []struct {
		input        string
		defaultValue int8
		expected     int8
	}{
		{"", 0, 0},
		{"", 1, 1},
		{"0", 0, 0},
		{"0", 1, 0},
		{"1", 0, 1},
		{"1", 2, 1},
		{"-1", 0, -1},
		{"-1", 2, -1},
		{"a", 0, 0},
		{"a", 2, 2},
		{"128", 0, 0},
		{"128", 2, 2},
		{"-129", 0, 0},
		{"-129", 2, 2},
	}

	for i, tc := range testCases {
		tc := tc

		t.Run(fmt.Sprintf("#%02d", i), func(t *testing.T) {
			t.Setenv("TEST_INT8", tc.input)

			if actual := env.Int8("TEST_INT8", tc.defaultValue); actual != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}

	defaultValue := env.Int8("TEST_INT8", 1)
	if defaultValue != 1 {
		t.Errorf("expected %v, got %v", 1, defaultValue)
	}
}

func TestInt16(t *testing.T) {
	testCases := []struct {
		input        string
		defaultValue int16
		expected     int16
	}{
		{"", 0, 0},
		{"", 1, 1},
		{"0", 0, 0},
		{"0", 1, 0},
		{"1", 0, 1},
		{"1", 2, 1},
		{"-1", 0, -1},
		{"-1", 2, -1},
		{"a", 0, 0},
		{"a", 2, 2},
		{"32768", 0, 0},
		{"32768", 2, 2},
		{"-32769", 0, 0},
		{"-32769", 2, 2},
	}

	for i, tc := range testCases {
		tc := tc

		t.Run(fmt.Sprintf("#%02d", i), func(t *testing.T) {
			t.Setenv("TEST_INT16", tc.input)

			if actual := env.Int16("TEST_INT16", tc.defaultValue); actual != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}

	defaultValue := env.Int16("TEST_INT16", 1)
	if defaultValue != 1 {
		t.Errorf("expected %v, got %v", 1, defaultValue)
	}
}

func TestInt32(t *testing.T) {
	testCases := []struct {
		input        string
		defaultValue int32
		expected     int32
	}{
		{"", 0, 0},
		{"", 1, 1},
		{"0", 0, 0},
		{"0", 1, 0},
		{"1", 0, 1},
		{"1", 2, 1},
		{"-1", 0, -1},
		{"-1", 2, -1},
		{"a", 0, 0},
		{"a", 2, 2},
		{"2147483648", 0, 0},
		{"2147483648", 2, 2},
		{"-2147483649", 0, 0},
		{"-2147483649", 2, 2},
	}

	for i, tc := range testCases {
		tc := tc

		t.Run(fmt.Sprintf("#%02d", i), func(t *testing.T) {
			t.Setenv("TEST_INT32", tc.input)

			if actual := env.Int32("TEST_INT32", tc.defaultValue); actual != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}

	defaultValue := env.Int32("TEST_INT32", 1)
	if defaultValue != 1 {
		t.Errorf("expected %v, got %v", 1, defaultValue)
	}
}

func TestInt64(t *testing.T) {
	testCases := []struct {
		input        string
		defaultValue int64
		expected     int64
	}{
		{"", 0, 0},
		{"", 1, 1},
		{"0", 0, 0},
		{"0", 1, 0},
		{"1", 0, 1},
		{"1", 2, 1},
		{"-1", 0, -1},
		{"-1", 2, -1},
		{"a", 0, 0},
		{"a", 2, 2},
		{"9223372036854775808", 0, 0},
		{"9223372036854775808", 2, 2},
		{"-9223372036854775809", 0, 0},
		{"-9223372036854775809", 2, 2},
	}

	for i, tc := range testCases {
		tc := tc

		t.Run(fmt.Sprintf("#%02d", i), func(t *testing.T) {
			t.Setenv("TEST_INT64", tc.input)

			if actual := env.Int64("TEST_INT64", tc.defaultValue); actual != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}

	defaultValue := env.Int64("TEST_INT64", 1)
	if defaultValue != 1 {
		t.Errorf("expected %v, got %v", 1, defaultValue)
	}
}

func TestUint(t *testing.T) {
	testCases := []struct {
		input        string
		defaultValue uint
		expected     uint
	}{
		{"", 0, 0},
		{"", 1, 1},
		{"0", 0, 0},
		{"0", 1, 0},
		{"1", 0, 1},
		{"1", 2, 1},
		{"a", 0, 0},
		{"a", 2, 2},
		{"-1", 0, 0},
		{"-1", 2, 2},
		{"-4294967297", 2, 2},
	}

	for i, tc := range testCases {
		tc := tc

		t.Run(fmt.Sprintf("#%02d", i), func(t *testing.T) {
			t.Setenv("TEST_UINT", tc.input)

			if actual := env.Uint("TEST_UINT", tc.defaultValue); actual != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}

	defaultValue := env.Uint("TEST_UINT", 1)
	if defaultValue != 1 {
		t.Errorf("expected %v, got %v", 1, defaultValue)
	}
}

func TestUint8(t *testing.T) {
	testCases := []struct {
		input        string
		defaultValue uint8
		expected     uint8
	}{
		{"", 0, 0},
		{"", 1, 1},
		{"0", 0, 0},
		{"0", 1, 0},
		{"1", 0, 1},
		{"1", 2, 1},
		{"a", 0, 0},
		{"a", 2, 2},
		{"256", 0, 0},
		{"256", 2, 2},
		{"-1", 0, 0},
		{"-1", 2, 2},
		{"-257", 0, 0},
		{"-257", 2, 2},
	}

	for i, tc := range testCases {
		tc := tc

		t.Run(fmt.Sprintf("#%02d", i), func(t *testing.T) {
			t.Setenv("TEST_UINT8", tc.input)

			if actual := env.Uint8("TEST_UINT8", tc.defaultValue); actual != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}

	defaultValue := env.Uint8("TEST_UINT8", 1)
	if defaultValue != 1 {
		t.Errorf("expected %v, got %v", 1, defaultValue)
	}
}

func TestUint16(t *testing.T) {
	testCases := []struct {
		input        string
		defaultValue uint16
		expected     uint16
	}{
		{"", 0, 0},
		{"", 1, 1},
		{"0", 0, 0},
		{"0", 1, 0},
		{"1", 0, 1},
		{"1", 2, 1},
		{"a", 0, 0},
		{"a", 2, 2},
		{"65536", 0, 0},
		{"65536", 2, 2},
		{"-1", 0, 0},
		{"-1", 2, 2},
		{"-65537", 0, 0},
		{"-65537", 2, 2},
	}

	for i, tc := range testCases {
		tc := tc

		t.Run(fmt.Sprintf("#%02d", i), func(t *testing.T) {
			t.Setenv("TEST_UINT16", tc.input)

			if actual := env.Uint16("TEST_UINT16", tc.defaultValue); actual != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}

	defaultValue := env.Uint16("TEST_UINT16", 1)
	if defaultValue != 1 {
		t.Errorf("expected %v, got %v", 1, defaultValue)
	}
}

func TestUint32(t *testing.T) {
	testCases := []struct {
		input        string
		defaultValue uint32
		expected     uint32
	}{
		{"", 0, 0},
		{"", 1, 1},
		{"0", 0, 0},
		{"0", 1, 0},
		{"1", 0, 1},
		{"1", 2, 1},
		{"a", 0, 0},
		{"a", 2, 2},
		{"4294967296", 0, 0},
		{"4294967296", 2, 2},
		{"-1", 0, 0},
		{"-1", 2, 2},
		{"-4294967297", 0, 0},
		{"-4294967297", 2, 2},
	}

	for i, tc := range testCases {
		tc := tc

		t.Run(fmt.Sprintf("#%02d", i), func(t *testing.T) {
			t.Setenv("TEST_UINT32", tc.input)

			if actual := env.Uint32("TEST_UINT32", tc.defaultValue); actual != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}

	defaultValue := env.Uint32("TEST_UINT32", 1)
	if defaultValue != 1 {
		t.Errorf("expected %v, got %v", 1, defaultValue)
	}
}

func TestUint64(t *testing.T) {
	testCases := []struct {
		input        string
		defaultValue uint64
		expected     uint64
	}{
		{"", 0, 0},
		{"", 1, 1},
		{"0", 0, 0},
		{"0", 1, 0},
		{"1", 0, 1},
		{"1", 2, 1},
		{"a", 0, 0},
		{"a", 2, 2},
		{"18446744073709551616", 0, 0},
		{"18446744073709551616", 2, 2},
		{"-1", 0, 0},
		{"-1", 2, 2},
		{"-18446744073709551617", 0, 0},
		{"-18446744073709551617", 2, 2},
	}

	for i, tc := range testCases {
		tc := tc

		t.Run(fmt.Sprintf("#%02d", i), func(t *testing.T) {
			t.Setenv("TEST_UINT64", tc.input)

			if actual := env.Uint64("TEST_UINT64", tc.defaultValue); actual != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}

	defaultValue := env.Uint64("TEST_UINT64", 1)
	if defaultValue != 1 {
		t.Errorf("expected %v, got %v", 1, defaultValue)
	}
}

func TestFloat32(t *testing.T) {
	testCases := []struct {
		input        string
		defaultValue float32
		expected     float32
	}{
		{"", 0, 0},
		{"", 1, 1},
		{"0", 0, 0},
		{"0", 1, 0},
		{"1", 0, 1},
		{"1", 2, 1},
		{"a", 0, 0},
		{"a", 2, 2},
		{"-1", 0, -1},
		{"-1", 2, -1},
		{"-1.1", 0, -1.1},
		{"-1.1", 2, -1.1},
		{"1.1", 0, 1.1},
		{"1.1", 2, 1.1},
		{"1.1.1", 0, 0},
		{"1.1.1", 2, 2},
	}

	for i, tc := range testCases {
		tc := tc

		t.Run(fmt.Sprintf("#%02d", i), func(t *testing.T) {
			t.Setenv("TEST_FLOAT32", tc.input)

			if actual := env.Float32("TEST_FLOAT32", tc.defaultValue); actual != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}

	defaultValue := env.Float32("TEST_FLOAT32", 1)
	if defaultValue != 1 {
		t.Errorf("expected %v, got %v", 1, defaultValue)
	}
}

func TestFloat64(t *testing.T) {
	testCases := []struct {
		input        string
		defaultValue float64
		expected     float64
	}{
		{"", 0, 0},
		{"", 1, 1},
		{"0", 0, 0},
		{"0", 1, 0},
		{"1", 0, 1},
		{"1", 2, 1},
		{"a", 0, 0},
		{"a", 2, 2},
		{"-1", 0, -1},
		{"-1", 2, -1},
		{"-1.1", 0, -1.1},
		{"-1.1", 2, -1.1},
		{"1.1", 0, 1.1},
		{"1.1", 2, 1.1},
		{"1.1.1", 0, 0},
		{"1.1.1", 2, 2},
	}

	for i, tc := range testCases {
		tc := tc

		t.Run(fmt.Sprintf("#%02d", i), func(t *testing.T) {
			t.Setenv("TEST_FLOAT64", tc.input)

			if actual := env.Float64("TEST_FLOAT64", tc.defaultValue); actual != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}

	defaultValue := env.Float64("TEST_FLOAT64", 1)
	if defaultValue != 1 {
		t.Errorf("expected %v, got %v", 1, defaultValue)
	}
}
