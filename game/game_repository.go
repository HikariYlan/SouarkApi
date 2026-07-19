package game

import (
	"souark/api/base"
	"souark/api/services/mongodb"
)

var repository = base.Repository[Game]{
	Collection: mongodb.Collection("game"),
}

func Repository() base.Repository[Game] {
	return repository
}
