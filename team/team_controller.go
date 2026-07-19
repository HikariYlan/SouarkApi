package team

import (
	"encoding/json"
	"net/http"
	"souark/api/base"
	"souark/api/player"
	"souark/api/services/send"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func createFilter(req *http.Request) bson.D {

	playerOne := req.URL.Query().Get("playerOne")
	playerTwo := req.URL.Query().Get("playerTwo")

	var playerOneFilter bson.E
	var playerTwoFilter bson.E

	if playerOne != "" {
		playerOneFilter = bson.E{Key: "playerOne", Value: playerOne}
	}

	if playerTwo != "" {
		playerTwoFilter = bson.E{Key: "playerTwo", Value: playerTwo}
	}

	return bson.D{playerOneFilter, playerTwoFilter}
}

func ValidateData(req *http.Request) (bool, Team) {
	decoder := json.NewDecoder(req.Body)

	var document Team

	if err := decoder.Decode(&document); err != nil {
		return false, document
	}

	return true, document
}

func NewTeam(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	var config Config

	if err := decoder.Decode(&config); err != nil {
		send.Json(map[string]string{"error": "Invalid data"}, res, http.StatusBadRequest)
		return
	}

	player1, err := player.Repository().FindById(config.Player1)
	if err != nil {
		send.Json(map[string]string{"error": "Invalid data"}, res, http.StatusBadRequest)
		return
	}

	player2, err := player.Repository().FindById(config.Player2)
	if err != nil {
		send.Json(map[string]string{"error": "Invalid data"}, res, http.StatusBadRequest)
		return
	}

	document := Team{
		Player1: player1.Entity,
		Player2: player2.Entity,
	}

	newDocument, err := controller.Repository.InsertOne(document)

	if err != nil {
		send.Json(err, res, http.StatusBadRequest)
		return
	}

	send.Json(newDocument, res, http.StatusCreated)
}

var controller = base.Controller[Team]{
	Repository:   Repository(),
	CreateFilter: createFilter,
	ValidateData: ValidateData,
	Routes:       []string{"get", "getId", "delete"},
}
