postgres:
	docker run --name pg -e POSTGRES_PASSWORD=123 -p 5432:5432 -d postgres:12-alpine

createdb:
	docker exec -it pg createdb --username=postgres --owner=postgres simple_bank

dropdb:
	docker exec -it pg dropdb --username=postgres simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://postgres:123@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:123@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc