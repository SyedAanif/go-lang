//go:build debug || test

package log

// this file will be included for testing or debug builds only

func DebugMessage(msg string) {
	// print debug message
	println("[DEBUG]", msg)
}
