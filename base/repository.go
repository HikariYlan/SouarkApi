package base

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// Generic repository type
type Repository[Entity any] struct {
	Collection *mongo.Collection
}

// Returns filtered data from the collection, or an error in case of failure
func (repository Repository[Entity]) FindAll(filter bson.D) ([]Document[Entity], error) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := repository.Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	data := make([]Document[Entity], 0)
	if err = cursor.All(context.TODO(), &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (repository Repository[Entity]) InsertOne(document Entity) (Document[Entity], error) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var newDocument Document[Entity]

	result, err := repository.Collection.InsertOne(ctx, document)

	if err != nil {
		return newDocument, err
	}

	filter := bson.D{{Key: "_id", Value: result.InsertedID}}
	singleResult := repository.Collection.FindOne(ctx, filter)

	if err = singleResult.Decode(&newDocument); err != nil {
		return newDocument, err
	}
	return newDocument, nil
}

func (repository Repository[Entity]) FindById(id string) (Document[Entity], error) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var document Document[Entity]

	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return document, err
	}

	filter := bson.D{{Key: "_id", Value: objectID}}

	result := repository.Collection.FindOne(ctx, filter)

	if err := result.Decode(&document); err != nil {
		return document, err
	}
	return document, nil
}

func (repository Repository[Entity]) Delete(id string) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}

	filter := bson.D{{Key: "_id", Value: objectID}}

	result, err := repository.Collection.DeleteOne(ctx, filter)

	if err != nil {
		return false, err
	}

	if result.DeletedCount == 0 {
		return false, err
	}

	return true, nil
}
