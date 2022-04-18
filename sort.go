package main

import (
	// "fmt"
	"sort"
)

func SortMoves(moves *Moves) {
	sort.Sort(moves)
}

func (a Moves) Len() int {
	return len(a)
}

func (a Moves) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Moves) Less(i, j int) bool {
	iScore, jScore := Eval(a[i].board), Eval(a[j].board)
	//fmt.Printf("i: %v; j: %v\n", iScore, jScore)
	return iScore > jScore
}
