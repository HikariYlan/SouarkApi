package game

import (
	"souark/api/game/enum"
	"souark/api/team"
)

type Teams struct {
	Team1 team.Team `json:"team1"`
	Team2 team.Team `json:"team2"`
}

type TeamsConfig struct {
	Team1 string `json:"team1"`
	Team2 string `json:"team2"`
}

type Score struct {
	Sets   int `json:"sets"`
	Points int `json:"points"`
	Fouls  int `json:"fouls"`
}

type State struct {
	CurrentSet int      `json:"currentSet"`
	Scores     [2]Score `json:"scores"`
	Winner     int      `json:"winner"`
}

type Game struct {
	Championship enum.Championship `json:"championship"`
	Teams        Teams             `json:"teams"`
	State        State             `json:"state"`
}

type Config struct {
	Championship enum.Championship `json:"championship"`
	Teams        TeamsConfig       `json:"teams"`
	State        State             `json:"state"`
}
