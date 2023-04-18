include .env

.DEFAULT_GOAL := rebuild-api

docs_init:
	go mod download
	swag init -g cmd/api/main.go

# Updating swagger docs and than running docker-compose
start:
	first_init
	docker-compose up -d --build

update:
	git pull
	docs_init
	docker-compose up -d --build api

restart:
	docker-compose restart

down:
	docker-compose down

race:
	go build -race cmd/api/main.go

migrate:
	echo migrations

test:
	echo test