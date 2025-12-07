package file_reader

import (
	"bufio"
	"errors"
	"io"
	"os"
)

type FileReader struct {
	file    *os.File
	scanner *bufio.Scanner
}

func New(path string) (*FileReader, error) {
	if path == "" {
		return nil, errors.New("file path is empty")
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return &FileReader{
		file:    f,
		scanner: bufio.NewScanner(f),
	}, nil
}

func (fr *FileReader) Close() error {
	if fr.file != nil {
		return fr.file.Close()
	}
	return nil
}

func (fr *FileReader) ReadLine() (string, error) {
	if fr.scanner.Scan() {
		return fr.scanner.Text(), nil
	}
	if err := fr.scanner.Err(); err != nil {
		return "", err
	}
	return "", io.EOF
}

func (fr *FileReader) ReadAll() ([]string, error) {
	var lines []string
	for {
		line, err := fr.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, err
		}

		lines = append(lines, line)
	}
	return lines, nil
}
