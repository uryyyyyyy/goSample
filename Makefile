# use go1.8

setup:
	go get github.com/Masterminds/glide
	go get github.com/golang/lint/golint
	go get golang.org/x/tools/cmd/goimports


deps:
	glide install

build:
	go build -v -o ./bin/goSample ./src/main.go

test:
	echo "test"

lint:
	go fmt ./src/...
	go vet ./src/...
	golint -min_confidence=0.1 ./src/...


