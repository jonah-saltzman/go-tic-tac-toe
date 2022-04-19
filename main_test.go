package main

import (
	"fmt"
	"testing"
)

func TestDriver(t *testing.T) {
	boards := ReadBoards()
	b0 := Driver(boards[0], 8)
	//b1 := Driver(boards[1])
	b2 := Driver(boards[2], 8)
	b3 := Driver(boards[3], 8)
	b4 := Driver(boards[4], 8)
	b5 := Driver(boards[5], 8)
	if b0 != 10 {
		fmt.Println("board 0 error")
	}
	if b2 != 22 {
		fmt.Println("board 2 error")
	}
	if b3 != 14 {
		fmt.Println("board 3 error")
	}
	if b4 != 12 {
		fmt.Println("board 4 error")
	}
	if b5 != 12 {
		fmt.Println("board 5 error")
	}
}

func BenchmarkMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main()
	}
}
