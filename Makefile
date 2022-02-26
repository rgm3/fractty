BIN=fractty

$(BIN): go.mod main.go
	go build -o $(BIN) main.go

.PHONY: all
all: $(BIN)

.PHONY: build
build: $(BIN)

.PHONY: test
test:
	go test -v main.go

.PHONY: run
run: $(BIN)
	./$(BIN)

.PHONY: clean
clean:
	go clean
	rm $(BIN)
