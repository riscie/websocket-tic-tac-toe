package main

import (
	"fmt"
	"testing"
)

func TestCheckForWin(t *testing.T) {

	var winPatterns []gameState

	// Checking winning patterns
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

	// Checking non-win patterns
	var nonWinPatterns []gameState

	// Different non-winning patterns
	gs.Fields = []field{x, {}, x, {}, {}, {}, {}, {}, x}
	nonWinPatterns = append(nonWinPatterns, gs)

	gs.Fields = []field{{}, {}, {}, x, x, {}, {}, {}, x}
	nonWinPatterns = append(nonWinPatterns, gs)

	gs.Fields = []field{x, {}, {}, {}, {}, {}, {}, x, x}
	nonWinPatterns = append(nonWinPatterns, gs)

	gs.Fields = []field{x, {}, {}, x, {}, {}, {}, {}, x}
	nonWinPatterns = append(nonWinPatterns, gs)

	fmt.Printf("Testing %v winning patterns\n", len(winPatterns))

	for _, p := range winPatterns {
		if w, _ := p.checkForWin(); !w {
			t.Error("No Player detected as winning")
		}
	}
	fmt.Printf("Testing %v non-winning patterns\n", len(nonWinPatterns))

	for _, p := range nonWinPatterns {
		if w, _ := p.checkForWin(); w {
			t.Error("Detected win which is non-win")
		}
	}
}
