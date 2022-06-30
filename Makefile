## Init .env file
init-env:
	cp .env.example .env
	
install-tools:
	go install golang.org/x/tools/cmd/stringer@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/cosmtrek/air@latest
	
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

clean:
	go clean
	go clean --modcache
	go clean --cache
