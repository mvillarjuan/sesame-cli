golang_version := 1.18
VERSION ?= v0.0.1
PROJECT ?= sesame-cli

build-osx:
	@env GOOS=darwin GOARCH=amd64 go build -o bin/$(PROJECT)-osx

build-linux:
	@env GOOS=linux GOARCH=amd64 go build -o bin/$(PROJECT)-linux

build-docker:
	@docker build -t $(PROJECT):$(VERSION) -f Dockerfile \
		--build-arg GOLANG_VERSION=$(golang_version) \
		--build-arg PROJECT_NAME=$(PROJECT) \
		.

launch-tests:
	@docker build -t $(PROJECT):$(VERSION)_test -f DockerfileTesting \
		--build-arg GOLANG_VERSION=$(golang_version) \
		--build-arg PROJECT_NAME=$(PROJECT) \
		.

tests:
	go test -v

push:
	docker push $(PROJECT):$(VERSION)

build-all: build-osx build-linux build-docker

install-deps:
	go mod init || true
	go mod tidy
