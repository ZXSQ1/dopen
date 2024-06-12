
BIN = devdocs-tui
COMPILE = go build

$(BIN):
	$(COMPILE) -o $(BIN) *.go

clean:
	rm -rf $(BIN)
