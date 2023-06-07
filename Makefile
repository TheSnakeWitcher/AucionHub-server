##################################################
# variables
##################################################


BIN=./bin
APP="AuctionHub-server"


##################################################
# rules
##################################################


build:
	make build-linux
	make build-windows

build-linux:
	go env -w GOOS=linux && go build -o $(BIN)/ ./...

build-windows:
	go env -w GOOS=windows && go build -o $(BIN)/ ./...

clean:
	make clean_windows
	make clean_linux

clean-linux:
	rm ./$(BIN)/$(APP)

clean-windows:
	rm ./$(BIN)/$(APP).exe

run:
	go run ./...

test:
	go test ./... -v

contract:
	abigen --abi "../AuctionHub-contract/data/AuctionHub.abi" --pkg AuctionHub --out ./AuctionHub/AuctionHub.go

gen-proto:
	cd data && protoc --go_out=.. --go-grpc_out=.. ./*.proto

.PHONY:  build build-linux build-windows \
 	 	 clean clean-linux clean-windows \
 	 	 run test \
