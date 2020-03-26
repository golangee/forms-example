## the binary name
ARTIFACT_NAME = blub

## for the module itself
MODULE_PATH = github.com/worldiety/wtk-example

## the path which contains the main package to execute
MAIN_PATH = github.com/worldiety/wtk-example/cmd/wasm

## for ldflags replacement
BUILD_FILE_PATH = ${MODULE_PATH}

## which linter version to use?
GOLANGCI_LINT_VERSION = v1.24.0

LDFLAGS = -X $(MODULE_PATH).BuildGitCommit=$(CI_COMMIT_SHA) -X $(MODULE_PATH).BuildGitBranch=$(CI_COMMIT_REF_NAME)

TMP_DIR = $(TMPDIR)/$(MODULE_PATH)
BUILD_DIR = .build
TOOLSDIR = $(TMP_DIR)
GO = go
GOLANGCI_LINT = ${TOOLSDIR}/golangci-lint
GOLINT = ${TOOLSDIR}/golint
TMP_GOPATH = $(CURDIR)/${TOOLSDIR}/.gopath

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
	${GO} generate ./...
	GOOS=js GOARCH=wasm ${GO} build -ldflags "${LDFLAGS}" -o ${BUILD_DIR}/app.wasm ${MAIN_PATH}


	cp "${GOROOT}/misc/wasm/wasm_exec.js" ${BUILD_DIR}
	cp -r static/. ${BUILD_DIR}


compress: build ## Applies gzip and brotli compression to build files
	$(shell find -E ${BUILD_DIR} -regex '.*\.(wasm|js|html)'  -exec gzip -k -f --best "{}" \; )
	$(shell find -E ${BUILD_DIR} -regex '.*\.(wasm|js|html)'  -exec brotli -f -Z -w 24 "{}" \; )


run: clean build ## Starts the compiled program
	${GO} build -o ${BUILD_DIR}/srv ${MODULE_PATH}/cmd/srv
	${BUILD_DIR}/srv -d=${BUILD_DIR}


generate: ## Executes go generate
	${GO} generate

setup: installGolangCi ## Installs golangci-lint

installGolangCi:
	mkdir -p ${TOOLSDIR}
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(TOOLSDIR) $(GOLANGCI_LINT_VERSION)


help: ## Shows this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help

