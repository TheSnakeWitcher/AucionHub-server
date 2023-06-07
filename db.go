package main

import (
    "fmt"
    "context"
    "log"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func InitDbClient(Config Configuration) *mongo.Client {
    DbUri := fmt.Sprintf("%s://%s:%d",
        Config.DbDriver,
        Config.DbHost,
        Config.DbPort,
    )
	// clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	clientOptions := options.Client().ApplyURI(DbUri)
    ctx := context.Background()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
        // InitErrCheck(InitDbErr)
	}
    // Logger.Trace().Msg("db connection seted")

    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal(err)
        // InitErrCheck(InitDbErr)
    }
    // Logger.Trace().Msg("db connection verified")

    return client
    
}
