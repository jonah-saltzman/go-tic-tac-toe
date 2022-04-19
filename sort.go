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
	// fmt.Println("=====================")
	// a[i].print()
	// fmt.Printf("iScore: %v\n\n", iScore)
	// a[j].print()
	// fmt.Printf("jScore: %v\n\n", jScore)
	return iScore > jScore
}
