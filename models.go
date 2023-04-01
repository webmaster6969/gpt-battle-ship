package main

type Location struct {
	X int
	Y int
}

type Player struct {
	// Имя игрока
	Name string
	// Доска игрока
	Board  *Board
	Hits   int
	Misses int
}

type Ship struct {
	// Длина корабля
	Size int
	// Расположение
	Location []Location
	// Жизни корабля
	Hits int
}
