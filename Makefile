.PHONY: run build test docker-build docker-run tidy

run:
	go run cmd/api/main.go

build:
	go build -o server cmd/api/main.go

test:
	go test ./...

tidy:
	go mod tidy

docker-build:
	docker build -t product-api .

docker-run:
	docker run -p 8080:8080 product-api
