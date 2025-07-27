package utils

import (
	"testing"
)

func TestReplaceEmojis(t *testing.T) {
	tests := []struct {
		input    string
		replace  string
		expected string
	}{
		{
			input:    "This string has no emojis.",
			replace:  " ",
			expected: "This string has no emojis.",
		},
		{
			input:    "😃🚀👍💀",
			replace:  " ",
			expected: "    ",
		},
		{
			input:    "This string contains emojis 💀😃🚀👍",
			replace:  " ",
			expected: "This string contains emojis     ",
		},
		{
			input:    "This string contains non-printable characters \x02",
			replace:  " ",
			expected: "This string contains non-printable characters \x02",
		},
		{
			input:    "This string contains ANSI escape codes \x1b[31m\x1b[1mred and bold\x1b[0m.",
			replace:  " ",
			expected: "This string contains ANSI escape codes \x1b[31m\x1b[1mred and bold\x1b[0m.",
		},
	}

	for _, tt := range tests {
		output := ReplaceEmojis(tt.input, tt.replace)
		if output != tt.expected {
			t.Errorf("ReplaceEmojis(%q, %q): expected %q, but got %q", tt.input, tt.replace, tt.expected, output)
		}
	}
}

func TestStripANSI(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "This string has no ANSI escape codes.",
			expected: "This string has no ANSI escape codes.",
		},
		{
			input:    "\x1b[31m\x1b[1mThis string has only ANSI escape codes\x1b[0m.",
			expected: "This string has only ANSI escape codes.",
		},
		{
			input:    "This string has mixed \x1b[32m\x1b[1mred and bold\x1b[0m text.",
			expected: "This string has mixed red and bold text.",
		},
		{
			input:    "\x1b[33mThis string has multiple \x1b[33myellow\x1b[0m ANSI escape codes.",
			expected: "This string has multiple yellow ANSI escape codes.",
		},
		{
			input:    "This string contains ANSI escape codes \x1b[31m\x1b[1mred and bold\x1b[0m.",
			expected: "This string contains ANSI escape codes red and bold.",
		},
		{
			input:    "This string contains a non-printable character \x02.",
			expected: "This string contains a non-printable character \x02.",
		},
		{
			input:    "This string contains an emoji 😃.",
			expected: "This string contains an emoji 😃.",
		},
		{
			input:    "\033[1;31mhello world\033[0m",
			expected: "hello world",
		},
		{
			input:    "hello 😃 world",
			expected: "hello 😃 world",
		},
		{
			input:    "\033[1;31mhello 😃 world\033[0m",
			expected: "hello 😃 world",
		},
		{
			input:    "😃",
			expected: "😃",
		},
		{
			input:    "\x1b[1;31mhello 😃 world \033[0;32m🌍!\033[0m",
			expected: "hello 😃 world 🌍!",
		},
	}

	for _, test := range tests {
		output := StripANSI(test.input)
		if output != test.expected {
			t.Errorf("StripANSI: expected %q, but got %q", test.expected, output)
		}
	}
}

func TestRemoveNonPrintable(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "This string contains a non-printable character \x02.",
			expected: "This string contains a non-printable character .",
		},
		{
			input:    "This string contains an emoji 😃.",
			expected: "This string contains an emoji 😃.",
		},
		{
			input:    "This string contains a whitespace character \t and an emoji 🎉.",
			expected: "This string contains a whitespace character \t and an emoji 🎉.",
		},
		{
			input:    "This\x00string\x1b[32mis not\x1b[0m printable.",
			expected: "Thisstringis not printable.",
		},
		{
			input:    "\x0aThis is\na test.\r\n",
			expected: "\nThis is\na test.\r\n",
		},
		{
			input:    "\x00\x01\x02\x03\x04\x05\x06\x07\x08\t\x0b\x0c\x0e\x0f\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f",
			expected: "\t\v\f",
		},
	}

	for _, test := range tests {
		output := RemoveNonPrintable(test.input)
		if output != test.expected {
			t.Errorf("RemoveNonPrintable: expected %q, but got %q for input %q", test.expected, output, test.input)
		}
	}
}

func TestPlaintextString(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{
			input:    "",
			expected: "",
		},
		{
			input:    "\x1b[32mThis is green.\x1b[0m",
			expected: "This is green.",
		},
		{
			input:    "👋 Hello, 🌍!",
			expected: "👋 Hello, 🌍!",
		},
		{
			input:    "This\x00string\x1b[32mis not\x1b[0m printable.",
			expected: "Thisstringis not printable.",
		},
		{
			input:    "   This string has leading/trailing whitespace.  \n\t",
			expected: "   This string has leading/trailing whitespace.  \n\t",
		},
		{
			input:    "\x1b[32m👋\x1b[0m \x00Hello,\x1b[32m 🌍!\x1b[0m",
			expected: "👋 Hello, 🌍!",
		},
	}

	for _, tc := range testCases {
		output := PlaintextString(tc.input)
		if output != tc.expected {
			t.Errorf("PlaintextString(%q): expected %q, but got %q", tc.input, tc.expected, output)
		}
	}
}

func TestPlaintextStringLength(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			input:    "This is a normal string.",
			expected: 24,
		},
		{
			input:    "This string contains ANSI escape codes \x1b[31m\x1b[1mred and bold\x1b[0m.",
			expected: 52,
		},
		{
			input:    "This string contains non-printable characters \x00\x07.",
			expected: 47,
		},
		{
			input:    "This string contains non-printable and whitespace\t characters \x00\x07.",
			expected: 63,
		},
		{
			input:    "This string contains both ANSI escape codes \x1b[32mand \x1b[31mnon-printable\x1b[0m characters.",
			expected: 73,
		},
		{
			input:    "This string contains emojis 😃🚀👍.",
			expected: 32,
		},
		{
			input:    "This string contains a mixture of ANSI escape codes \x1b[33mand emojis 😃🚀.",
			expected: 66,
		},
		{
			input:    "\033[1;31mhello world\033[0m",
			expected: 11,
		},
		{
			input:    "hello 😃 world",
			expected: 13,
		},
		{
			input:    "\033[1;31mhello 😃 world\033[0m",
			expected: 13,
		},
		{
			input:    "\x1b[1;31mhello 😃 world \033[0;32m🌍!\033[0m",
			expected: 16,
		},
	}

	for _, test := range tests {
		output := PlaintextStringLength(test.input)
		if output != test.expected {
			t.Errorf("PlaintextStringLength(%q): expected %d, but got %d", test.input, test.expected, output)
		}
	}
}
