.PHONY: clean test

GOROOT=${shell echo ${HOME}}/goroot
GOPATH=${shell echo ${HOME}}/go2
PKG=github.com/kigichang/goscala

test:
	env GOROOT=${GOROOT} GOPATH=${GOPATH} ${GOROOT}/bin/go test -v -cover -gcflags -G=3 .
	env GOROOT=${GOROOT} GOPATH=${GOPATH} ${GOROOT}/bin/go test -v -cover -gcflags -G=3 ${PKG}/either
	env GOROOT=${GOROOT} GOPATH=${GOPATH} ${GOROOT}/bin/go test -v -cover -gcflags -G=3 ${PKG}/opt
<<<<<<< HEAD
	env GOROOT=${GOROOT} GOPATH=${GOPATH} ${GOROOT}/bin/go test -v -cover -gcflags -G=3 ${PKG}/try
=======
	env GOROOT=${GOROOT} GOPATH=${GOPATH} ${GOROOT}/bin/go test -v -cover -gcflags -G=3 ${PKG}/slices
	env GOROOT=${GOROOT} GOPATH=${GOPATH} ${GOROOT}/bin/go test -v -cover -gcflags -G=3 ${PKG}/try
	
>>>>>>> 0ec6c5065beced5aa1c8726cf96ee1da6ef6d566

tidy:
	env GOROOT=${GOROOT} GOPATH=${GOPATH} ${GOROOT}/bin/go mod tidy