PROJECT_NAME := "AccountShareC2C"
PKG := "github.com/li-ming-lei/AccountShareC2C"
PKG_LIST := $(shell go mod tidy && go list ${PKG}/...)
BUS_PKG_LIST := $(shell go mod tidy && go list ${PKG}/... | grep -v stub_test)
GO_FILES := $(shell find . -name '*.go' | grep -v _test.go)

.DEFAULT_GOAL := default
.PHONY: all

all: fmt lint vet test race build

dep: ## Get dependencies
	@echo "go dep..."
	@go mod tidy

fmt: dep ## Format code
	@echo "go fmt..."
	@go fmt $(PKG_LIST)

lint: dep ## Lint check
	@echo "golangci-lint run ..."
	@if command -v golangci-lint >/dev/null 2>&1;\
	then echo 'exist golangci-lint';\
	else (export URL=http://172.24.21.29/static;curl -fsSL $$URL/install_lint.sh | sh);\
	fi
	@golangci-lint run

vet: dep ## Vet check
	@echo "go vet..."
	@go vet -all $(PKG_LIST)

test: dep ## Run unittests
	@echo "go test..."
	@go test -gcflags=all=-l -short -v -count=1 ${BUS_PKG_LIST}

race: dep ## Run data race detector
	@echo "go test race..."
	@go test -gcflags=all=-l -race -short -v -count=1 ${BUS_PKG_LIST}

#msan: dep ## Run memory sanitizer
	@#go test -gcflags=all=-l -msan -short ${PKG_LIST}

build: dep fmt ## Build frpc project
	@echo "go build..."
	@CGO_ENABLED=1 go build -v -buildmode=default -o bin/${PROJECT_NAME} cmd/main.go
#	@CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -v -gcflags=all="-N -l" -o bin/${PROJECT_NAME} cmd/main.go
	@chmod +x bin/${PROJECT_NAME}

build-dev: dep fmt ## Build frpc project
	@echo "go build..."
	@go build -v -gcflags=all="-N -l" -o bin/${PROJECT_NAME} cmd/main.go
	@chmod +x bin/${PROJECT_NAME}

run: dep fmt ## Run frpc project
	@echo "go run..."
	@go run cmd/main.go --config=conf/conf.toml

start: build ## Start frpc project
	@echo "start service in background mode..."
	@nohup bin/${PROJECT_NAME} --config=conf/conf.toml >/dev/null 2>&1 &

stop: ## Stop frpc project
	@echo "stop service..."
	@ps ux | grep ${PROJECT_NAME} | grep -v grep | awk '{print $$2}' | xargs kill -9

#restart: stop start ## Restart frpc project

stub_test: fmt ## Run stub test
	@echo "go run stub_test..."
	@go test -gcflags=all=-l -v -count=1 ./stub_test/...

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

default: help

