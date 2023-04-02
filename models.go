package main

// Location Позиции кораблей
type Location struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Player struct {
	// Имя игрока
	Name string `json:"name"`
	// Доска игрока
	Board  *Board `json:"board"`
	Hits   int    `json:"hits"`
	Misses int    `json:"misses"`
}

type Ship struct {
	// Длина корабля
	Size int `json:"size"`
	// Расположение
	Location []Location `json:"location"`
	// Жизни корабля
	Hits int `json:"hits"`
}
