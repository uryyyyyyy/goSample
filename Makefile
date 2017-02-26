# use go1.8

setup:
	go get github.com/Masterminds/glide
	go get github.com/golang/lint

deps:
	glide install

build:
	go build -v -o ./bin/goSample ./src/main.go

test:
	echo "test"

fmt:
	go fmt ./src/...

