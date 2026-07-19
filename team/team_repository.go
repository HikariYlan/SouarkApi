package team

import (
	"souark/api/base"
	"souark/api/services/mongodb"
)

var repository = base.Repository[Team]{
	Collection: mongodb.Collection("team"),
}

func Repository() base.Repository[Team] {
	return repository
}
