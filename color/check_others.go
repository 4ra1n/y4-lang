//go:build !windows

package color

func isWindowsColorSupported() bool {
	return true
}
