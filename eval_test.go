package main

import (
	"fmt"
	"testing"
)

func TestOver(t *testing.T) {
	var score int
	var over bool
	score, over = Over(&notOver)
	if score != 0 {
		t.Errorf("notOver should not have a score")
	}
	if over {
		t.Errorf("notOver should not be over")
	}

	score, over = Over(&xWin)
	if score != -Max {
		t.Errorf("xWin score should be -1000")
	}
	if !over {
		t.Errorf("xWin should be over")
	}

	score, over = Over(&oWin)
	if score != Max {
		t.Errorf("oWin score should be 1000")
	}
	if !over {
		t.Errorf("oWin should be over")
	}

	score, over = Over(&draw)
	if score != 0 {
		t.Errorf("draw score should be 0")
	}
	if !over {
		t.Errorf("draw should be over")
	}

	score, over = Over(&xFullWin)
	if score != -Max {
		t.Errorf("xFullWin score should be -1000")
	}
	if !over {
		t.Errorf("xFullWin should be over")
	}

	score, over = Over(&oFullWin)
	if score != Max {
		t.Errorf("oFullWin score should be 1000")
	}
	if !over {
		t.Errorf("oFullWin should be over")
	}
	boards := ReadBoards()
	overs := overMap(boards, Over)
	for i, v := range overs {
		if v.score != 0 {
			t.Errorf(fmt.Sprintf("Over failed on board %v\n", i))
		}
	}
}

func TestEval(t *testing.T) {
	var score int
	score = Eval(&notOver)
	if score != 0 {
		t.Errorf("notOver should not have a score")
	}

	score = Eval(&xWin)
	if score != -Max {
		t.Errorf("xWin score should be -1000")
	}

	score = Eval(&oWin)
	if score != Max {
		t.Errorf("oWin score should be 1000")
	}

	score = Eval(&draw)
	if score != 0 {
		t.Errorf("draw score should be 0")
	}

	score = Eval(&xFullWin)
	if score != -Max {
		t.Errorf("xFullWin score should be -1000")
	}

	score = Eval(&oFullWin)
	if score != Max {
		t.Errorf("oFullWin score should be 1000")
	}
	boards := ReadBoards()
	evals := evalMap(boards, Eval)
	for i, v := range evals {
		if v != 0 {
			t.Errorf(fmt.Sprintf("Eval failed on board %v\n", i))
		}
	}
}

func evalMap(boards []board, f func(b *board) int) []int {
	vsm := make([]int, len(boards))
	for i, v := range boards {
		vsm[i] = f(&v)
	}
	return vsm
}

func overMap(boards []board, f func(b *board) (int, bool)) []Result {
	vsm := make([]Result, len(boards))
	for i, v := range boards {
		score, over := f(&v)
		r := Result{score: score, over: over}
		vsm[i] = r
	}
	return vsm
}

func BenchmarkOver(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, board := range allBoards {
			Over(&board)
		}
	}
}

func BenchmarkEval(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, board := range allBoards {
			Eval(&board)
		}
	}
}

var allBoards []board = []board{notOver, xWin, oWin, draw, oFullWin, xFullWin}

var notOver board = [25]uint{
	1, 2, 0, 0, 0,
	1, 0, 2, 0, 0,
	0, 0, 0, 0, 0,
	1, 0, 0, 0, 0,
	1, 0, 0, 0, 0,
}

var xWin board = [25]uint{
	1, 2, 0, 0, 0,
	1, 0, 2, 0, 0,
	1, 0, 0, 0, 0,
	1, 0, 0, 0, 0,
	1, 2, 2, 2, 2,
}

var oWin board = [25]uint{
	1, 2, 0, 0, 0,
	1, 2, 2, 1, 1,
	0, 2, 1, 0, 0,
	1, 2, 1, 0, 0,
	1, 2, 1, 0, 0,
}

var draw board = [25]uint{
	1, 2, 1, 1, 2,
	1, 1, 2, 2, 2,
	2, 2, 2, 2, 1,
	1, 1, 1, 1, 2,
	1, 1, 1, 1, 2,
}

var xFullWin board = [25]uint{
	1, 2, 1, 1, 1,
	1, 1, 2, 1, 2,
	2, 2, 1, 2, 1,
	1, 1, 1, 1, 2,
	1, 1, 1, 1, 2,
}

var oFullWin board = [25]uint{
	1, 2, 1, 1, 2,
	1, 1, 2, 2, 2,
	2, 2, 2, 2, 1,
	1, 1, 1, 1, 2,
	2, 2, 2, 2, 2,
}
