DB_URL=postgresql://root:523127@localhost:5432/booking?sslmode=disable
sqlc:
	sqlc generate

createdb:
	docker run -p 3306:3306 --name mysql --network mysql -e MYSQL_ROOT_PASSWORD=523127 -e MYSQL_DATABASE=booking  -d mysql:8.0
dropdb:
	docker rm -f postgres12
migrateup1:
	migrate -path internal/model/migration -database "$(DB_URL)" -verbose up 1
migrateup2:
	migrate -path internal/model/migration -database "$(DB_URL)" -verbose up 2
migratedown:
	migrate -path internal/model/migration -database "$(DB_URL)" -verbose down
mock:
	mockgen  -package mock_repository  -source internal/model/domain/querier.go Querier  -destination internal/model/mock/mock.go >  internal/model/mock_repository/mock.go
postgres:
	docker run --name postgres12 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=523127 -d postgres:12-alpine
createdb:
	docker exec -it postgres12 createdb --username=root --owner=root booking
.PHONY: sqlc createdb migratedown migrateup1 mock
