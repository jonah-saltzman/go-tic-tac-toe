package main

import "fmt"

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

type Result struct {
	over  bool
	score int
}

func (m *Move) print() string {
	b := m.board
	return fmt.Sprintf("[\n%s%s%s%s%s]\n", pl(b[0:5]), pl(b[5:10]), pl(b[10:15]), pl(b[15:20]), pl(b[20:25]))
}

func (b *board) print() string {
	return fmt.Sprintf("[\n%s%s%s%s%s]\n\n", pl(b[0:5]), pl(b[5:10]), pl(b[10:15]), pl(b[15:20]), pl(b[20:25]))
}

func pl(line []uint) string {
	return fmt.Sprintf("   % 2v % 2v % 2v % 2v % 2v\n", line[0], line[1], line[2], line[3], line[4])
}

type Moves []Move

func (a *Moves) print() {
	for _, m := range *a {
		m.print()
		fmt.Printf("to: %d\n", m.to)
	}
}

func (c *Counter) inc() {
	c.count += 1
}

const Max int = 1000
const bLen int = 25
const wLen int = 5
