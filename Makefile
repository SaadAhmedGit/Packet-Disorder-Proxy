.PHONY: build

default:
	@go run .

build:
	@go build -o ./build/main ./main.go ./packet_heap.go
