

GO := go
GOFMT := gofmt
# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell $(GO) env GOBIN))
GOBIN=$(shell $(GO) env GOPATH)/bin
else
GOBIN=$(shell $(GO) env GOBIN)
endif

check-gopath:
ifndef GOPATH
	$(error GOPATH is not set)
endif
.PHONY: check-gopath

verify: check-gopath
	$(GO) vet \
		./cmd/... \
		./pkg/... 
.PHONY: verify

binary: verify
	$(GO) build ./cmd/uninterface
.PHONY: binary

container:
	docker build -t quay.io/famargon/uninterface:latest .
	docker push quay.io/famargon/uninterface:latest
.PHONY: container