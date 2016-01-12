package main

var state gameState

type gameState struct {
	NumberOfPlayers	int 	`json:"numPlayers"`
	SomeSlice []string	`json:"someSlice"`
}

func (gs *gameState) AddPlayer() {
	gs.NumberOfPlayers++
}

func InitializeGameState(){
	state = gameState{
		NumberOfPlayers: 0,
		SomeSlice: []string{"These", "Are", "Some", "Tests"},
	}
}