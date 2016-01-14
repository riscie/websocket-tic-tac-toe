package main

import (
	"encoding/json"
	"log"
)

func newGameState() gameState {
	gs := gameState{
		NumberOfPlayers: 0,
		StatusMessage:   "waiting for second player...",
		Fields: []field{
			field{Set: false, Symbol: ""},
			field{Set: false, Symbol: ""},
			field{Set: false, Symbol: ""},
			field{Set: false, Symbol: ""},
			field{Set: false, Symbol: ""},
			field{Set: true, Symbol: "X"},
			field{Set: false, Symbol: ""},
			field{Set: true, Symbol: "O"},
			field{Set: false, Symbol: ""},
		},
	}
	return gs
}

type gameState struct {
	NumberOfPlayers int     `json:"numPlayers"`
	StatusMessage   string  `json:"statusMessage"`
	Fields          []field `json:"fields"`
}

type field struct {
	Set    bool   `json:"set"`
	Symbol string `json:"symbol"`
}

func (gs *gameState) AddPlayer() {
	gs.NumberOfPlayers++
}

func stateToJSON(gs gameState) []byte {
	json, err := json.Marshal(gs)
	if err != nil {
		log.Fatal("Error in marshalling json:", err)
	}
	return json
}
