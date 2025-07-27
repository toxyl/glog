package ansi

import (
	"io"
	"os"
)

// ANSI represents an ANSI escape sequence
type ANSI struct {
	sequence string
}

// New creates a new ANSI escape sequence
func New(sequence string) *ANSI {
	return &ANSI{sequence: sequence}
}

// String returns the ANSI escape sequence as a string
func (a *ANSI) String() string {
	return a.sequence
}

// ToStdout writes the ANSI escape sequence to stdout
func (a *ANSI) ToStdout() {
	_, _ = a.Write(os.Stdout)
}

// ToStderr writes the ANSI escape sequence to stderr
func (a *ANSI) ToStderr() {
	_, _ = a.Write(os.Stderr)
}

// Write writes the ANSI escape sequence to the given writer
func (a *ANSI) Write(w io.Writer) (int, error) {
	return w.Write([]byte(a.sequence))
}

// Combine combines multiple ANSI sequences
func (a *ANSI) Combine(others ...*ANSI) *ANSI {
	combined := a.sequence
	for _, other := range others {
		combined += other.sequence
	}
	return New(combined)
}

// Apply applies the ANSI sequence to a string
func (a *ANSI) Apply(text string) string {
	return a.sequence + text + Reset().sequence
}
