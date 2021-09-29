build:
	go build cmd/app/main.go

run:
	go run cmd/main.go

swag:
	swag init -g cmd/app/main.go
