OUTPUT:=./bin/ende
COMMIT:=$(shell git rev-parse --short HEAD)
DATE:=$(shell date '+%Y%m%d%H%M%S; unix: %s')

GO_MOD:=github.com/inaneverb/ende
GO_FILE:=./main.go
GO_TAGS:=
GO_FLAGS:=-ldflags="\
-X '${GO_MOD}/internal/pkg/version.Commit=${COMMIT}'\
-X '${GO_MOD}/internal/pkg/version.Date=${DATE}'\
-s\
"
GO_PRIVATE :=
GO_ENV := GONOPROXY=${GO_PRIVATE} GONOSUMDB=${GO_PRIVATE} GOPRIVATE=${GO_PRIVATE}

env:
	${GO_ENV} go env

codegen:
	${GO_ENV} go generate ./...

clean:
	rm -f ${OUTPUT} || true

deps:
	${GO_ENV} go mod tidy
	${GO_ENV} go mod download

.PHONY: build
build: clean deps codegen
	${GO_ENV} go build -tags=${GO_TAGS} ${GO_FLAGS} -o ${OUTPUT} ${GO_FILE}

.PHONY: install
install: clean deps codegen
	${GO_ENV} go install -tags=${GO_TAGS} ${GO_FLAGS}

.PHONY: run
run: build
	${OUTPUT}
