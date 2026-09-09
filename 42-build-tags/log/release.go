//go:build !debug && !test

package log

// this file will be included for release builds only

func DebugMessage(msg string) {
	// do nothing in release mode
	println("just to show this is being called")
}
