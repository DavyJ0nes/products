#### VARIABLES ####
USERNAME = davyj0nes
APP_NAME = products-api

DOCKER_ADDR ?= 192.168.99.100
LOCAL_PORT ?= 8080
APP_PORT ?= 8080

GO_VERSION ?= 1.10.3
GO_PROJECT_PATH ?= github.com/davyj0nes/products
GO_FILES = $(shell go list ./... | grep -v /vendor/)

RELEASE = 0.0.1
COMMIT = $(shell git rev-parse HEAD | cut -c 1-6)
BUILD_TIME = $(shell date -u '+%Y-%m-%d_%I:%M:%S%p')

BUILD_PREFIX = CGO_ENABLED=0 GOOS=linux
BUILD_FLAGS = -a -tags netgo --installsuffix netgo
LDFLAGS = -ldflags "-s -w -X ${GO_PROJECT_PATH}/api/version.Release=${RELEASE} -X ${GO_PROJECT_PATH}/api/version.Commit=${COMMIT} -X ${GO_PROJECT_PATH}/api/version.BuildTime=${BUILD_TIME}"
DOCKER_GO_BUILD = docker run --rm -v "$(GOPATH)":/go -v "$(CURDIR)":/go/src/app -w /go/src/app golang:${GO_VERSION}
GO_BUILD_STATIC = $(BUILD_PREFIX) go build $(BUILD_FLAGS) $(LDFLAGS)
GO_BUILD_OSX = GOOS=darwin GOARCh=amd64 go build $(LDFLAGS)
GO_BUILD_WIN = GOOS=windows GOARCh=amd64 go build $(LDFLAGS)

#### COMMANDS ####
.PHONY: compile
compile:
	$(call blue, "# Building Golang Binary...")
	@${DOCKER_GO_BUILD} sh -c 'go get && ${GO_BUILD_STATIC} -o ${APP_NAME}'

.PHONY: build
build: compile
	$(call blue, "# Building Docker Image...")
	@docker build --no-cache --label APP_VERSION=${RELEASE} --label BUILT_ON=${BUILD_TIME} --label GIT_HASH=${COMMIT} -t ${USERNAME}/${APP_NAME}:${RELEASE} .
	@docker tag ${USERNAME}/${APP_NAME}:${RELEASE} ${USERNAME}/${APP_NAME}:latest
	@docker tag ${USERNAME}/${APP_NAME}:${RELEASE} ${USERNAME}/${APP_NAME}:v1
	@$(MAKE) clean

.PHONY: publish
publish: build
	$(call blue, "# Publishing Docker Image...")
	@docker push docker.io/${USERNAME}/${APP_NAME}:${RELEASE}
	@docker push docker.io/${USERNAME}/${APP_NAME}:v1

.PHONY: run
run:
	$(call blue, "# Running App...")
	@docker run -it --rm -v "$(GOPATH)":/go -v "$(CURDIR)":/go/src/app -p ${LOCAL_PORT}:${APP_PORT} -w /go/src/app golang:${GO_VERSION} go run main.go

.PHONY: run-docker
run-docker: build
	$(call blue, "# Running Docker Image Locally...")
	@docker run -it --rm --name ${APP_NAME} -p ${LOCAL_PORT}:${APP_PORT} ${USERNAME}/${APP_NAME}:${RELEASE} 

.PHONY: deploy
deploy: build
	$(call blue, "# Deploying to Kubernetes...")
	@kubectl create -f kubernetes/service.yml
	@kubectl create -f kubernetes/deployment.yml

.PHONY: test
test:
	$(call blue, "# Testing Golang Code...")
	@docker run --rm -it -v "$(GOPATH):/go" -v "$(CURDIR)":/go/src/app -w /go/src/app golang:${GO_VERSION} sh -c 'go test -v -race ${GO_FILES}' 

.PHONY: transaction-test
transaction-test:
	$(call blue, "# Creating a new Transaction...")
	curl -XPOST -d '{"location": "United Kingdom","product_skus": ["CM01-W","Co01-B","GT01-G"]}' ${DOCKER_ADDR}:${LOCAL_PORT}/api/v1/transaction

.PHONY: clean
clean: 
	@rm -f ${APP_NAME} 

#### FUNCTIONS ####
define blue
	@tput setaf 4
	@echo $1
	@tput sgr0
endef
