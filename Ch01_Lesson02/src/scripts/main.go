package main

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"log"
)

type timeZoneConversion struct {
	TimeZone		string `json:"timeZone"`
	TimeDifference	string `json:"timeDifference"`
	Name			string `json:"name"`
}

type tzs []timeZoneConversion

func main() {
	const database = "packt"
	const username = "root"
	const password = "example"
	const collectionName = "timeZones"

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(
		options.Credential{
			Username: username,
			Password: password,
		})


	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Couldn't connect to MongoDB")
	}
	defer client.Disconnect(context.TODO())

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Connected to MongoDB!")

	collection := client.Database(database).Collection(collectionName)
	err = collection.Drop(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	log.Printf(`"%s" database deleted`, database)

	startDatabase := client.Database(database)
	timezoneCollection := startDatabase.Collection(collectionName)
	log.Printf(`"%s" database "%s" collection created`, database, collectionName)

	data, err := ioutil.ReadFile("all_timezones.json")
	if err != nil {
		log.Fatal("Couldn't open file")
	}

	var timeZones tzs
	err = json.Unmarshal(data, &timeZones)
	if err != nil {
		log.Fatal("Couldn't unmarshall Json")
	}

	i := 0
	for _, v := range timeZones {
		timezoneCollection.InsertOne(context.TODO(), v)
		i++
	}
	log.Printf("%d record inserted", i)
}