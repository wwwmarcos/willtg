.PHONY: build clean deploy

deps:
	go get -v -t -d ./...

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/hello telegram/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose
