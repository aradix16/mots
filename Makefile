all: build

dev:
	air .

build:
	go build -o bin/app

clean:
	rm -rf bin
