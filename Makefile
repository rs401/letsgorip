
help: ## Show this help
	@printf "***\nUsage: Make {target}\nAvailable targets:\n\n"
	@egrep '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

build: build-api build-auth build-forums build-places ## build the api and services

build-api: ## build the Auth API
	@go build -o api/api api/main.go

build-auth: ## build the Auth service
	@go build -o auth/authsvc auth/main.go

build-forums: ## build the Forum service
	@go build -o forums/forumsvc forums/main.go

build-places: ## build the Place service
	@go build -o places/placesvc places/main.go

build-docker: build-api-docker build-auth-docker build-forums-docker build-places-docker build-ui-docker ## Build all docker images
build-gke: build-api-gke build-auth-gke build-forums-gke build-places-gke build-ui-gke ## Build all docker images for GKE

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

build-places-docker: ## build the Place service docker image
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s' -o docker/places/placesvc places/main.go
	@docker build -t rs401/letsgoripplacesvc:latest docker/places
	@rm docker/places/placesvc

build-ui-docker: ## build the ui docker image
	@docker build -t rs401/letsgoripui:latest ui/

build-api-gke: ## build the API docker image for gke
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s' -o docker/api/api api/main.go
	@docker build -t us-east1-docker.pkg.dev/letsgorip/lgr-repo/lgrapi docker/api
	@rm docker/api/api

build-auth-gke: ## build the Auth service docker image for gke
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s' -o docker/auth/authsvc auth/main.go
	@docker build -t us-east1-docker.pkg.dev/letsgorip/lgr-repo/lgrauthsvc docker/auth
	@rm docker/auth/authsvc

build-forums-gke: ## build the Forum service docker image for gke
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s' -o docker/forums/forumsvc forums/main.go
	@docker build -t us-east1-docker.pkg.dev/letsgorip/lgr-repo/lgrforumsvc docker/forums
	@rm docker/forums/forumsvc

build-places-gke: ## build the Place service docker image for gke
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s' -o docker/places/placesvc places/main.go
	@docker build -t us-east1-docker.pkg.dev/letsgorip/lgr-repo/lgrplacesvc docker/places
	@rm docker/places/placesvc

build-ui-gke: ## build the ui docker image for gke
	@docker build -t us-east1-docker.pkg.dev/letsgorip/lgr-repo/lgrui ui/


kube: ## Run kubectl apply on kubernetes config directory
	@kubectl apply -f k8s/

kube-down: ## Run kubectl delete on kubernetes config directory
	@kubectl delete -f k8s/

prev-readme: ## Compile readme and preview in browser
	@pandoc -s README.md -o README.html
	@$$BROWSER README.html

proto: ## Run protoc compiler
	@protoc -I=./messages --go_out=./pb --go_opt=paths=source_relative --go-grpc_out=./pb --go-grpc_opt=paths=source_relative ./messages/*.proto

push-gke: ## Push project images to G artifacts registry
	@docker push us-east1-docker.pkg.dev/letsgorip/lgr-repo/lgrplacesvc:latest
	@docker push us-east1-docker.pkg.dev/letsgorip/lgr-repo/lgrforumsvc:latest
	@docker push us-east1-docker.pkg.dev/letsgorip/lgr-repo/lgrauthsvc:latest
	@docker push us-east1-docker.pkg.dev/letsgorip/lgr-repo/lgrapi:latest
	@docker push us-east1-docker.pkg.dev/letsgorip/lgr-repo/lgrui:latest


run-docker: ## Run docker commands to start docker containers
	@docker run -d --rm --name lgrauthsvc --net test -p 9001:9001 rs401/letsgoripauthsvc
	@docker run -d --rm --name lgrforumsvc --net test -p 9002:9002 rs401/letsgoripforumsvc
	@docker run -d --rm --name lgrplacesvc --net test -p 9003:9003 rs401/letsgoripplacesvc
	@docker run -d --rm --name lgrapi --net test -p 9000:9000 rs401/letsgoripapi

stop-docker: ## Stop docker containers running from 'run-docker' target
	@docker stop lgrauthsvc lgrforumsvc lgrplacesvc lgrapi

test: ## Run all tests
	@go test -v -cover ./...