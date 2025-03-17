SRC = ./cmd/server

all: up run

run:
	go run $(SRC)/main.go

up:
	docker-compose up -d

down:
	docker-compose down

clean: down
	docker volume rm `docker volume ls -q`