.PHONY: clean test

GOROOT=${shell echo ${HOME}}/goroot
GOPATH=${shell echo ${HOME}}/go2
PKG=github.com/kigichang/goscala

test:
	env GOROOT=${GOROOT} GOPATH=${GOPATH} ${GOROOT}/bin/go test -v -cover -gcflags -G=3 .
	env GOROOT=${GOROOT} GOPATH=${GOPATH} ${GOROOT}/bin/go test -v -cover -gcflags -G=3 ${PKG}/either
	env GOROOT=${GOROOT} GOPATH=${GOPATH} ${GOROOT}/bin/go test -v -cover -gcflags -G=3 ${PKG}/opt
#	env GOROOT=${GOROOT} GOPATH=${GOPATH} ${GOROOT}/bin/go test -v -cover -gcflags -G=3 ${PKG}/slices
	env GOROOT=${GOROOT} GOPATH=${GOPATH} ${GOROOT}/bin/go test -v -cover -gcflags -G=3 ${PKG}/try
	

tidy:
	env GOROOT=${GOROOT} GOPATH=${GOPATH} ${GOROOT}/bin/go mod tidy