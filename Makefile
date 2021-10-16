.PHONY: clean test

GOROOT=/Users/kigi/goroot
GOPATH=/Users/kigi/go2
PKG=github.com/kigichang/goscala

test:
	env GOROOT=${GOROOT} GOPATH=${GOPATH} ${GOROOT}/bin/gofmt -w .
	env GOROOT=${GOROOT} GOPATH=${GOPATH} ${GOROOT}/bin/go test -v -cover -gcflags -G=3 .
	env GOROOT=${GOROOT} GOPATH=${GOPATH} ${GOROOT}/bin/go test -v -cover -gcflags -G=3 ${PKG}/either
	env GOROOT=${GOROOT} GOPATH=${GOPATH} ${GOROOT}/bin/go test -v -cover -gcflags -G=3 ${PKG}/future
	env GOROOT=${GOROOT} GOPATH=${GOPATH} ${GOROOT}/bin/go test -v -cover -gcflags -G=3 ${PKG}/iter
#	env GOROOT=${GOROOT} GOPATH=${GOPATH} ${GOROOT}/bin/go test -v -cover -gcflags -G=3 ${PKG}/maps
	env GOROOT=${GOROOT} GOPATH=${GOPATH} ${GOROOT}/bin/go test -v -cover -gcflags -G=3 ${PKG}/opt
	env GOROOT=${GOROOT} GOPATH=${GOPATH} ${GOROOT}/bin/go test -v -cover -gcflags -G=3 ${PKG}/try
	env GOROOT=${GOROOT} GOPATH=${GOPATH} ${GOROOT}/bin/go test -v -cover -gcflags -G=3 ${PKG}/seq
	env GOROOT=${GOROOT} GOPATH=${GOPATH} ${GOROOT}/bin/go test -v -cover -gcflags -G=3 ${PKG}/slices
	
tidy:
	env GOROOT=${GOROOT} GOPATH=${GOPATH} ${GOROOT}/bin/go mod tidy