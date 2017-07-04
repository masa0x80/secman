NAME := secman
REVISION := $(shell git rev-parse --short HEAD)
VERSION = $(shell grep 'Version string =' $</version.go | sed -E 's/.*"(.+)"$$/\1/')
LDFLAGS = -X 'main.version=$(VERSION)' \
					-X 'main.revision=$(REVISION)'

setup:
	go get github.com/Masterminds/glide
	go get github.com/golang/lint/golint
	go get golang.org/x/tools/cmd/goimports

test: deps
	go test $$(glide novendor)

deps: setup
	glide install

update: setup
	glide update

lint: setup
	go vet $$(glide novendor)
	for pkg in $$(glide novendor -x); do \
		golint -set_exit_status $$pkg || exit $$?; \
	done

fmt: setup
	goimports -w $$(glide nv -x)

bin/%: cmd/% deps
	go build -ldflags "$(LDFLAGS)" -o $@ $</*.go

.PHONY: setup deps update test lint fmt build
