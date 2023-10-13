package database

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoDb struct {
	collectionName string
	ctx            context.Context
	dbName         string
	MongoURL       string
}

func NewMongoCollection(collaction string) *mongoDb {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatal("Error Loading env File")
	}

	return &mongoDb{
		collectionName: collaction,
		ctx:            context.Background(),
		dbName:         os.Getenv("MONGODB_DB_NAME"),
		MongoURL:       os.Getenv("MONGODB_URL"),
	}
}

func (m *mongoDb) connect() (Client *mongo.Client, Error error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(m.MongoURL))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(m.ctx)

	return client, err
}

func (m *mongoDb) GetCollection() (collection *mongo.Collection) {
	client, err := m.connect()
	if err != nil {
		log.Fatal(err)
	}
	var mongoCollection *mongo.Collection = client.Database(m.dbName).Collection(m.collectionName)

	return mongoCollection
}
