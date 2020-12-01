maria:
	docker run \
		-d \
		-e MYSQL_RANDOM_ROOT_PASSWORD=true \
		-e MYSQL_DATABASE=godo \
		-e MYSQL_USER=godo \
		-e MYSQL_PASSWORD=secret \
		-p 5500:3306 \
		--name godoDatabase \
		mariadb:latest

redis:
	docker run \
		-d \
		-p 6600:6379 \
		--name godoStorage \
		redis:6.0.9-alpine

infra:
	make maria && make redis

migrateon:
	migrate create -ext sql -dir database/migration -seq  init_schema

migrateup:
	migrate -path database/migration -database "mysql://godo:secret@tcp(127.0.0.1:5500)/godo" -verbose up

migratedown:
	migrate -path database/migration -database "mysql://godo:secret@tcp(127.0.0.1:5500)/godo" -verbose down

run:
	go run main.go

.PHONY: maria redis infra migrateon migrateup migratedown run