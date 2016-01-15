package main

import (
	"encoding/json"
	"log"
)

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
		PlayerSymbols:  []string{0: "X", 1: "O"},
		Started: false,
	}
	return gs
}

type gameState struct {
	//renaming json values here to confirm the standard (lowercase var names)
	StatusMessage   string  	`json:"statusMessage"`
	Fields          []field 	`json:"fields"`
	numberOfPlayers int     	`json:"-"` //not exported to JSON
	PlayerSymbols	[]string	`json:"playerSymbols"`
	Started		bool		`json:"started"`
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
		gs.Started = true
	}
}

func (gs *gameState) makeMove(playerNum int, field int) {
	gs.Fields[field].Set = true
	gs.Fields[field].Symbol = gs.PlayerSymbols[playerNum] //X atm
}

func stateToJSON(gs gameState) []byte {
	json, err := json.Marshal(gs)
	if err != nil {
		log.Fatal("Error in marshalling json:", err)
	}
	return json
}
