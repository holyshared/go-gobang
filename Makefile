.PHONY: setup

setup:
	go get github.com/gin-gonic/gin
	go get github.com/olahol/melody
	go get github.com/gin-gonic/contrib/static
	go get github.com/Sirupsen/logrus

.PHONY: build

build:
	go build -o server main.go

.PHONY: test

test:
	go test ./gobang
	go test ./gobang-server
