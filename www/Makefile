## for the module itself
MODULE_PATH = github.com/golangee/forms-example/www

## which linter version to use?
GOLANGCI_LINT_VERSION = v1.32.2

## for ldflags replacement
BUILD_FILE_PATH = ${MODULE_PATH}

## which linter version to use?
GOLANGCI_LINT_VERSION = v1.24.0

buildVersion := "v0.1.0"
currentDate := $(shell date +"%Y.%m.%d-%H:%M:%S")
currentCommit := $(shell git rev-parse HEAD)
hostName := $(shell hostname)

LDFLAGS = -X $(MODULE_PATH)/internal/build.Time=${currentDate} -X $(MODULE_PATH)/internal/build.Commit=${currentCommit} -X $(MODULE_PATH)/internal/build.Host=${hostName} -X $(MODULE_PATH)/internal/build.Version=${buildVersion}



TMP_DIR = $(TMPDIR)/$(MODULE_PATH)
ifndef BUILD_DIR
    $(info BUILD_DIR is undefined, using $(TMP_DIR))
	BUILD_DIR = $(TMP_DIR)
endif

GO = go
GOLANGCI_LINT = ${TMP_DIR}/golangci-lint
GOLINT = ${TMP_DIR}/golint
GOROOT = $(shell ${GO} env GOROOT)

all: generate lint test build compress ## Runs lint, test and build and compression for release

clean: ## Removes any temporary and output files
	rm -rf ${BUILD_DIR}

lint: ## Executes all linters
	${GOLANGCI_LINT} run --enable-all --exclude-use-default=false

test: ## Executes the tests
	${GO} test -race ./...

.PHONY: build generate setup

build: generate ## Performs a build and puts everything into the build directory
	GOOS=js GOARCH=wasm ${GO} build -ldflags "${LDFLAGS}" -o ${BUILD_DIR}/app.wasm ${MODULE_PATH}/cmd/wasm


	cp "${GOROOT}/misc/wasm/wasm_exec.js" ${BUILD_DIR}
	cp -r static/. ${BUILD_DIR}


compress: build ## Applies gzip and brotli compression to build files
	$(shell find -E ${BUILD_DIR} -regex '.*\.(wasm|js|html)'  -exec gzip -k -f --best "{}" \; )
	$(shell find -E ${BUILD_DIR} -regex '.*\.(wasm|js|html)'  -exec brotli -f -Z -w 24 "{}" \; )



generate: ## Executes go generate
	${GO} generate ./...

setup: installGolangCi ## Installs golangci-lint

installGolangCi:
	mkdir -p ${TOOLSDIR}
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(TMP_DIR) $(GOLANGCI_LINT_VERSION)


help: ## Shows this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help

