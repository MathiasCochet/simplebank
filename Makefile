postgres:
	docker run --name postgres14 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it postgres14 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres14 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc-version:
	docker run --rm -v "$(CURDIR):/src" -w //src kjconroy/sqlc version
sqlc-init:
	docker run --rm -v "$(CURDIR):/src" -w //src kjconroy/sqlc init
sqlc-compile:
	docker run --rm -v "$(CURDIR):/src" -w //src kjconroy/sqlc compile
sqlc-generate:
	docker run --rm -v "$(CURDIR):/src" -w //src kjconroy/sqlc generate

test:
	go test -v -cover ./...

server: 
	go run main.go
	
mock: 
	mockgen -package mockdb -destination db/mock/store.go hsehld.dev/m/v2/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc-version sqlc-init sqlc-compile sqlc-generate test server mock