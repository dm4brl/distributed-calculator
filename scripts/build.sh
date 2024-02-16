#!/bin/bash

git clone https://github.com/dm4brl/distributed-calculator.git

cd distributed-calculator

POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=distributed_calculator
POSTGRES_HOST=postgres
POSTGRES_PORT=5432

REDIS_HOST=redis
REDIS_PORT=6379

RABBITMQ_HOST=rabbitmq
RABBITMQ_PORT=5672

docker-compose up -d

go get -u github.com/golang/dep/cmd/dep
dep ensure

go build -o server ./cmd/server/main.go
go build -o agent ./cmd/agent/main.go


./server

./agent


