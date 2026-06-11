package player

import (
	"souark/api/base"
	"souark/api/services/mongodb"
)

var repository = base.Repository[Player]{
	Collection: mongodb.Collection("player"),
}

func Repository() base.Repository[Player] {
	return repository
}
