package main

import "buildTags/log"

// old syntax: +build <tag-name>
// new syntax: //go:build <tag-name>
// these are used to determine which files to include/exclude in the build based on tags
// also can be used to control platform dependant functions
// ex: //go:build linux will have syscall mount and unmount, but not on other platforms
// command: go run -tags=debug main.go -> debug message printed
// command: go run -tags=test main.go -> debug message printed
// command: go run main.go -> no debug message printed
func main() {
	log.DebugMessage("This is a test message. Will be printed only when debug or test tag")
}
