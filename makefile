build-local:
	go get ./...
	go build -o react-static-serve

install:
	make build-local
	mv react-static-serve /usr/local/bin/
