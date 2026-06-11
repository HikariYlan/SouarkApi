package mongodb

import (
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// Open a MongoDB connection and returns the application database.
// Required values (username, password, database) are being read from the environment file
func connectToDB() *mongo.Database {
	err := godotenv.Load()
	if err != nil {
		return nil
	}

	username := os.Getenv("MONGO_USER")
	password := os.Getenv("MONGO_PASSWORD")
	host := os.Getenv("MONGO_HOST")
	port := os.Getenv("MONGO_PORT")

	uri := "mongodb://" + username + ":" + password + "@" + host + ":" + port
	opts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(opts)
	if err != nil {
		panic(err)
	}

	db := os.Getenv("MONGO_DB")
	return client.Database(db)
}

// MongoDB singleton
var db = connectToDB()

// Returns a collection identified by its name
func Collection(name string) *mongo.Collection {
	return db.Collection(name)
}
