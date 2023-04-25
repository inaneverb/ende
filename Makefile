OUTPUT_BUILD:=./bin/ende

GO_BUILD_FILE:=./main.go
GO_BUILD_TAGS:=

GOPRIVATE :=
GO_ENV := GONOPROXY=${GOPRIVATE} GONOSUMDB=${GOPRIVATE} GOPRIVATE=${GOPRIVATE}

env:
	${GO_ENV} go env

codegen:
	${GO_ENV} go generate ./...

deps:
	${GO_ENV} go mod tidy
	${GO_ENV} go mod download

.PHONY: build
build: env deps codegen
	${GO_ENV} go build -tags=${GO_BUILD_TAGS} -o ${OUTPUT_BUILD} ${GO_BUILD_FILE}

.PHONY: install
install: build
	mkdir -p /opt/local/bin || true
	cp -f ${OUTPUT_BUILD} /opt/local/bin/

.PHONY: run
run: build
	${OUTPUT_BUILD}
