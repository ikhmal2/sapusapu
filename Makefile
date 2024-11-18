all: build run


run:
	./main

build:
	go build -o main ./cmd/web/*

clean:
	rm main
