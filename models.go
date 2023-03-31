package main

type Player struct {
	Name   string
	Board  *Board
	Hits   int
	Misses int
}

type Ship struct {
	Name     string
	Size     int
	Location [][2]int
	Hits     [][2]int
}
