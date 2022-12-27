package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/donggyuLim/suino-server/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const URL = ""

var dbName string
var url string
var db *mongo.Client

func DBConnect() {
	url = utils.LoadENV("DBURL")
	dbName = utils.LoadENV("DBNAME")
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(url).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("============DB connect==================")
	db = client
}


func Close() {
	err := db.Disconnect(context.TODO())
	utils.HandleErr(err)
	fmt.Println("=========Connection to MongoDB closed=============")
}

func Insert(collectionName string, data interface{}) error {
	exp := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), exp)
	defer cancel()
	collection := db.Database(dbName).Collection(collectionName)
	insertResult, err := collection.InsertOne(ctx, data)
	if err != nil {
		return err
	}

	fmt.Println("Insert Complete", insertResult.InsertedID)
	return nil
}

// func Exsits(dbName, collectionName, key, value string) bool {
// 	collection := db.Database(dbName).Collection(collectionName)
// 	filter := bson.D{{Key: key}}
// 	results := collection.FindOne(context.TODO(), filter)

// }

func FindOne(collectionName, key, value string, data interface{}) error {
	exp := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), exp)
	defer cancel()
	collection := db.Database(dbName).Collection(collectionName)
	filter := bson.D{{Key: key, Value: value}}
	err := collection.FindOne(ctx, filter).Decode(&data)
	return err
}

// 어떻게 찍히나 확인해야함.
func Find(collectionName, key, value string, limit int64) ([]bson.M, error) {
	exp := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), exp)
	defer cancel()
	collection := db.Database(dbName).Collection(collectionName)
	findOptions := options.Find()
	findOptions.SetLimit(limit)
	findOptions.SetSort(bson.D{{Key: "timestamp", Value: -1}})
	var filter primitive.D
	if key == "" && value == "" {
		filter = bson.D{}
	} else {
		filter = bson.D{{Key: key, Value: value}}
	}

	cur, _ := collection.Find(ctx, filter, findOptions)
	var curs []bson.M
	err := cur.All(context.TODO(), &curs)
	return curs, err
}

func FindAndUpdate(collectionName, key, value string, data interface{}) {
	exp := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), exp)
	defer cancel()
	collection := db.Database(dbName).Collection(collectionName)
	filter := bson.D{{Key: key, Value: value}}
	collection.FindOneAndReplace(ctx, filter, data)
}
