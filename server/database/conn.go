package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const URI = "mongodb://root:msdnmm@192.168.0.152:27017/"

var Client, MongodbErr = mongo.Connect(context.TODO(), options.Client().ApplyURI(URI))
var UsersColl = Client.Database("golang_server").Collection("users")
