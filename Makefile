build-up:
	docker compose -f ./docker-compose.local.yml up --build

start:
	air

install:
	docker exec -it cws-back-cws-back-1 yarn install

bash:
	docker exec -it cws-back-cws-back-1 bash

migration-create:
	dmate new $(name)

migration-migrate:
	dbmate migrate

.PHONY: test

test:
	docker exec -it cws-back-cws-back-1 yarn test


build:
	docker compose up --build -d --remove-orphans
up:
	docker compose up -d
down:
	docker compose down
show_logs:
	docker compose logs