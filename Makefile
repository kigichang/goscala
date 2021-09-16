.PHONY: clean test

GOROOT=${shell echo ${HOME}}/goroot
GOPATH=${shell echo ${HOME}}/go2
PKG=github.com/kigichang/goscala

test:
	env GOROOT=${GOORT} GOPATH=${GOPATH} ${GOROOT}/bin/go test -gcflags -G=3 .

tidy:
	env GOROOT=${GOORT} GOPATH=${GOPATH} ${GOROOT}/bin/go mod tidy