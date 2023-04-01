package main

type Player struct {
	// Имя игрока
	Name string
	// Доска игрока
	Board  *Board
	Hits   int
	Misses int
}

type Ship struct {
	// Имя
	Name string
	// Длина кооробля
	Size int
	// Расположение
	Location [][2]int
	// Жизни кооробля
	Hits [][2]int
}
