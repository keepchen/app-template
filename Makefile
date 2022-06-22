# workdir info
PACKAGE=app-template
PREFIX=$(shell pwd)
CMD_PACKAGE=${PACKAGE}
OUTPUT_DIR=${PREFIX}/bin
OUTPUT_FILE=${OUTPUT_DIR}/app-template
COMMIT_ID=$(shell git rev-parse --short HEAD)
VERSION=$(shell git describe --tags || echo "v0.0.0")
VERSION_IMPORT_PATH=${PACKAGE}/cmd
BUILD_TIME=$(shell date '+%Y-%m-%dT%H:%M:%S%Z')
VCS_BRANCH=$(shell git symbolic-ref --short -q HEAD)

# build args
BUILD_ARGS = \
    -ldflags "-X $(VERSION_IMPORT_PATH).appName=$(PACKAGE) \
    -X $(VERSION_IMPORT_PATH).version=$(VERSION) \
    -X $(VERSION_IMPORT_PATH).revision=$(COMMIT_ID) \
    -X $(VERSION_IMPORT_PATH).branch=$(VCS_BRANCH) \
    -X $(VERSION_IMPORT_PATH).buildDate=$(BUILD_TIME)"
EXTRA_BUILD_ARGS=

# which cli tools
GOLINT=$(shell which golangci-lint || echo '')
SWAG=$(shell which swag || echo '')

export GOCACHE=
export GOPROXY=https://goproxy.io,direct
export GOSUMDB=

default: lint test build

lint:
	@echo "+ $@"
	@$(if $(GOLINT), , \
		$(error Please install golint: "https://golangci-lint.run/usage/install/#linux-and-windows"))
	golangci-lint run --deadline=10m -E gofmt  -E errcheck ./...

test:
	@echo "+ test"
	go test -cover $(EXTRA_BUILD_ARGS) ./...

.PHONY:build
build:
	@echo "+ build"
	#go build -tags prometheus $(BUILD_ARGS) $(EXTRA_BUILD_ARGS) -o ${OUTPUT_FILE} $(CMD_PACKAGE)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(BUILD_ARGS) -o /output/app-template

dist: build
	@echo "+ $@"
	mkdir -p dist/
	@tar -cvf dist/app-template-${VERSION}.tar README.md \
         		bin/app-template \
         		config/config.yaml

clean:
	@echo "+ $@"
	@rm -r "${OUTPUT_DIR}"

gen-rsa-key:
	openssl genrsa -out $(PREFIX)/static/certifications/rsa_private_key.pem 2048 && \
	openssl rsa -in $(PREFIX)/static/certifications/rsa_private_key.pem \
		-pubout -out ./static/certifications/rsa_public_key.pem

# swag version >= 1.6.7
# go get -u github.com/swaggo/swag/cmd/swag@v1.6.7
gen-swag-httpserver:
	@echo "+ $@"
	@$(if $(SWAG), , \
		$(error Please install swag cli, using go: "go get -u github.com/swaggo/swag/cmd/swag@v1.6.7"))
	swag init --dir pkg/app/httpserver \
 		--output pkg/app/httpserver/docs \
 		--parseDependency --parseInternal \
 		--generalInfo httpserver.go
