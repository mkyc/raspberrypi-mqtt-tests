OUTPUT := ./out/

all: clean get build

build:
	go vet ./cmd/...
	go fmt ./cmd/...
	go build -x -o $(OUTPUT)sender github.com/mkyc/raspberrypi-mqtt-tests/cmd/sender
	go build -x -o $(OUTPUT)receiver github.com/mkyc/raspberrypi-mqtt-tests/cmd/receiver

get:
	go get -d -v ./...

clean:
	go clean
	rm -rf $(OUTPUT)
