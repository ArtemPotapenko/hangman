.PHONY: format vet tidy build run clean

format:
	go fmt ./...

vet:
	go vet ./...

tidy:
	go mod tidy

build: vet tidy
	go build -o hangman ./cmd

run: build
	./hangman $(FILE)

clean:
	rm -f hangman