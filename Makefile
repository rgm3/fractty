BIN=fractty
.PHONY=build test run clean

all: $(BIN)

$(BIN): go.mod main.go
	go build -o $(BIN) main.go

build:
	go build

test:
	go test -v main.go

run: $(BIN)
	./$(BIN)

clean:
	go clean
	rm $(BIN)
