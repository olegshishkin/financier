#!/bin/sh
echo "Install mockery"
go install github.com/vektra/mockery/v3@latest

echo "Install oapi-codegen"
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

echo "Install wire"
go install github.com/google/wire/cmd/wire@latest

echo "Install golangci-lint"
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest