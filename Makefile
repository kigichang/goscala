.PHONY: clean test

PKG=github.com/kigichang/goscala

test:
	env GOROOT=${shell echo ${HOME}}/goroot ${shell echo ${HOME}}/goroot/bin/go tool go2go test
clean:
	-rm *.go