postgres:
	docker-compose up -d

createdb:
	docker exec -it youtube-tech-school-golang-dev-postgres-1 createdb --username=yout --owner=yout simple_bank

drodb:
	docker exec -it youtube-tech-school-golang-dev-postgres-1 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://yout:youtpass@localhost:15434/simple_bank?sslmode=disable" -verbose up

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: createdb