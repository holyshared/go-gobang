.PHONY: setup build test clean

setup:
	go get github.com/gin-gonic/gin
	go get github.com/olahol/melody
	go get github.com/gin-gonic/contrib/static
	go get github.com/Sirupsen/logrus

build:
	go build -o server main.go

test:
	go test ./gobang
	go test ./gobang-server

clean:
	rm server
