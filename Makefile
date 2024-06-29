
MAIN = main.go
COMPILE = go build


VERSION = 0.0.1
BIN = dopen

build:
	GOOS=windows GOARCH=amd64 $(COMPILE) -o bin/$(BIN)-$(VERSION)-windows.exe $(MAIN)
	GOOS=linux GOARCH=amd64 $(COMPILE) -o bin/$(BIN)-$(VERSION)-linux $(MAIN)

clean:
	rm -rf bin/*
