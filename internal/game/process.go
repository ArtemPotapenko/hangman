package game

import (
	"errors"
	"strings"
)

const MaxMistakeCount = 5

type Process struct {
	word         *WordInfo
	usingLetters map[rune]struct{}
	mistakeCount int
}

func NewProcess(word string) Process {
	gameWord := newWordInfo(word)
	return Process{word: gameWord, usingLetters: make(map[rune]struct{}), mistakeCount: 0}
}

func (process *Process) GuessLetter(letter rune) (bool, error) {
	if _, ok := process.usingLetters[letter]; ok {
		return false, errors.New("letter already used")
	} else {
		process.usingLetters[letter] = struct{}{}
		guessed := process.word.guessLetter(letter)
		if guessed {
			return true, nil
		} else {
			process.mistakeCount++
			return false, nil
		}
	}
}

func (process *Process) IsGuessed() bool {
	return process.word.guessLittersCount == process.word.wordLength
}

func (process *Process) GetGuessWord() string {
	stringBuilder := strings.Builder{}
	for i, guess := range process.word.guessLitters {
		if guess {
			stringBuilder.WriteString(string(process.word.hiddenWord[i]))
		} else {
			stringBuilder.WriteString("_")
		}
	}
	return stringBuilder.String()
}

func (process *Process) IsLost() bool {
	return process.mistakeCount == MaxMistakeCount
}

func (process *Process) GetHiddenWord() string {
	return string(process.word.hiddenWord)
}

func (process *Process) GetMistakeCount() int {
	return process.mistakeCount
}
