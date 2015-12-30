PACKAGE=github.com/maraino/go-enum

all:
	go build $(PACKAGE)

test:
	go test -cover $(PACKAGE)

cover:
	go test -coverprofile=c.out $(PACKAGE)
	go tool cover -html=c.out
