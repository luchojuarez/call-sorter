BINARY=service

build:
	go build -o ${BINARY} cmd/*.go

run:
	./${BINARY}

web: build run