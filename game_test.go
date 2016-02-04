package main

import (
	"fmt"
	"testing"
)

func TestCheckForWin(t *testing.T) {

	var winPatterns []gameState

	gs := newGameState()
	x := field{Set: true, Symbol: "X"}

	// Different winning patterns
	// top row
	gs.Fields = []field{x, x, x, {}, {}, {}, {}, {}, {}}
	winPatterns = append(winPatterns, gs)
	// second row
	gs.Fields = []field{{}, {}, {}, x, x, x, {}, {}, {}}
	winPatterns = append(winPatterns, gs)
	// third row
	gs.Fields = []field{{}, {}, {}, {}, {}, {}, x, x, x}
	winPatterns = append(winPatterns, gs)
	// first column
	gs.Fields = []field{x, {}, {}, x, {}, {}, x, {}, {}}
	winPatterns = append(winPatterns, gs)
	// second column
	gs.Fields = []field{{}, x, {}, {}, x, {}, {}, x, {}}
	winPatterns = append(winPatterns, gs)
	// third column
	gs.Fields = []field{{}, {}, x, {}, {}, x, {}, {}, x}
	winPatterns = append(winPatterns, gs)
	// diagonal 1
	gs.Fields = []field{x, {}, {}, {}, x, {}, {}, {}, x}
	winPatterns = append(winPatterns, gs)
	// diagonal 2
	gs.Fields = []field{{}, {}, x, {}, x, {}, x, {}, {}}
	winPatterns = append(winPatterns, gs)

	fmt.Printf("Testing %v winning patterns: ", len(winPatterns))

	for _, p := range winPatterns {
		if w, _ := p.checkForWin(); !w {
			t.Error("no Player detected as winning")
		}
	}

}
