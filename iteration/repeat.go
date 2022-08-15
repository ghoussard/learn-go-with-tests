package iteration

import "strings"

// Repeat a character n times
func Repeat(character string, times int) (repeated string) {
	repeated = strings.Repeat(character, times)

	return
}
