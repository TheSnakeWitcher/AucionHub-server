package main

import (
    "context"

	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"
)

const (
    auctionColl string = "auctions" 
    bidsColl string = "bids" 
    confColl string = "config" 
)

type Service struct {
    db *mongo.Database
}

func NewService(client *mongo.Client,Config Configuration) *Service {
    return &Service{db: client.Database(Config.DbName)}
}

func (svc Service) collection(coll string) *mongo.Collection {
    return svc.db.Collection(coll)
}


func (svc *Service) InsertAuction(log any) *mongo.InsertOneResult {
    return svc.insertItem(log,auctionColl)
}

func (svc Service) ListAuction() []bson.M {
    return svc.listCollItems(auctionColl)
}

func (svc Service) GetAuction(seller string) (bson.M,error) {
    return svc.getItem(seller,auctionColl)
}

func (svc *Service) UpdateAuction(filter,replacement bson.D) *mongo.UpdateResult {
    return svc.updateItem(filter,replacement,auctionColl)
}

func (svc *Service) DeleteAuction(seller string) *mongo.DeleteResult {
    return svc.deleletItem(seller,bidsColl)
}

func (svc Service) ListBids() []bson.M {
    return svc.listCollItems(bidsColl)
}

func (svc Service) GetBid(seller string) (bson.M,error) {
    return svc.getItem(seller,bidsColl)
}

func (svc *Service) InsertBid(log any) *mongo.InsertOneResult {
    return svc.insertItem(log,bidsColl)
}

func (svc *Service) UpdateBid(filter,replacement bson.D) *mongo.UpdateResult {
    return svc.updateItem(filter,replacement,bidsColl)
}

func (svc *Service) DeleteBid(seller string) *mongo.DeleteResult {
    return svc.deleletItem(seller,bidsColl)
}


func (svc Service) listCollItems(collName string ) []bson.M {
    var result []bson.M
    coll := svc.collection(collName)
    ctx := context.Background()

    cursor,err := coll.Find(ctx,nil)
    if err != nil {
        panic(err)
    }

    cursor.All(ctx,result)
    return result
}

func (svc *Service) deleletItem(seller string,collName string) *mongo.DeleteResult {
    coll := svc.collection(collName)
    result, err := coll.DeleteOne(context.Background(), seller)
    if err != nil {
            panic(err)
    }
    return result
}

func (svc *Service) insertItem(log any,collName string) *mongo.InsertOneResult {
    // TODO: implement convertion from log to auctionData
    coll := svc.collection(collName)
    // auctionFromLog(log)
    auctionData := bson.D{{"fullName", "User 1"}, {"age", 30}}
    result, err := coll.InsertOne(context.Background(), auctionData)
    if err != nil {
        panic(err)
    }
    return result
}

func (svc Service) getItem(seller string,collName string) (bson.M,error) {
    var result bson.M
    // filter := bson.M{seller} // try with filter too
    coll := svc.collection(collName)
    err := coll.FindOne(context.TODO(), seller).Decode(&result)
    if err != nil {
        return nil , err
    }
    return result , nil
}

func (svc *Service) updateItem(filter,replacement bson.D,collName string) *mongo.UpdateResult {
    coll := svc.collection(collName)
    result, err := coll.ReplaceOne(context.Background(), filter, replacement)
    if err != nil {
        panic(err)
    }
    return result
}
