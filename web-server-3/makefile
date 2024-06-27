include .env
export $(shell sed 's/=.*//' .env)

up-server:
	docker-compose up --build

up-db:
	docker-compose -f docker-compose.db.yml up --build