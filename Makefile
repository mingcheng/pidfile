VERSION=`date +%Y%m%d`

ifneq ("$(wildcard /go)","")
   GOPATH=/go
   GOROOT=/usr/local/go
endif

GO_ENV=CGO_ENABLED=0
GO_FLAGS=-ldflags=""
GO=env $(GO_ENV) $(GOROOT)/bin/go

PACKAGES=`go list ./... | grep -v /vendor/`
GOFILES=`find . -name "*.go" -type f -not -path "./vendor/*"`

build: 
	@$(GO) build $(GO_FLAGS) .

fmt:
	@gofmt -s -w ${GOFILES}

list:
	@echo ${PACKAGES}
	@echo ${VETPACKAGES}
	@echo ${GOFILES}

test:
	@go test -cpu=1,2,4 -v -tags integration ./...

clean:
	@$(GO) clean ./...

.PHONY: fmt  test clean target   
