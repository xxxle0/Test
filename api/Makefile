run:
	@echo "======= Backend's Running ======="
	go run main.go

build:
	@echo "======= Build Docker Image ======="
	docker build -t duybkit13/guard-test -f ./deployments/Dockerfile .

swagger-migrate:
	@echo "======= Swagger Migrate ======="
	swag init

test:
	go test ./... -v
