# Gin Framework Boilerplate
This is a boilerplate code which can be used for any project using Gin as the framework (Go language) and PostgreSQL as the database.

## Run the Apps
To run the apps locally, run following command.
```go run cmd/api/main.go```

## Run Test & Get Coverage
To run all tests and get the coverage, run following command.
```go test ./... --cover```

## Initialize Swagger
To update Swagger documentation after any changes occurred in the code, run following command.
```swag init -g cmd/api/server/server.go -o cmd/api/docs```

## Initialize Mock for Specific Interface
To initialize mock of specific interface, run following command.
```mockery --dir path-to-interface-dir --name interface-name --filename file-name --output path-to-output-dir --outpkg package-name```
For example:
```mockery --dir pkg/jwt --name JWTService --filename jwt_test.go --output internal/mocks --outpkg mocks```
