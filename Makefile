#! /usr/bin/make
#
# Makefile for goa v2
#
# Targets:
# - "depend" retrieves the Go packages needed to run the linter and tests
# - "lint" runs the linter and checks the code format using goimports
# - "test" runs the tests
#
# Meta targets:
# - "all" is the default target, it runs all the targets in the order above.
#
MAJOR=2
MINOR=2
BUILD=6

GOOS=$(shell go env GOOS)
GO_FILES=$(shell find . -type f -name '*.go')
DIR=$(shell pwd)

ifeq ($(GOOS),windows)
EXAMPLES_DIR="$(GOPATH)\src\goa.design\examples"
PLUGINS_DIR="$(GOPATH)\src\goa.design\plugins"
GOBIN="$(GOPATH)\bin"
else
EXAMPLES_DIR=$(GOPATH)/src/goa.design/examples
PLUGINS_DIR=$(GOPATH)/src/goa.design/plugins
GOBIN=$(GOPATH)/bin
endif

export GO111MODULE=off

# Only list test and build dependencies
# Standard dependencies are installed via go get
DEPEND=\
	golang.org/x/lint/golint \
	golang.org/x/tools/cmd/goimports \
	github.com/golang/protobuf/protoc-gen-go \
	github.com/golang/protobuf/proto \
	honnef.co/go/tools/cmd/staticcheck \
	github.com/getkin/kin-openapi/openapi3 \
	github.com/google/uuid

all: lint test

travis: depend all

# Install protoc
PROTOC_VERSION=3.14.0
ifeq ($(GOOS),linux)
PROTOC=protoc-$(PROTOC_VERSION)-linux-x86_64
PROTOC_EXEC=$(PROTOC)/bin/protoc
else
	ifeq ($(GOOS),darwin)
PROTOC=protoc-$(PROTOC_VERSION)-osx-x86_64
PROTOC_EXEC=$(PROTOC)/bin/protoc
	else
		ifeq ($(GOOS),windows)
PROTOC=protoc-$(PROTOC_VERSION)-win32
PROTOC_EXEC="$(PROTOC)\bin\protoc.exe"
		endif
	endif
endif
depend:
	@go get -v $(DEPEND)
	@go get github.com/hashicorp/go-getter/cmd/go-getter && \
		go-getter https://github.com/google/protobuf/releases/download/v$(PROTOC_VERSION)/$(PROTOC).zip $(PROTOC) && \
		cp $(PROTOC_EXEC) $(GOBIN) && \
		rm -r $(PROTOC) && \
		echo "`protoc --version`"
	@go get -t -v ./...

lint:
	@if [ "`goimports -l $(GO_FILES) | tee /dev/stderr`" ]; then \
		echo "^ - Repo contains improperly formatted go files" && echo && exit 1; \
	fi
	@if [ "`golint ./... | grep -vf .golint_exclude | tee /dev/stderr`" ]; then \
		echo "^ - Lint errors!" && echo && exit 1; \
	fi
	@if [ "`staticcheck -checks all ./... | grep -v ".pb.go" | grep -v "SA1019" | tee /dev/stderr`" ]; then \
		echo "^ - staticcheck errors!" && echo && exit 1; \
	fi

test:
	go test ./...

release:
	# First make sure all is clean
	git diff-index --quiet HEAD
	cd $(GOPATH)/src/goa.design/examples && \
		git checkout v$(MAJOR) && \
		git pull origin v$(MAJOR) && \
		git diff-index --quiet HEAD
	cd $(GOPATH)/src/goa.design/plugins && \
		git checkout v$(MAJOR) && \
		git pull origin v$(MAJOR) && \
		git diff-index --quiet HEAD
	# Bump version number, commit and push
	sed 's/Major = .*/Major = $(MAJOR)/' pkg/version.go > _tmp && mv _tmp pkg/version.go
	sed 's/Minor = .*/Minor = $(MINOR)/' pkg/version.go > _tmp && mv _tmp pkg/version.go
	sed 's/Build = .*/Build = $(BUILD)/' pkg/version.go > _tmp && mv _tmp pkg/version.go
	sed 's/Current Release: `v.*/Current Release: `v$(MAJOR).$(MINOR).$(BUILD)`/' README.md > _tmp && mv _tmp README.md
	sed 's/Code generated by goa .*/Code generated by goa v$(MAJOR).$(MINOR).$(BUILD), DO NOT EDIT./' README.md > _tmp && mv _tmp README.md
	git add .
	git commit -m "Release v$(MAJOR).$(MINOR).$(BUILD)"
	git tag v$(MAJOR).$(MINOR).$(BUILD)
	cd cmd/goa && go install
	git push origin v$(MAJOR)
	git push origin v$(MAJOR).$(MINOR).$(BUILD)
	# Update examples
	cd $(GOPATH)/src/goa.design/examples && \
		make && \
		git add . && \
		git commit -m "Release v$(MAJOR).$(MINOR).$(BUILD)" && \
		git tag v$(MAJOR).$(MINOR).$(BUILD) && \
		git push origin v$(MAJOR)
		git push origin v$(MAJOR).$(MINOR).$(BUILD)
	# Update plugins
	cd $(GOPATH)/src/goa.design/plugins && \
		make && \
		git add . && \
		git commit -m "Release v$(MAJOR).$(MINOR).$(BUILD)" && \
		git tag v$(MAJOR).$(MINOR).$(BUILD) && \
		git push origin v$(MAJOR) && \
		git push origin v$(MAJOR).$(MINOR).$(BUILD)
	echo DONE RELEASING v$(MAJOR).$(MINOR).$(BUILD)!
