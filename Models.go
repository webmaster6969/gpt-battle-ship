package main

// Location Позиции кораблей
type Location struct {
	X int `json:"X"`
	Y int `json:"Y"`
}

type Player struct {
	// Имя игрока
	Name string `json:"Name"`
	// Доска игрока
	Board  *Board `json:"Board"`
	Hits   int    `json:"Hits"`
	Misses int    `json:"Misses"`
}

type Ship struct {
	// Длина корабля
	Size int `json:"Size"`
	// Расположение
	Location []Location `json:"Location"`
	// Жизни корабля
	Hits int `json:"Hits"`
}
