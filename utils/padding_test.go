package utils

import "testing"

func TestGetPadLength(t *testing.T) {
	tests := []struct {
		input     string
		maxLength int
		expected  int
	}{
		{
			input:     "<nil>",
			maxLength: 10,
			expected:  5,
		},
		{
			input:     "hello world",
			maxLength: 20,
			expected:  9,
		},
		{
			input:     "\033[1;31mhello world\033[0m",
			maxLength: 20,
			expected:  9,
		},
		{
			input:     "hello 😃 world",
			maxLength: 20,
			expected:  6,
		},
		{
			input:     "\033[1;31mhello 😃 world\033[0m",
			maxLength: 20,
			expected:  6,
		},
		{
			input:     "\033[1;31mhello 😃 world \033[0;32m🌍!\033[0m",
			maxLength: 20,
			expected:  2,
		},
		{
			input:     "1.00 μm",
			maxLength: 20,
			expected:  13,
		},
	}

	for _, tt := range tests {
		got := getPadLength(tt.input, tt.maxLength, ' ')
		if got != tt.expected {
			t.Errorf("getPadLength(%q, %d) = %d, want %d", tt.input, tt.maxLength, got, tt.expected)
		}
	}
}
