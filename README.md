# Golang Clean Architecture Template

## Description

This is golang clean architecture template forked from https://github.com/khannedy/golang-clean-architecture. Here I try to improve a features and functionalities.

## Architecture

![Clean Architecture](architecture.png)

1. External system perform request (HTTP, gRPC, Messaging, etc)
2. The Delivery creates various Model from request data
3. The Delivery calls Use Case, and execute it using Model data
4. The Use Case create Entity data for the business logic
5. The Use Case calls Repository, and execute it using Entity data
6. The Repository use Entity data to perform database operation
7. The Repository perform database operation to the database
8. The Use Case create various Model for Gateway or from Entity data
9. The Use Case calls Gateway, and execute it using Model data
10. The Gateway using Model data to construct request to external system
11. The Gateway perform request to external system (HTTP, gRPC, Messaging, etc)

## Tech Stack

- Golang : https://github.com/golang/go
- MySQL (Database) : https://github.com/mysql/mysql-server
- Apache Kafka : https://github.com/apache/kafka

## Framework & Library

- GoFiber (HTTP Framework) : https://github.com/gofiber/fiber
- GORM (ORM) : https://github.com/go-gorm/gorm
- Viper (Configuration) : https://github.com/spf13/viper
- Golang Migrate (Database Migration) : https://github.com/golang-migrate/migrate
- Go Playground Validator (Validation) : https://github.com/go-playground/validator
- Logrus (Logger) : https://github.com/sirupsen/logrus
- Confluent Kafka Golang : https://github.com/confluentinc/confluent-kafka-go

## Configuration

All configuration is in `config.json` file.

## API Spec

All API Spec is in `api` folder.

## Install command-line tool (only if you haven't install it yet)

- Install `migrate` globally: `go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`
- Docker and Kafka configuration: https://github.com/saadabdurrazaq/docker-setup-and-instalation

## Database Migration

All database migration is in `db/migrations` folder.

### Create Migration

```sh
migrate create -ext sql -dir db/migrations create_table_xxx

```

### Run Migration

```sh
migrate -database "mysql://root:root@tcp(localhost:3306)/golang_clean_architecture?charset=utf8mb4&parseTime=True&loc=Local" -path db/migrations up

```

## Run Application

### Run Docker

```bash
- Open docker desktop
- cd <your go directory>
- then execute command below

docker-compose down
docker-compose up -d

```

### Run unit test

```bash
go test -v ./test/

```

### Run web server

```bash
go run cmd/web/main.go

```

### Run worker (very much like a cron job with a few differences)

```bash
go run cmd/worker/main.go

```

### Open phpmyadmin

- Go to http://localhost:8080/
- username: root
- password: root

### Token Usage in Postman

- Go to the Headers tab.
- Add this key-value pair:<br>
Key: Authorization<br>
Value: your_token (you can get the token after you successfully login)