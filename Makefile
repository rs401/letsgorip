
help: ## Show this help
	@printf "***\nUsage: Make {target}\nAvailable targets:\n\n"
	@egrep '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

build: build-api build-auth build-forums ## build the api and services

build-api: ## build the Auth API
	@go build -o api/api api/main.go

build-auth: ## build the Auth service
	@go build -o auth/authsvc auth/main.go

build-forums: ## build the Forum service
	@go build -o forums/forumsvc forums/main.go

build-docker: build-api-docker build-auth-docker build-forums-docker ## Build both docker images

build-api-docker: ## build the API docker image
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s' -o docker/api/api api/main.go
	@docker build -t rs401/letsgoripapi:latest docker/api
	@rm docker/api/api

build-auth-docker: ## build the Auth service docker image
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s' -o docker/auth/authsvc auth/main.go
	@docker build -t rs401/letsgoripauthsvc:latest docker/auth
	@rm docker/auth/authsvc

build-forums-docker: ## build the Forum service docker image
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s' -o docker/forums/forumsvc forums/main.go
	@docker build -t rs401/letsgoripforumsvc:latest docker/forums
	@rm docker/forums/forumsvc

kube: ## Run kubectl apply on kubernetes config directory
	@kubectl apply -f k8s/

kube-down: ## Run kubectl delete on kubernetes config directory
	@kubectl delete -f k8s/

proto: ## Run protoc compiler
	@protoc -I=./messages --go_out=./pb --go_opt=paths=source_relative --go-grpc_out=./pb --go-grpc_opt=paths=source_relative ./messages/*.proto

run-docker: ## Run docker commands to start docker containers
	@docker run -d --rm --name lgrauthsvc --net test -p 9001:9001 rs401/letsgoripauthsvc
	@docker run -d --rm --name lgrforumsvc --net test -p 9002:9002 rs401/letsgoripforumsvc
	@docker run -d --rm --name lgrapi --net test -p 9000:9000 rs401/letsgoripapi

stop-docker: ## Stop docker containers running from 'run-docker' target
	@docker stop lgrauthsvc lgrforumsvc lgrapi

test: ## Run all tests
	@go test -v -cover ./...