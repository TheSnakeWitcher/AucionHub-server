package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
    InitErr error
    DbClient *mongo.Client
    client *ethclient.Client
)

func InitErrCheck(err error) {
    if err != nil {
        panic(err)
    }
}

func init() {
    InitErr = InitConfig()
	InitErrCheck(InitErr)

    DbClient = InitDbClient(Config)

    client, InitErr = ethclient.Dial(Config.Web3Provider)
	InitErrCheck(InitErr)
}

func main() {
    fmt.Println("AuctionHub-server starts...")
    defer fmt.Println("AuctionHub-server stops...")

	defer func() {
        if err := DbClient.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

    svc := NewService(DbClient,Config)
    srv := NewServer(*svc,client,Config)

    errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
        srvAddres := fmt.Sprintf("%s:%s",Config.ServerHost,Config.ServerPort)
        errs <- srv.Start(srvAddres)
	}()

    log.Fatal(<-errs)
}
