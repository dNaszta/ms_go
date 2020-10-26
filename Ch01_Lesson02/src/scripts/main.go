package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type timeZoneConversion struct {
	TimeZone		string `bson:"timeZone" json:"timeZone"`
	TimeDifference	string `bson:"timeDifference" json:"timeDifference"`
	Name			string `bson:"name" json:"name"`
}

type tzs []timeZoneConversion

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("Couldn't connect to MongoDB")
	}
}