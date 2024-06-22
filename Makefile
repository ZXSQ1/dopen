
BIN = dopen
COMPILE = go build
MAIN = main/main.go

$(BIN):
	$(COMPILE) -o $(BIN) $(MAIN) 

clean:
	rm -rf $(BIN)
