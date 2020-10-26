package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type TZConvertion struct {
	TimeZone       string `bson:"timezone" json:"timeZone"`
	TimeDifference string `bson:"timedifference" json:"timeDifference"`
}

type Repository struct {
	ctx           context.Context
	client        *mongo.Client
	dbServer      string
	dbDatabase    string
	dbCollection  string
	collection    *mongo.Collection
	clientOptions *options.ClientOptions
}

func NewRepository(dbServer string, dbDatabase string, dbCollection string, dbUser string, dbPass string) *Repository {
	repo := new(Repository)
	repo.ctx = context.TODO()
	repo.dbServer = dbServer
	repo.dbDatabase = dbDatabase
	repo.dbCollection = dbCollection
	repo.clientOptions = options.Client().ApplyURI("mongodb://" + dbServer).SetAuth(
		options.Credential{
			Username: dbUser,
			Password: dbPass,
		})

	client, err := mongo.Connect(repo.ctx, repo.clientOptions)
	if err != nil {
		log.Fatal("Couldn't connect to MongoDB")
	}
	repo.client = client
	repo.collection = client.Database(dbDatabase).Collection(dbCollection)
	return repo
}

func (repo *Repository) Close() {
	err := repo.client.Disconnect(repo.ctx)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}

func (repo *Repository) FindAll() ([]TZConvertion, error) {
	var tzcs []TZConvertion
	cur, err := repo.collection.Find(repo.ctx, bson.D{{}})

	for cur.Next(repo.ctx) {
		var elem TZConvertion
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		tzcs = append(tzcs, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(repo.ctx)
	return tzcs, err
}

func (repo *Repository) FindByTimeZone(tz string) (TZConvertion, error) {
	var tzc TZConvertion
	err := repo.collection.FindOne(repo.ctx, bson.M{"timezone": tz}).Decode(&tzc)
	if err != nil {
		log.Fatal(err)
	}
	return tzc, err
}

func (repo *Repository) Insert(tzc TZConvertion) error {
	_, err := repo.collection.InsertOne(repo.ctx, tzc)
	return err
}

func (repo *Repository) Delete(tzc TZConvertion) error {
	_, err := repo.collection.DeleteMany(context.TODO(), bson.M{"timezone": tzc.TimeZone})
	return err
}

func (repo *Repository) Update(tz string, tzc TZConvertion) error {
	filter := bson.D{{"timezone", tz}}
	_, err := repo.collection.UpdateOne(repo.ctx, filter, &tzc)
	return err
}
