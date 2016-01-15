package main

import (
	"encoding/json"
	"log"
)

// these constants represent different game status messages
const waitPaired = "Waiting to get paired"
const gameBegins = "Game begins!"
const draw = "Draw!"
const resetWaitPaired = "Opponent has been disconnected :( Waiting to get paired again"
const winner = " wins! Congratulations!"

// gameState is the struct which represents the gameState between two players
type gameState struct {
	//renaming json values here to confirm the standard (lowercase var names)
	StatusMessage string   `json:"statusMessage"`
	Fields        []field  `json:"fields"`
	PlayerSymbols []string `json:"playerSymbols"`
	Started       bool     `json:"started"`
	Over          bool     `json:"over"`
	//These are not exported to JSON
	numberOfPlayers int
	playersTurn     int
	numberOfMoves   int
}

// field represents one of the nine tic-tac-toe fields
type field struct {
	Set    bool   `json:"set"`
	Symbol string `json:"symbol"`
}

// newGameState is the constructor for the gameState struct and creates the initial gameState Struct (empty board)
func newGameState() gameState {
	gs := gameState{
		StatusMessage: waitPaired,
		Fields:        emptyFields(),
		PlayerSymbols: []string{0: "X", 1: "O"},
		Started:       false,
		//These are not exported to JSON
		numberOfPlayers: 0,
		playersTurn:     0,
	}
	return gs
}

func emptyFields() []field {
	return []field{
		field{}, field{}, field{}, // row1
		field{}, field{}, field{}, // row2
		field{}, field{}, field{}, // row3
	}
}

// addPlayer informs the gameState about the new player and alters the statusMessage
func (gs *gameState) addPlayer() {
	gs.numberOfPlayers++
	switch gs.numberOfPlayers {
	case 1:
		gs.StatusMessage = waitPaired
	case 2:
		gs.StatusMessage = gameBegins
		gs.Started = true
	}
}

// makeMove checks if it's the
func (gs *gameState) makeMove(playerNum int, moveNum int) {
	if moveNum <= 9 {
		if gs.isPlayersTurn(playerNum) {
			if gs.isLegalMove(moveNum) {
				gs.Fields[moveNum].Set = true
				gs.Fields[moveNum].Symbol = gs.PlayerSymbols[playerNum]
				gs.switchPlayersTurn(playerNum)
				gs.numberOfMoves++
				if won, symbol := gs.checkForWin(); won {
					gs.setWinner(symbol)
				} else {
					gs.checkForDraw()
				}
			}
		}
	} else {
		gs.specialMove(moveNum)
	}
}

// special move processes moves which are not board moves like restarting the game
func (gs *gameState) specialMove(moveNum int) {
	switch moveNum {
	//restart game
	case 10:
		gs.restartGame()
	}
}

// restartGame sets the gameState to a state so that a new game between the same
// players can begin
func (gs *gameState) restartGame() {
	gs.StatusMessage = gameBegins
	gs.Fields = emptyFields()
	gs.Over = false
	gs.numberOfMoves = 0

}

// resetGame is needed, when one player drops out. It sets the gameState to a state so that
// the player who is left can wait for a new opponent.
func (gs *gameState) resetGame() {
	gs.restartGame()
	gs.Started = false
	gs.StatusMessage = resetWaitPaired
}

// checkForWin checks if a player has three in a row
func (gs *gameState) checkForWin() (bool, string) {
	// TODO: There are way beter, more dynamic implementations for this. We could limit the search space
	// TODO: by looking at the field where the last move was made eg.
	//rows
	//check row1
	if gs.Fields[0].Symbol == gs.Fields[1].Symbol && gs.Fields[1].Symbol == gs.Fields[2].Symbol && gs.Fields[2].Symbol != "" {
		return true, gs.Fields[0].Symbol
	}
	//check row2
	if gs.Fields[3].Symbol == gs.Fields[4].Symbol && gs.Fields[4].Symbol == gs.Fields[5].Symbol && gs.Fields[5].Symbol != "" {
		return true, gs.Fields[3].Symbol
	}
	//check row2
	if gs.Fields[6].Symbol == gs.Fields[7].Symbol && gs.Fields[7].Symbol == gs.Fields[8].Symbol && gs.Fields[8].Symbol != "" {
		return true, gs.Fields[7].Symbol
	}

	//columns
	//column 1
	if gs.Fields[0].Symbol == gs.Fields[3].Symbol && gs.Fields[3].Symbol == gs.Fields[6].Symbol && gs.Fields[6].Symbol != "" {
		return true, gs.Fields[0].Symbol
	}
	//column 2
	if gs.Fields[1].Symbol == gs.Fields[4].Symbol && gs.Fields[4].Symbol == gs.Fields[7].Symbol && gs.Fields[7].Symbol != "" {
		return true, gs.Fields[1].Symbol
	}
	//column 3
	if gs.Fields[2].Symbol == gs.Fields[5].Symbol && gs.Fields[5].Symbol == gs.Fields[8].Symbol && gs.Fields[8].Symbol != "" {
		return true, gs.Fields[2].Symbol
	}

	//diagonals
	//diagonal1
	if gs.Fields[0].Symbol == gs.Fields[4].Symbol && gs.Fields[4].Symbol == gs.Fields[8].Symbol && gs.Fields[8].Symbol != "" {
		return true, gs.Fields[0].Symbol
	}
	//diagonal2
	if gs.Fields[2].Symbol == gs.Fields[4].Symbol && gs.Fields[4].Symbol == gs.Fields[6].Symbol && gs.Fields[6].Symbol != "" {
		return true, gs.Fields[2].Symbol
	}
	return false, ""
}

func (gs *gameState) setWinner(symbol string) {
	gs.StatusMessage = symbol + winner
	gs.Over = true
}

// checkForDraw checks for draws
func (gs *gameState) checkForDraw() {
	if gs.numberOfMoves == 9 {
		gs.StatusMessage = draw
		gs.Over = true
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
