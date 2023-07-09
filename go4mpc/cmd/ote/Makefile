BINARY_NAME=ote

build:
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME} main.go

run:
	./${BINARY_NAME} -i	100 1

clean:
	@echo $(MPCFGO_ROOT)
	go clean
 #rm ${BINARY_NAME}-darwin
	rm ${BINARY_NAME}
 #rm ${BINARY_NAME}-windows