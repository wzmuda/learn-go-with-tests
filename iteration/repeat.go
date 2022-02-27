package iteration

import (
	"strings"
)

func Repeat(c string, count int) (repeated string) {
	if count < 0 {
		return c
	}
	return strings.Repeat(c, count)
}
