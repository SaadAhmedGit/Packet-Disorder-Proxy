.PHONY: build

default:
	@go run ./main.go

build:
	@go build -o ./build/main ./main.go
