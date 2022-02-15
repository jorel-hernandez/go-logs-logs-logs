package logs

import (
	"strings"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	//panic("Please implement the Application() function")
	var app string
	for _, char := range log {
		if char == '❗' {
			app = "recommendation"
			break
		} else if char == '🔍' {
			app = "search"
			break
		} else if char == '☀' {
			app = "weather"
			break
		} else {
			app = "default"
		}
	}
	return app
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	//panic("Please implement the Replace() function")
	return strings.ReplaceAll(log, string(oldRune), string(newRune))
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	//panic("Please implement the WithinLimit() function")
	return utf8.RuneCountInString(log) <= limit
}
