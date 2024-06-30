build:
	go build -o . cmd/main.go

test:
	go test -v ./...

up-server: build
	./main

up-db:
	docker-compose -f docker-compose.db.yml up --build

migration-create:
	migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migration-up:
	go run cmd/migrate/main.go up

migration-down:
	go run cmd/migrate/main.go down

mysql:
	docker exec -it web-server-4-mysql-db-1 mysql --user=root --password=strong_password
