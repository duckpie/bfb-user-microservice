ENV=local

.PHONY: run
run:
	sudo docker-compose -f docker-compose.$(ENV).yml build
	sudo docker-compose -f docker-compose.$(ENV).yml up


.PHONY: build
build:
	sudo docker-compose -f docker-compose.$(ENV).yml build


.PHONY: down
down:
	sudo docker-compose -f docker-compose.$(ENV).yml down --volumes --remove-orphans


.PHONY: test
test:
	sudo docker-compose -f docker-compose.test.yml build
	sudo docker-compose -f docker-compose.test.yml up \
			--remove-orphans \
			--exit-code-from user_ms_service_test \
			--abort-on-container-exit user_ms_service_test
	sudo docker stop bfb-user-microservice_db_test_1
			

.PHONY: count
count:
	find . -name tests -prune -o -type f -name '*.go' | xargs wc -l


.DEFAULT_GOAL := run
