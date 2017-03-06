# use go1.8

setup:
	go get github.com/Masterminds/glide
	go get github.com/golang/lint/golint
	go get golang.org/x/tools/cmd/goimports


deps:
	glide install

build:
	go build -v -o ./bin/goSample ./main.go

run:
	go run ./main.go

run_build:
	./bin/goSample

test:
	glide novendor | xargs go test

lint:
	glide novendor | xargs go fmt
	glide novendor | xargs go vet
	glide novendor | xargs golint -min_confidence=0.8

