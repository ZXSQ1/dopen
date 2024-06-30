
MAIN = main.go
COMPILE = go build

BIN = dopen

build:
	GOOS=windows GOARCH=amd64 $(COMPILE) -o bin/$(BIN)-windows-amd64.exe $(MAIN)
	GOOS=windows GOARCH=386 $(COMPILE) -o bin/$(BIN)-windows-386.exe $(MAIN)
	GOOS=windows GOARCH=arm $(COMPILE) -o bin/$(BIN)-windows-arm.exe $(MAIN)

	GOOS=linux GOARCH=amd64 $(COMPILE) -o bin/$(BIN)-linux-amd64 $(MAIN)
	GOOS=linux GOARCH=arm $(COMPILE) -o bin/$(BIN)-linux-arm $(MAIN)
	GOOS=linux GOARCH=386 $(COMPILE) -o bin/$(BIN)-linux-386 $(MAIN)
	GOOS=linux GOARCH=arm64 $(COMPILE) -o bin/$(BIN)-linux-arm64 $(MAIN)

	GOOS=darwin GOARCH=amd64 $(COMPILE) -o bin/$(BIN)-darwin-arm $(MAIN)

	GOOS=netbsd GOARCH=amd64 $(COMPILE) -o bin/$(BIN)-netbsd-amd64 $(MAIN)
	GOOS=netbsd GOARCH=arm $(COMPILE) -o bin/$(BIN)-netbsd-arm $(MAIN)
	GOOS=netbsd GOARCH=386 $(COMPILE) -o bin/$(BIN)-netbsd-386 $(MAIN)

	GOOS=openbsd GOARCH=amd64 $(COMPILE) -o bin/$(BIN)-openbsd-amd64 $(MAIN)
	GOOS=openbsd GOARCH=arm $(COMPILE) -o bin/$(BIN)-openbsd-arm $(MAIN)
	GOOS=openbsd GOARCH=386 $(COMPILE) -o bin/$(BIN)-openbsd-386 $(MAIN)

	GOOS=freebsd GOARCH=amd64 $(COMPILE) -o bin/$(BIN)-freebsd-amd64 $(MAIN)
	GOOS=freebsd GOARCH=arm $(COMPILE) -o bin/$(BIN)-freebsd-arm $(MAIN)
	GOOS=freebsd GOARCH=386 $(COMPILE) -o bin/$(BIN)-freebsd-386 $(MAIN)

clean:
	rm -rf bin/*
