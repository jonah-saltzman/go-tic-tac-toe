package main

import (
	"sort"
	"testing"
)

func TestSortMoves(t *testing.T) {
	boards := ReadBoards()
	c := Counter{}
	moves := GenMoves(&boards[6], 2, &c)
	sorted := GenMoves(&boards[6], 2, &c)
	sort.Sort(sorted)
	if sorted[0].to == moves[0].to {
		t.Errorf("Did not sort")
	}
}
