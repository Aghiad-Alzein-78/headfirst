package testo4

import (
	"strings"
)

//JoinWithCommas join strings
func JoinWithCommas(phrase []string) string {
	if len(phrase) == 1 {
		return phrase[0]
	}
	result := strings.Join(phrase[:len(phrase)-1], ",")
	result += " and "
	result += phrase[len(phrase)-1]
	return result
}
