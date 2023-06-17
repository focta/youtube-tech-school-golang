postgres:
	docker-compose up -d

createdb:
	docker exec -it youtube-tech-school-golang_dev-postgres_1 createdb --username=yout --owner=yout simple_bank

drodb:
	docker exec -it youtube-tech-school-golang-dev-postgres-1 dropdb simple_bank

.PHONY: createdb