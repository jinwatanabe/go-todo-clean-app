run:
	DB_USER=root DB_PASSWORD=password DB_HOST=localhost DB_PORT=3306 go run .
test:
	go test -v ./...
gen:
	wire di/wire.go
db:
	docker-compose up -d
stop-db:
	docker-compose down