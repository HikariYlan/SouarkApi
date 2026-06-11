package player

import (
	"net/http"
	"souark/api/base"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func createFilter(req *http.Request) bson.D {

	firstName := req.URL.Query().Get("firstName")
	lastName := req.URL.Query().Get("lastName")
	division := req.URL.Query().Get("division")

	var firstNameFilter bson.E
	var lastNameFilter bson.E
	var divisionFilter bson.E

	if firstName != "" {
		firstNameFilter = bson.E{Key: "firstName", Value: firstName}
	}

	if lastName != "" {
		lastNameFilter = bson.E{Key: "lastName", Value: lastName}
	}

	if division != "" {
		divisionFilter = bson.E{Key: "division", Value: division}
	}

	return bson.D{firstNameFilter, lastNameFilter, divisionFilter}
}

var controller = base.Controller[Player]{
	Repository:   Repository(),
	CreateFilter: createFilter,
}
