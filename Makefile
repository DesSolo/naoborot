PROJECT_NAME=$(shell basename "$(PWD)")
BIN_DIR="bin"

_cp_certs:
	@ mkdir ${BIN_DIR}
	@cp certs/* ${BIN_DIR}/

clean:
	@rm -rf ${BIN_DIR}

build: clean _cp_certs
	@go build -ldflags "-s -w" -o ${BIN_DIR}/${PROJECT_NAME} *.go

run: build
	@LISTEN_PORT=9000 ./${BIN_DIR}/${PROJECT_NAME}