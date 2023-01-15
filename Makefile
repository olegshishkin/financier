MAIN_FILE_DIR	:= ./cmd/financier
BINARY_NAME 	:= $(shell basename `pwd`)
OUT_DIR 		:= ./.out
REPORT_DIR 		:= report
BINARY_DIR 		:= bin
GREEN_COLOR  	:= $(shell tput -Txterm setaf 2)
RESET_COLOR  	:= $(shell tput -Txterm sgr0)
.DEFAULT_GOAL 	:= build

clean:
	@echo ''
	@echo '$(GREEN_COLOR)Step: clean$(RESET_COLOR)'
	go clean
	rm -rf $(OUT_DIR)
.PHONY: clean

fmt: clean
	@echo ''
	@echo '$(GREEN_COLOR)Step: fmt$(RESET_COLOR)'
	go fmt ./...
.PHONY:fmt

dep: fmt
	@echo ''
	@echo '$(GREEN_COLOR)Step: dep$(RESET_COLOR)'
	go get -d -x ./...
.PHONY:dep

lint: dep
	@echo ''
	@echo '$(GREEN_COLOR)Step: lint$(RESET_COLOR)'
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run ./... -c ./.lint/.golangci.yml
.PHONY:lint

vet: lint
	@echo ''
	@echo '$(GREEN_COLOR)Step: vet$(RESET_COLOR)'
	go vet -c=10 ./...
.PHONY:vet

test: vet
	@echo ''
	@echo '$(GREEN_COLOR)Step: test$(RESET_COLOR)'
	go test ./... -v -race
.PHONY:test

cover: test
	@echo ''
	@echo '$(GREEN_COLOR)Step: cover$(RESET_COLOR)'
	mkdir -p $(OUT_DIR)/$(REPORT_DIR)
	go test ./... -cover -covermode=count -coverprofile=$(OUT_DIR)/$(REPORT_DIR)/profile.out
	go tool cover -html=$(OUT_DIR)/$(REPORT_DIR)/profile.out -o $(OUT_DIR)/$(REPORT_DIR)/coverage.html
.PHONY:cover

build: cover
	@echo ''
	@echo '$(GREEN_COLOR)Step: build$(RESET_COLOR)'
	mkdir -p $(OUT_DIR)/$(BINARY_DIR)
	go build -o $(OUT_DIR)/$(BINARY_DIR)/$(BINARY_NAME) $(MAIN_FILE_DIR)
	ls -alh $(OUT_DIR)/$(BINARY_DIR)/$(BINARY_NAME)
.PHONY:build
