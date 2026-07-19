package player

import "souark/api/player/enum"

type Player struct {
	FirstName string        `json:"firstName"`
	LastName  string        `json:"lastName"`
	Division  enum.Division `json:"division"`
}
