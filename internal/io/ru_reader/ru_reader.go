package ru_reader

import (
	"bufio"
	"fmt"
	"hangman/internal/russian"
	"os"
	"strings"
)

type RuReader struct {
	fd     int
	reader *bufio.Reader
}

func New(fd int) *RuReader {
	return &RuReader{
		fd:     fd,
		reader: bufio.NewReader(os.Stdin),
	}
}

func (r *RuReader) ReadRune() (rune, int, error) {
	fmt.Print("Введите русскую букву: ")
	for {
		line, err := r.reader.ReadString('\n')
		if err != nil {
			return 0, 0, err
		}
		line = strings.TrimSpace(line)
		runes := []rune(line)
		if len(runes) != 1 {
			fmt.Print("Введите русскую букву: ")
			continue
		}
		if !russian.IsRussianLetter(runes[0]) {
			fmt.Print("Введите русскую букву: ")
			continue
		}
		return runes[0], 1, nil
	}
}
