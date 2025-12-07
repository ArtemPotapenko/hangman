package main

import (
	"fmt"
	"hangman/internal/dict"
	"hangman/internal/game"
	"hangman/internal/io/drawer"
	"hangman/internal/io/filereader"
	"hangman/internal/io/rureader"
	"io"
	"math/rand"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Используй: go run main.go <filename>")
		os.Exit(1)
	}
	filename := os.Args[1]
	reader, err := filereader.New(filename)
	if err != nil {
		fmt.Println("Ошибка открытия файла " + filename)
		os.Exit(1)
	}
	defer func(reader *filereader.FileReader) {
		err := reader.Close()
		if err != nil {
			panic(err)
		}
	}(reader)

	stringArr, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	stringArr = dict.FilterRussian(stringArr)
	if len(stringArr) == 0 {
		fmt.Println("Словарь пустой либо все строки не подходят для игры")
		os.Exit(1)
	}
	fmt.Printf("Обнаружено %d строк\n", len(stringArr))

	process := game.NewProcess(stringArr[rand.Intn(len(stringArr))])

	fmt.Println(process.GetGuessWord())
	ruReader := rureader.New(int(os.Stdin.Fd()))

	for !process.IsGuessed() && !process.IsLost() {
		letter, _, err := ruReader.ReadRune()

		if err != nil {
			if err == io.EOF {
				fmt.Println("Завершение программы")
				os.Exit(0)
			}
			panic(err)
		}

		guessLetter, err := process.GuessLetter(letter)
		if err != nil {
			fmt.Printf("Буква \"%v\" уже была использована\n", string(letter))
			continue
		}

		if guessLetter {
			fmt.Printf("Буква \"%v\" угадана\n", string(letter))

		} else {
			fmt.Printf("Увы, буква \"%v\" отсутвует\n", string(letter))
		}

		fmt.Println()
		err = drawer.DrawHangman(process.GetMistakeCount())
		if err != nil {
			panic(err)
		}
		fmt.Println()
		fmt.Println()
		fmt.Println(process.GetGuessWord())
		fmt.Println()

	}

	if process.IsGuessed() {
		fmt.Println()
		fmt.Printf("Вы угадали слово: %s!\n", process.GetGuessWord())
	} else {
		fmt.Println()
		fmt.Printf("Вы проиграли. Загаданное слово: %s.\n", process.GetHiddenWord())
	}
}
