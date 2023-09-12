POSTGRES_URL=postgres://mess:mess@localhost:5432/mess?sslmode=disable

migrate:
	migrate -database $(POSTGRES_URL) -path db/migrations up

up:
	docker-compose up

bg:
	docker-compose up -d

down:
	docker-compose down

sleep:
	sleep 3

setup: bg sleep migrate
