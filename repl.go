package main

import (
	"strings"
)

func cleanInput(text string) []string {
	lowercase := strings.ToLower(text)
	words  := strings.Fields(lowercase)
	return words
}

// lol, i reimplemented this like a caveman..
//func cleanInput(text string) []string {
//	var result []string
//	new_word := true
//	for _, char := range text {
//		if char == ' ' {
//			new_word = true
//			continue
//		}
//		if new_word {
//			result = append(result, "")
//		}
//		result[len(result) - 1] += string(char)
//		new_word = false
//	}
//	return result
//}

