BINARY_NAME=gtransl

# TODO: parameterize this 
RAW_PATH=${MPCFGO_ROOT}/example/cryptonet/raw


build:
 #GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin main.go
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME} main.go
 #GOARCH=amd64 GOOS=window go build -o ${BINARY_NAME}-windows main.go
	
run:
	./${BINARY_NAME} -L	$(RAW_PATH)
#	./${BINARY_NAME} -L /home/fanyx/mine/mpcfgo/example/cryptonet/raw
#	./${BINARY_NAME} -L /home/fanyx/me/mpcfgo/example/cryptonet/raw

build_and_run: build run

clean:
	@echo $(MPCFGO_ROOT)
	go clean
 #rm ${BINARY_NAME}-darwin
	rm ${BINARY_NAME}
 #rm ${BINARY_NAME}-windows