package player

import (
	"encoding/json"
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

func ValidateData(req *http.Request) (bool, Player) {
	decoder := json.NewDecoder(req.Body)

	var document Player

	if err := decoder.Decode(&document); err != nil {
		return false, document
	}

	return true, document
}

var controller = base.Controller[Player]{
	Repository:   Repository(),
	CreateFilter: createFilter,
	ValidateData: ValidateData,
	Routes:       []string{"get", "post", "getId", "delete"},
}
