package dict

import (
	"hangman/internal/russian"
	"strings"
)

func FilterRussian(words []string) []string {
	var filteredWords []string
	for _, word := range words {
		runes := []rune(strings.TrimSpace(word))
		skipWord := false
		if len(runes) < 4 || len(runes) > 8 {
			skipWord = true
		}

		for _, curRune := range runes {
			if curRune == ' ' || curRune == '\t' {
				skipWord = true
			}
			if !russian.IsRussianLetter(curRune) {
				skipWord = true
			}
		}
		if !skipWord {
			filteredWords = append(filteredWords, strings.ToLower(string(runes)))
		}
	}
	return filteredWords
}
