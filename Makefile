SRCS != find . -type f -name '*.go' | grep -v _test.go
TESTDIRS != find . -type f -name '*_test.go' | xargs dirname | uniq

BIN != basename $(CURDIR)

$(BIN): $(SRCS)
	go build

run: $(SRCS)
	go run hello.go

test:
	go test -v $(TESTDIRS)

clean:
	rm -f ./$(BIN)

deps:
	go mod tidy

install:
	go install

uninstall:
	rm -f $(shell which $(BIN))

.PHONY: run test clean deps install uninstall
