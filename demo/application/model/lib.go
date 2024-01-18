package model

import "go.mongodb.org/mongo-driver/mongo"

type Lib struct {
	MongoDb *mongo.Database
}
