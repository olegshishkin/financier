package financier

// Mocks generation
//go:generate mockery --name=.*Storage --recursive --keeptree --with-expecter

// OpenAPI model generation
//go:generate oapi-codegen --package=api -generate=types,spec,gin -o=./api/openapi.go ./api/v1/openapi.yaml

// Dependency injection
//go:generate wire ./...
