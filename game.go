package main

import (
	"encoding/json"
	"log"
	"fmt"
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
		playerSymbols: map[int]string{
			0: "X",
			1: "O",
		},
	}
	return gs
}

type gameState struct {
	NumberOfPlayers int     `json:"numPlayers"`
	StatusMessage   string  `json:"statusMessage"`
	Fields          []field `json:"fields"`
	playerSymbols	map[int]string
}

type field struct {
	Set    bool   `json:"set"`
	Symbol string `json:"symbol"`
}

func (gs *gameState) makeMove(playerNum int, field int) {
	gs.Fields[field].Set = true
	gs.Fields[field].Symbol = gs.playerSymbols[playerNum] //X atm
	fmt.Println("Move: Player",playerNum,"Field",field)
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
