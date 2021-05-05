gPHONY: api-build-local
api-build-local:
	docker build -t dev-api -f build/api/Dockerfile .

gPHONY: api-build-prd
api-build-prd:
	docker build -t prd-api -f build/api/Dockerfile-prd .

gPHONY: compose
compose:
	docker-compose up -d

gPHONY: build-dev
build-dev:
	docker-compose -f docker-compose.yml -f docker-compose.dev.yml -f db.yml build

gPHONY: compose-dev
compose-dev:
	docker-compose -f docker-compose.yml -f docker-compose.dev.yml -f db.yml up -d

gPHONY: migrate-up
migrate-up:
	docker exec -it todo-api sql-migrate up -config=./migrations/dbconfig.yml

gPHONY: migrate-down
migrate-down:
	docker exec -it todo-api sql-migrate down -config=./migrations/dbconfig.yml
