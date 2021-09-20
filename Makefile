.PHONY: clean test

GOROOT=${shell echo ${HOME}}/goroot
GOPATH=${shell echo ${HOME}}/go2
PKG=github.com/kigichang/goscala

test:
	env GOROOT=${GOORT} GOPATH=${GOPATH} ${GOROOT}/bin/go test -v -cover -gcflags -G=3 .
	env GOROOT=${GOORT} GOPATH=${GOPATH} ${GOROOT}/bin/go test -v -cover -gcflags -G=3 ${PKG}/monad

tidy:
	env GOROOT=${GOORT} GOPATH=${GOPATH} ${GOROOT}/bin/go mod tidy