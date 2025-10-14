up:
	docker-compose -f docker-compose.yaml up --build

down:
	docker-compose -f docker-compose.yaml down

logs:
	docker-compose -f docker-compose.yaml logs -f

ps:
	docker-compose -f docker-compose.yaml ps

connect-db:
	docker exec -it Database_container psql -U user -d appdb

migrations:
	migrate create -ext sql -dir ./database/migrations -seq $(filter-out $@, $(MAKECMDGOALS))
%:
	@:

generate-docs:
	swag init

.PHONY: up down logs ps connect-db migrations generate-docs