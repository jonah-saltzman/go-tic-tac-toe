package main

import (
	"fmt"
)

type board [bLen]uint
type win [wLen]uint

type Score struct {
	to    int
	score int
}

type Counter struct {
	count int
}

type Move struct {
	board *board
	to    int
}

func (c *Counter) inc() {
	c.count += 1
}

const Max int = 1000
const bLen int = 25
const wLen int = 5

func main() {
	var board board = [25]uint{
		1, 2, 0, 0, 0,
		1, 0, 2, 0, 0,
		0, 0, 0, 0, 0,
		1, 0, 0, 0, 0,
		1, 0, 0, 0, 0,
	}
	Eval(&board)
	var scores = []Score{}
	alpha := -Max * 10
	beta := Max * 10
	c := Counter{}
	initialMoves := GenMoves(&board, 2, &c)
	for _, move := range initialMoves {
		score := Score{to: move.to}
		score.score = alphabeta(*move.board, 0, alpha, beta, false, &c)
		//alpha = MinInt(alpha, score.score)
		scores = append(scores, score)
	}
	for _, score := range scores {
		fmt.Printf("Move to %d: score = %d\n", score.to, score.score)
	}
	fmt.Printf("Total moves: %d\n", c.count)
}

func alphabeta(b board, depth int, alpha int, beta int, isMax bool, c *Counter) int {
	score, over := Over(&b)

	if over {
		if score > 0 {
			return score - depth
		} else if score < 0 {
			return score + depth
		}
		return score
	}

	var player uint
	if isMax {
		player = 2
	} else {
		player = 1
	}
	if isMax {
		best := -Max
		for i := 0; i < bLen; i++ {
			if b[i] == 0 {
				c.inc()
				b[i] = player
				best = MaxInt(best, alphabeta(b, depth+1, alpha, best, !isMax, c))
				b[i] = 0
				alpha = MaxInt(alpha, best)
				if beta <= alpha {
					break
				}
			}
		}
		return best
	} else {
		best := Max
		for i := 0; i < bLen; i++ {
			if b[i] == 0 {
				c.inc()
				b[i] = player
				best = MinInt(best, alphabeta(b, depth+1, alpha, beta, !isMax, c))
				b[i] = 0
				beta = MinInt(beta, best)
				if beta <= alpha {
					break
				}
			}
		}
		return best
	}
}
