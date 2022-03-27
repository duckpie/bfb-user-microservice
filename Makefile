SERVICE=app

.PHONY: run

run:
	go build -o $(SERVICE) main.go
	clear		
	./$(SERVICE) run


.PHONY: count

count:
	find . -name tests -prune -o -type f -name '*.go' | xargs wc -l


.DEFAULT_GOAL := run