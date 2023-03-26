#!/bin/sh

# OpenAPI
echo "Generate OpenAPI code"
oapi-codegen -generate types -o ./api/v1/model.go -package v1 ./api/v1/openapi.yaml
oapi-codegen -generate gin -o ./api/v1/handlers.go -package v1 ./api/v1/openapi.yaml
oapi-codegen -generate spec -o ./api/v1/swagger.go -package v1 ./api/v1/openapi.yaml

# Mocks
echo "Generate mocks code"
mockery --name=.*Storage --recursive --keeptree --with-expecter --config=./scripts/.mockery

# Dependency injection
echo "Generate dependency injection code"
wire ./...