## Init .env file
init-env:
	cp .env.example .env
	
install-tools:
	go install golang.org/x/tools/cmd/stringer@v0.1.0
	go install golang.org/x/tools/cmd/goimports@v0.1.0
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.8.3
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.42.1
	go install github.com/cosmtrek/air@v1.27.3
	
deps:
	go mod download
	go mod tidy

gen:
	go generate ./...

gen-oapi:
	mkdir -p ./app/gen/api
	oapi-codegen -generate "types" -package api ./openapi/v1/openapi.yml > ./app/gen/api/vtm.types.go
	oapi-codegen -generate "server" -package api ./openapi/v1/openapi.yml > ./app/gen/api/vtm.server.go

lint:
	golangci-lint run

test: gen
	ENVCODE=unit go test -v -race -coverprofile=cover.out $(shell go list ./... | grep -vE "(test|gen)/")
	@go tool cover -func=cover.out | grep "total:" | tr '\t' ' '
	go tool cover -html=cover.out -o cover.html
    
up:
	air -c .air.toml