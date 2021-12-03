# project param
GO_PACKAGE := github.com/feifeifeimoon/mrunc
BIN_DIR := bin
OS := $(shell go env GOOS)
ARCH := $(shell go env GOARCH)
OUTPUT = ${BIN_DIR}/${OS}-${ARCH}/
BUILD_DATE := $(shell date -u +'%Y-%m-%dT%H:%M:%SZ')

# git param
GIT_VERSION := $(shell git rev-parse --abbrev-ref HEAD)
GIT_COMMIT := $(shell git rev-parse HEAD)

# go param
GO_BUILD_PACKAGE = ${GO_PACKAGE}/...
CGO_ENABLED = 0

GO_BUILD_ARGS = -gcflags "all=-N -l" -ldflags "\
-X ${GO_PACKAGE}/pkg/version.gitVersion=${GIT_VERSION} \
-X ${GO_PACKAGE}/pkg/version.gitCommit=${GIT_COMMIT} \
-X ${GO_PACKAGE}/pkg/version.buildDate=${BUILD_DATE}"


.PHONY: all
all: build

.PHONY: build
build:
	mkdir -p $(OUTPUT)
	GOOS=${OS} GOARCH=${ARCH} CGO_ENABLED=${CGO_ENABLED} \
	go build  ${GO_BUILD_ARGS} -o ${OUTPUT} ${GO_BUILD_PACKAGE}
	chmod +x ${OUTPUT}*


clean:
	rm -rf ${BIN_DIR}