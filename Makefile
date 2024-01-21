CURRENT_DIR=$(shell pwd)

APP=$(shell basename ${CURRENT_DIR})
APP_CMD_DIR=${CURRENT_DIR}/cmd

TAG=latest
ENV_TAG=latest

migration-up:
	migrate -path ./migrations/postgres -database 'postgres://postgres:1@0.0.0.0:5432/news_blogs_service?sslmode=disable' up

migration-down:
	migrate -path ./migrations/postgres -database 'postgres://postgres:1@0.0.0.0:5432/news_blogs_service?sslmode=disable' down

migration-up-test:
	migrate -path ./migrations/postgres -database 'postgres://postgres:1@0.0.0.0:5432/test_news_blogs_service?sslmode=disable' up

migration-down-test:
	migrate -path ./migrations/postgres -database 'postgres://postgres:1@0.0.0.0:5432/test_news_blogs_service?sslmode=disable' down

swag-init:
	swag init -g api/api.go -o api/docs

run:
	go run cmd/main.go
