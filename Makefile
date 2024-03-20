.PHONY: install

install:
	go build -o lazywoo
	mv lazywoo $(GOPATH)/bin