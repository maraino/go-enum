PACKAGE=github.com/maraino/go-enum

all:
	go build $(PACKAGE)

test:
	go test -cover $(PACKAGE)

run:
	go run example/main.go

cover:
	go test -coverprofile=c.out $(PACKAGE)
	go tool cover -html=c.out
