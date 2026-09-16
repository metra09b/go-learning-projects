package main

import (
	"unicode"
)

func CheckOnlyASCII(s string) bool {
	for _, ch := range s {
		if ch > unicode.MaxASCII {
			return false
		}
	}
	return true
}
