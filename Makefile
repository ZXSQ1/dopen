
BIN = dopen
COMPILE = go build

$(BIN):
	$(COMPILE) -o $(BIN) *.go */*.go

clean:
	rm -rf $(BIN)
