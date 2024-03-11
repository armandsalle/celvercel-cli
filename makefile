.PHONY: run build

run:
	go run main.go

build:
	go build -o build/celvercel

dev:
	go build -o build/celvercel && ./build/celvercel