
PKG=./...
TEST=.

clean:
	rm *_gen.go

gen:
	go generate ./...
	go fmt ./...

test:
	go test -v $(PKG) -run $(TEST)
