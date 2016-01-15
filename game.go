package main

import (
	"encoding/json"
	"log"
)

// gameState is the struct which represents the gamestate between two players
type gameState struct {
	//renaming json values here to confirm the standard (lowercase var names)
	StatusMessage string   `json:"statusMessage"`
	Fields        []field  `json:"fields"`
	PlayerSymbols []string `json:"playerSymbols"`
	Started       bool     `json:"started"`
	//These are not exported to JSON
	numberOfPlayers int
	playersTurn     int
	numberOfMoves   int
}

// field represents one of the nine tic-tac-toe fileds
type field struct {
	Set    bool   `json:"set"`
	Symbol string `json:"symbol"`
}

// newGameState is the constructor for the gameState struct and creates the initial gameState Struct (emtpy board)
func newGameState() gameState {
	gs := gameState{
		StatusMessage: "Waiting to get paired",
		Fields: []field{
			field{}, field{}, field{}, // row1
			field{}, field{}, field{}, // row2
			field{}, field{}, field{}, // row3
		},
		PlayerSymbols: []string{0: "X", 1: "O"},
		Started:       false,
		//These are not exported to JSON
		numberOfPlayers: 0,
		playersTurn:     0,
	}
	return gs
}

// addPlayer informs the gamestate about the new player and alters the statusMessage
func (gs *gameState) addPlayer() {
	gs.numberOfPlayers++
	switch gs.numberOfPlayers {
	case 1:
		gs.StatusMessage = "Waiting to get paired"
	case 2:
		gs.StatusMessage = "Game begins!"
		gs.Started = true
	}
}

// makeMove checks if it's the
func (gs *gameState) makeMove(playerNum int, field int) {
	if gs.isPlayersTurn(playerNum) {
		if gs.isLegalMove(field) {
			gs.Fields[field].Set = true
			gs.Fields[field].Symbol = gs.PlayerSymbols[playerNum]
			gs.switchPlayersTurn(playerNum)
			gs.numberOfMoves++
			gs.checkForDraw()
		}
	}
}

// checkForDraw checks for draws
func (gs *gameState) checkForDraw() {
	//Todo: Implement
	if gs.numberOfMoves == 9 {
		gs.StatusMessage = "Draw!"
	}
}

// isLegalMove checks if a move is legal
func (gs *gameState) isLegalMove(field int) bool {
	return !gs.Fields[field].Set
}

// isPlayersTurn checks if it's the players turn
func (gs *gameState) isPlayersTurn(playerNum int) bool {
	return playerNum == gs.playersTurn
}

// switchPlayersTurn switches the playersTurn variable to the id of the other player
func (gs *gameState) switchPlayersTurn(playerNum int) {
	switch playerNum {
	case 0:
		gs.playersTurn = 1
	case 1:
		gs.playersTurn = 0
	}
}

// gameStateToJSON marshals the gameState struct to JSON represented by a slice of bytes
func (gs *gameState) gameStateToJSON() []byte {
	json, err := json.Marshal(gs)
	if err != nil {
		log.Fatal("Error in marshalling json:", err)
	}
	return json
}
