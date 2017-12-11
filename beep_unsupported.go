// +build !linux,!windows,!darwin,!js

package beeep

import (
	"errors"
	"runtime"
)

var (
	// Default frequency, Hz, middle A
	DefaultFreq = 0.0
	// Default duration in milliseconds
	DefaultDuration = 0
)

// Beep beeps the pc speaker (https://en.wikipedia.org/wiki/PC_speaker).
func Beep(freq float64, duration int) error {
	return errors.New("beeep: unsupported operating system: %v", runtime.GOOS)
}