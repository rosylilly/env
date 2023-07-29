package env_test

import (
	"bytes"
	"net"
	"testing"

	"github.com/rosylilly/env"
)

func TestText(t *testing.T) {
	localIP := net.IPv4(127, 0, 0, 1)
	classC := net.IPv4(192, 168, 0, 1)

	testCases := []struct {
		input        string
		defaultValue net.IP
		expected     net.IP
	}{
		{"", localIP, localIP},
		{"invalid", localIP, localIP},
		{"192.168.0.1", localIP, classC},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.input, func(t *testing.T) {
			t.Setenv("IP", tc.input)

			actual := env.Text("IP", tc.defaultValue)
			if !bytes.EqualFold(actual, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}

	defaultValue := env.Text("IP", localIP)
	if !bytes.EqualFold(defaultValue, localIP) {
		t.Errorf("expected %v, got %v", net.IPv4(127, 0, 0, 1), defaultValue)
	}
}
