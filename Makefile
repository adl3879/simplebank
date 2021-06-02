createdb:
	createdb --username=postgres --owner=postgres simple_bank

dropdb:
	dropdb --username=postgres --owner=postgres simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://postgres:adeleye@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:adeleye@localhost:5432/simple_bank?sslmode=disable" -verbose down

test:
	go test -v -cover ./...