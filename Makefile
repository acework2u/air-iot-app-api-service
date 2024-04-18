include .env

dev:
	docker-compose up
stop:
	docker-compose down

build:
	docker buildx build --platform=linux/amd64 -t go-ci-cd-prod . --target production -f Dockerfile.production --no-cache
	# docker buildx build --platform=linux/amd64 -t go-ci-cd-prod . --target production -f Dockerfile.production --no-cache

start:
	docker run -p 8000:8000 --name go-ci-cd-prod go-ci-cd-prod 

ec2:
	#GOOS=linux GOARCH=amd64 go build -o ./app/bin/server .
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./app/bin/server .

test-winds:
	GOOS=windows GOARCH=386 go build -o ./app/windows/server.exe .

test-linux:
	GOOS=linux GOARCH=amd64 go build -o ./app/linux/server ./cmd/main.go

api-docs:
	swag init

dev-app-api:
	@go run cmd/app/main.go
