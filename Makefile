
BIN = dopen
COMPILE = go build
MAIN = main/main.go

$(BIN):
	$(COMPILE) -o $(BIN) bin/$(MAIN) 

clean:
	rm -rf $(BIN)
