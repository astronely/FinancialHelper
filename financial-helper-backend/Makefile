postgres:
	docker run --name postgresFH -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=admin -d postgres:latest

createdb:
	docker exec -it postgresFH createdb --username=root --owner=root f_helper

dropdb:
	docker exec -it postgresFH dropdb f_helper

migrateup:
	migrate -path db/migration -database "postgresql://root:admin@localhost:5432/f_helper?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:admin@localhost:5432/f_helper?sslmode=disable" -verbose down

run:
	docker start postgresFH

stop:
	docker stop postgresFH

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 run stop sqlc test serve