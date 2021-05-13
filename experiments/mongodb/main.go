package main

import (
	"context"
	"fmt"
	"time"

	"github.com/LeeDark/go-experience/experiments/mongodb/maybe"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const localhostURI = "mongodb://127.0.0.1:27017"

type MNPCacheItem struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	InstanceID int64              `bson:"instid"`
	MSISDN     string             `bson:"msisdn"`
	CountryID  maybe.Int64        `bson:"countryid"`
	MCC        string             `bson:"mcc"`
	MNC        string             `bson:"mnc"`
	IsPorted   maybe.Bool         `bson:"isported"`
	IsValid    maybe.Bool         `bson:"isvalid"`
	IsNumberOK maybe.Bool         `bson:"isok"`
	SourceID   int64              `bson:"sourceid"`
	QueryDT    time.Time          `bson:"querydt"`
	FromID     string             `bson:"fromid"`
}

type IndiaCacheItem struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	MSISDN        string             `bson:"msisdn"`
	OriginalMcc   string             `bson:"originalMcc"`
	OriginalMnc   string             `bson:"originalMnc"`
	OriginNetwork string             `bson:"originNetwork"`
	OriginCircle  string             `bson:"originCircle"`
	MCC           string             `bson:"mcc"`
	MNC           string             `bson:"mnc"`
	PortedNetwork string             `bson:"portedNetwork"`
	PortedCircle  string             `bson:"portedCircle"`
	UpdatedAt     time.Time          `bson:"updatedAt"`
}

func GetItem(collection *mongo.Collection, MSISDN string) {
	var cacheItem MNPCacheItem
	filter := bson.M{"msisdn": MSISDN}

	err := collection.FindOne(context.TODO(), filter).Decode(&cacheItem)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Found item for %s with %s:%s\n", MSISDN, cacheItem.MCC, cacheItem.MNC)
}

func AddIndiaItem(collection *mongo.Collection, MSISDN, MCC, MNC string) {
	var cacheItem IndiaCacheItem
	filter := bson.D{primitive.E{Key: "msisdn", Value: MSISDN}}

	err := collection.FindOne(context.TODO(), filter).Decode(&cacheItem)
	if err != nil {
		fmt.Printf("Not found item for %s with %+v\n", MSISDN, cacheItem)

		insertCacheItem := IndiaCacheItem{MSISDN: MSISDN, MCC: MCC, MNC: MNC}
		insertResult, err := collection.InsertOne(context.TODO(), insertCacheItem)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Inserted India cache item with ID:", insertResult.InsertedID)
		return
	}

	fmt.Printf("Found item for %s with %+v\n", MSISDN, cacheItem)

	var updatedDocument IndiaCacheItem
	//opts := options.FindOneAndUpdate().SetUpsert(true)
	opts := options.FindOneAndReplace().SetUpsert(true)

	//updateCacheItem := IndiaCacheItem{MSISDN: MSISDN, MCC: MCC, MNC: MNC}
	//updateCacheItem := bson.D{{"$set", bson.D{{"mcc", MCC}, {"mnc", MNC}}}}

	replaceCacheItem := IndiaCacheItem{MSISDN: MSISDN, MCC: MCC, MNC: MNC}
	err = collection.FindOneAndReplace(context.TODO(), filter, replaceCacheItem, opts).Decode(&updatedDocument)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			return
		}
		fmt.Println(err)
		return
	}

	fmt.Println("Updated India cache item with ID:", updatedDocument.ID)
}

func GetIndiaItem(collection *mongo.Collection, MSISDN string) {
	var cacheItem IndiaCacheItem
	filter := bson.M{"msisdn": MSISDN}

	err := collection.FindOne(context.TODO(), filter).Decode(&cacheItem)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Found item for %s with %s:%s\n", MSISDN, cacheItem.MCC, cacheItem.MNC)
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI(localhostURI))
	if err != nil {
		fmt.Println(err)
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer client.Disconnect(ctx)

	cacheCollection := client.Database("mnpserver2").Collection("cache")

	GetItem(cacheCollection, "4528198446")
	GetItem(cacheCollection, "4520705981")

	indiaCollection := client.Database("indiadb").Collection("cache")

	AddIndiaItem(indiaCollection, "4528198446", "100", "11")
	AddIndiaItem(indiaCollection, "4520705981", "100", "15")
	AddIndiaItem(indiaCollection, "1238198446", "100", "22")
	AddIndiaItem(indiaCollection, "4568198446", "100", "44")

	GetIndiaItem(indiaCollection, "4528198446")
	GetIndiaItem(indiaCollection, "4520705981")
	GetIndiaItem(indiaCollection, "1238198446")
	GetIndiaItem(indiaCollection, "4568198446")
}
