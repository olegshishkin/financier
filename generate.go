package financier

// OpenAPI
//go:generate oapi-codegen -generate types -o ./api/v1/model.go -package v1 ./api/v1/openapi.yaml
//go:generate oapi-codegen -generate gin -o ./api/v1/handlers.go -package v1 ./api/v1/openapi.yaml
//go:generate oapi-codegen -generate spec -o ./api/v1/swagger.go -package v1 ./api/v1/openapi.yaml

// Mocks
//go:generate mockery --name=.*Storage --recursive --keeptree --with-expecter --config=./mocks/.mockery

// Dependency injection
//go:generate wire ./...
