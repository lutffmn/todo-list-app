build:
	@go build -o bin/todo

run: build
	@./bin/todo