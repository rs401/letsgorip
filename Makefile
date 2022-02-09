
help: ## Show this help
	@printf "***\nUsage: Make {target}\nAvailable targets:\n\n"
	@egrep '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'


build-api: ## build the Auth API
	@go build -o authapi/authapi authapi/main.go

build-auth: ## build the Auth service
	@go build -o auth/authsvc auth/main.go

build-authapi-docker: ## build the Auth API docker image
	@CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -a -installsuffix cgo -ldflags '-s' -o docker/authapi/authapi authapi/main.go
	@docker build -t rs401/letsgoripauthapi:latest docker/authapi

build-auth-docker: ## build the Auth service docker image
	@CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -a -installsuffix cgo -ldflags '-s' -o docker/auth/authsvc auth/main.go
	@docker build -t rs401/letsgoripauthsvc:latest docker/auth

test: ## Run all tests
	@go test -v -cover ./...

kube: ## Run kubectl apply on kubernetes config directory
	@kubectl apply -f k8s/

proto: ## Run protoc compiler
	@protoc -I=./messages --go_out=./pb --go_opt=paths=source_relative --go-grpc_out=./pb --go-grpc_opt=paths=source_relative ./messages/*.proto