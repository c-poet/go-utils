// Package os provides operating-system detection helpers.
package os

import "runtime"

const (
	OSNameWin   = "windows"
	OSNameMac   = "darwin"
	OSNameLinux = "linux"
)

// IsWindows reports whether the current program is running on Windows.
func IsWindows() bool {
	return runtime.GOOS == OSNameWin
}

// IsMac reports whether the current program is running on macOS.
func IsMac() bool {
	return runtime.GOOS == OSNameMac
}

// IsLinux reports whether the current program is running on Linux.
func IsLinux() bool {
	return runtime.GOOS == OSNameLinux
}
