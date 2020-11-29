postgres:
	docker run \
		-d \
		-e POSTGRES_USER=godo \
		-e POSTGRES_PASSWORD=secret \
		-e POSTGRES_DB=godo \
		-p 5500:5432  \
		--name godoDB \
		postgres:13.1-alpine

redis:
	docker run \
		-d \
		-p 6600:6379 \
		--name godoSTG \
		redis:6.0.9-alpine

infra:
	make postgres && make redis

migrateon:
	migrate create -ext sql -dir database/migration -seq  init_schema

migrateup:
	migrate -path database/migration -database "postgres://godo:secret@127.0.0.1:5500/godo?sslmode=disable" -verbose up

migratedown:
	migrate -path database/migration -database "postgres://godo:secret@127.0.0.1:5500/godo?sslmode=disable" -verbose down

run:
	go run main.go

.PHONY: postgres migrateon migrateup migratedown infra run