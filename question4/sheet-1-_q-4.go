package main

import (
	"fmt"
	"unicode"
)

// IsIsogram checks if this is isogram.
// Resource: https://github.com/Connor-Cahill/Go-Practice-Problems/blob/master/go/isogram/isogram.go
func IsIsogram(word string) bool {
	dict := make(map[rune]int)

	for _, c := range word {
		_, ok := dict[unicode.ToUpper(c)]

		if ok {
			return false
		}

		if unicode.IsLetter(c) {
			dict[unicode.ToUpper(c)] = 1
		}
	}
	return true
}

func main() {
	// IsIsogram
	fmt.Println(IsIsogram("word")) // false
}
