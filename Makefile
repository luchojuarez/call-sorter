BINARY=service

build:
	go build -o ${BINARY} cmd/*.go

run-example:
	./${BINARY} | jq