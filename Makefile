postgres:
	docker run --name postgres13 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=postgres -d postgres:13-alpine

createdb:
	docker exec -it postgres13 createdb --username=root --owner=root stockmarket-simulator

dropdb:
	docker exec -it postgres13 dropdb stockmarket-simulator

migrateup:
	 .migrate.linux-amd64 -path pkg/db/migration -database "postgresql://root:postgres@localhost:5432/stockmarket-simulator?sslmode=disable" -verbose up

migratedown:
	 .migrate.linux-amd64 -path pkg/db/migration -database "postgresql://root:postgres@localhost:5432/stockmarket-simulator?sslmode=disable" -verbose down
	
sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc
