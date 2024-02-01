package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"main/pkg/conf"
	"main/pkg/logger"
	"time"
)

var Client *mongo.Client
var DB *mongo.Database

func ConnectMongo() {
	uri := conf.GetConf("MONGO_URI")
	logger.DebugLog("MONGO_URI: " + uri)
	if uri == "" {
		logger.FatalLog("MONGO_URI not found.")
	}
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	Client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		logger.ErrorLog(err)
		time.Sleep(time.Second * 5)
		Client, err = mongo.Connect(context.TODO(), opts)
		if err != nil {
			logger.FatalLog(err.Error())
		}
	}
	var result bson.D
	err = Client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result)
	if err != nil {
		logger.FatalLog(err.Error())
	}
	DB = Client.Database("bd")
}
