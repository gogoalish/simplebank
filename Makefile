LINTVER=v1.49.0
LINTBIN=${BINDIR}/lint_${GOVER}_${LINTVER}
BINDIR=${CURDIR}/bin
REMOTEPSQL=postgres://root:AIUC8OKmiP4k2RGNgshL9pVtmF2zym7T@dpg-cp331d8l5elc73ag1m8g-a.oregon-postgres.render.com/simple_bank_fhar

network:
	docker network create bank-network

postgres:
	docker run --name postgres15 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine

start:
	docker start postgres15

stop:
	docker stop postgres15

remove:
	docker remove postgres15

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres15 dropdb simple_bank

migrateup-local:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrateup1-local:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown-local:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratedown1-local:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

migrateup:
	migrate -path db/migration -database "${REMOTEPSQL}" -verbose up

migrateup1:
	migrate -path db/migration -database "${REMOTEPSQL}" -verbose up 1

migratedown:
	migrate -path db/migration -database "${REMOTEPSQL}" -verbose down

migratedown1:
	migrate -path db/migration -database "${REMOTEPSQL}" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/gogoalish/simplebank/db/sqlc Store

pull:
	docker pull postgres:15-alpine

	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest 

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock migratedown1 migrateup1