package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	boards := ReadBoards()
	board := boards[0]
	var scores = []Score{}
	alpha := -Max * 10
	beta := Max * 10
	c := Counter{}
	initialMoves := GenMoves(&board, 2, &c)
	SortMoves(&initialMoves)
	for _, move := range initialMoves {
		score := Score{to: move.to}
		score.score = alphabeta(*move.board, 0, alpha, beta, false, &c)
		//alpha = MinInt(alpha, score.score)
		scores = append(scores, score)
	}
	var bestScore Score
	bestScore.score = -Max
	for _, score := range scores {
		fmt.Printf("Move to %d: score = %d\n", score.to, score.score)
		if score.score > bestScore.score {
			bestScore = score
		}
	}
	fmt.Printf("Best move to: %v, score = %v\n", bestScore.to, bestScore.score)
	fmt.Printf("Total moves: %d\n", c.count)
}

func alphabeta(b board, depth int, alpha int, beta int, isMax bool, c *Counter) int {
	score, over := Over(&b)
	if over || depth >= 10 {
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
				best = MaxInt(best, alphabeta(b, depth+1, alpha, best, false, c))
				b[i] = 0
				alpha = MaxInt(alpha, best)
				if best >= beta {
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
				best = MinInt(best, alphabeta(b, depth+1, alpha, beta, true, c))
				b[i] = 0
				beta = MinInt(beta, best)
				if best <= alpha {
					break
				}
			}
		}
		return best
	}
}

func ReadBoards() []board {
	var path string = "board.json"
	jsonFile, err := os.Open(path)
	if err != nil {
		panic("Error loading json")
	}
	defer jsonFile.Close()
	bytes, _ := ioutil.ReadAll(jsonFile)
	var boards []board
	json.Unmarshal(bytes, &boards)
	return boards
}