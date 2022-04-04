SERVICE=app

.PHONY: run

run:
	sudo docker-compose -f docker-compose.local.yml up


.PHONY: build

build:
	sudo docker-compose -f docker-compose.local.yml build


.PHONY: down

build:
	sudo docker-compose -f docker-compose.local.yml down --volumes --remove-orphans


.PHONY: count

count:
	find . -name tests -prune -o -type f -name '*.go' | xargs wc -l


.DEFAULT_GOAL := run