package game

import (
	"encoding/json"
	"net/http"
	"souark/api/base"
	"souark/api/services/send"
	"souark/api/team"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func createFilter(req *http.Request) bson.D {
	return bson.D{}
}

func ValidateData(req *http.Request) (bool, Game) {
	decoder := json.NewDecoder(req.Body)

	var document Game

	if err := decoder.Decode(&document); err != nil {
		return false, document
	}

	return true, document
}

func NewGame(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	var config Config

	if err := decoder.Decode(&config); err != nil {
		send.Json(map[string]string{"error": "Invalid data"}, res, http.StatusBadRequest)
		return
	}

	team1, err := team.Repository().FindById(config.Teams.Team1)
	if err != nil {
		send.Json(map[string]string{"error": "Invalid data"}, res, http.StatusBadRequest)
		return
	}

	team2, err := team.Repository().FindById(config.Teams.Team2)
	if err != nil {
		send.Json(map[string]string{"error": "Invalid data"}, res, http.StatusBadRequest)
		return
	}

	document := Game{
		Championship: config.Championship,
		Teams: Teams{
			Team1: team1.Entity,
			Team2: team2.Entity,
		},
		State: State{},
	}

	newDocument, err := controller.Repository.InsertOne(document)

	if err != nil {
		send.Json(err, res, http.StatusBadRequest)
		return
	}

	send.Json(newDocument, res, http.StatusCreated)
}

var controller = base.Controller[Game]{
	Repository:   Repository(),
	CreateFilter: createFilter,
	ValidateData: ValidateData,
	Routes:       []string{"get", "getId", "delete"},
}
