package base

import "go.mongodb.org/mongo-driver/v2/bson"

// Encapsulates an entity in a MongoDB document
type Document[E any] struct {
	Id     bson.ObjectID `bson:"_id" json:"_id"`
	Entity E             `bson:"inline" json:",inline"`
}
