package team

import "souark/api/player"

type Team struct {
	Player1 player.Player `json:"player1"`
	Player2 player.Player `json:"player2"`
}

type Config struct {
	Player1 string `json:"player1"`
	Player2 string `json:"player2"`
}
