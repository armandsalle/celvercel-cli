.PHONY: run build

run:
	go run main.go

build:
	go build -o build/celvercel

build-windows:
	GOOS=windows GOARCH=amd64 go build -o bin/app-windows.exe

build-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/app-linux

build-mac:
	GOOS=darwin GOARCH=amd64 go build -o bin/app-mac

build-mac-silicon:
	GOOS=darwin GOARCH=arm64 go build -o bin/app-mac-silicon

dev:
	go build -o build/celvercel && ./build/celvercel