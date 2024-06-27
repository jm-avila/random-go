include .env
export $(shell sed 's/=.*//' .env)

up-db:
	docker-compose -f docker-compose.db.yml up --build