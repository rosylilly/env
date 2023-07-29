package env_test

import (
	"testing"

	"github.com/rosylilly/env"
)

type TestJSONStruct struct {
	A string `json:"a"`
	B int    `json:"b"`
}

func TestJSON(t *testing.T) {
	testCases := []struct {
		input        string
		defaultValue TestJSONStruct
		expected     TestJSONStruct
	}{
		{"", TestJSONStruct{}, TestJSONStruct{}},
		{`{"a":"a","b":1}`, TestJSONStruct{}, TestJSONStruct{A: "a", B: 1}},
		{"/", TestJSONStruct{}, TestJSONStruct{}},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.input, func(t *testing.T) {
			t.Setenv("JSON", tc.input)

			actual := env.JSON("JSON", tc.defaultValue)
			if actual != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}

	defaultValue := env.JSON("JSON", TestJSONStruct{})
	if defaultValue != (TestJSONStruct{}) {
		t.Errorf("expected %v, got %v", TestJSONStruct{}, defaultValue)
	}
}
