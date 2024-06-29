build:
	go build -o . cmd/main.go

test:
	go test -v ./...

up-server: build
	./main

up-db:
	docker-compose -f docker-compose.db.yml up --build