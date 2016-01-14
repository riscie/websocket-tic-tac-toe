package main

import (
	"encoding/json"
	"log"
)

var playerSymbols = map[int]string{0: "X", 1: "O"}

func newGameState() gameState {
	gs := gameState{
		StatusMessage: "waiting for second player...",
		Fields: []field{
			field{Set: false, Symbol: ""},
			field{Set: false, Symbol: ""},
			field{Set: false, Symbol: ""},
			field{Set: false, Symbol: ""},
			field{Set: false, Symbol: ""},
			field{Set: false, Symbol: ""},
			field{Set: false, Symbol: ""},
			field{Set: true, Symbol: "O"},
			field{Set: false, Symbol: ""},
		},
	}
	return gs
}

type gameState struct {
	StatusMessage   string  `json:"statusMessage"`
	Fields          []field `json:"fields"`
	numberOfPlayers int     `json:"numberOfPlayers"`
}

type field struct {
	Set    bool   `json:"set"`
	Symbol string `json:"symbol"`
}

func (gs *gameState) addPlayer() {
	gs.numberOfPlayers++
	switch gs.numberOfPlayers {
	case 1:
		gs.StatusMessage = "waiting for second player..."
	case 2:
		gs.StatusMessage = "game begins!"
	}
}

func (gs *gameState) makeMove(playerNum int, field int) {
	gs.Fields[field].Set = true
	gs.Fields[field].Symbol = playerSymbols[playerNum] //X atm
}

func stateToJSON(gs gameState) []byte {
	json, err := json.Marshal(gs)
	if err != nil {
		log.Fatal("Error in marshalling json:", err)
	}
	return json
}
