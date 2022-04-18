package main

func GenMoves(b *board, player uint, c *Counter) []Move {
	var moves = []Move{}
	for i := range b {
		if b[i] == 0 {
			c.inc()
			move := Move{to: i}
			newBoard := *b
			newBoard[i] = player
			move.board = &newBoard
			moves = append(moves, move)
		}
	}
	return moves
}
