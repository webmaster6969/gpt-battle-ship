package main

import (
	"fmt"
	"math/rand"
	"time"
)

const BoardCellTypeEmpty = 0
const BoardCellTypeMiss = 1
const BoardCellTypeHit = 2
const BoardCellTypeDeath = 3

type Board struct {
	Size  int     `json:"Size"`
	Ships []*Ship `json:"ships"`
	Grid  [][]int `json:"grid"`
}

// NewBoard Создаем доску
func NewBoard(size int, shipCounts []int) *Board {
	rand.Seed(time.Now().UnixNano())

	// Создаем доску
	b := &Board{
		Size: size,
		Grid: make([][]int, size),
	}

	// Ставим каждой клетки изначальное значение
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			b.Grid[x] = make([]int, size)
			b.Grid[x][y] = BoardCellTypeEmpty
		}
	}

	// Ставим корабли на поле боя
	for _, count := range shipCounts {
		ship := b.placeShip(count)
		b.Ships = append(b.Ships, ship)
	}

	return b
}

// Пытаемся поставить корабль в нужную точку
func (b *Board) placeShip(size int) *Ship {
	for {
		ship := &Ship{
			Size: size,
		}

		// Выбираем случайное положение
		x := rand.Intn(b.Size)
		y := rand.Intn(b.Size)

		// Так же выбираем случайное расположение по вертикале или горизонтале
		vertical := rand.Intn(2) == 0

		// Проверям что точно можем рассположить корабль
		if b.canPlaceShip(ship, x, y, vertical) {
			for i := 0; i < size; i++ {
				if vertical {
					ship.Location = append(ship.Location, Location{x + i, y})
					//b.Grid[x+i][y] = "*"
				} else {
					//b.Grid[x][y+i] = "*"
					ship.Location = append(ship.Location, Location{x, y + i})
				}
			}
			return ship
		} else {
			return b.placeShip(size)
		}
	}
}

// Проверяем, свободна ли точка
func (b *Board) canPlaceShip(ship *Ship, x, y int, vertical bool) bool {

	// Проверка на выход за пределы поля
	if x+ship.Size >= b.Size || y+ship.Size >= b.Size {
		return false
	}

	// Проверяем каждую тотчку на соответсвие
	for i := 0; i < ship.Size; i++ {
		if vertical {
			if !b.isPointCell(x+i, y) {
				return false
			}
		} else {
			if !b.isPointCell(x, y+i) {
				return false
			}
		}
	}

	return true
}

// Сканируем доску для расмещения коробля
func (b *Board) isPointCell(_x, _y int) bool {
	if _x >= b.Size || _y >= b.Size || _x < 0 || _y < 0 {
		return false
	}

	// Создаем объект для проверки
	check := make([][]int, b.Size)
	for i := range check {
		check[i] = make([]int, b.Size)
	}

	// Создаем поле для проверки
	leftX := _x - 1
	centerX := _x
	rightX := _x + 1

	topY := _y - 1
	centerY := _y
	downY := _y + 1

	if leftX < 0 {
		leftX = 0
	}
	if topY < 0 {
		topY = 0
	}

	if rightX > b.Size {
		rightX = b.Size
	}

	if downY > b.Size {
		downY = b.Size
	}

	// Верх
	check[leftX][topY] = 1
	check[centerX][topY] = 1
	check[rightX][topY] = 1
	// Середина
	check[leftX][centerY] = 1
	check[centerX][centerY] = 1
	check[rightX][centerY] = 1
	// Низ
	check[leftX][downY] = 1
	check[centerX][downY] = 1
	check[rightX][downY] = 1

	// Проверяем поле
	for y := 0; y < b.Size; y++ {
		for x := 0; x < b.Size; x++ {
			if check[x][y] == 1 && b.FindShip(x, y) != nil {
				return false
			}
		}
	}

	return true
}

// FindShip Поиск части корабля
func (b *Board) FindShip(x, y int) *Ship {
	for _, ship := range b.Ships {

		// Перебираем позиции
		for _, loc := range ship.Location {

			if loc.X == x && loc.Y == y {
				return ship
			}
		}
	}
	return nil
}

// Shoot Выстрел по доске
func (b *Board) Shoot(x, y int) bool {
	if b.Grid[x][y] == BoardCellTypeHit || b.Grid[x][y] == BoardCellTypeMiss {
		return false
	}

	ship := b.FindShip(x, y)

	if ship != nil {
		ship.Hits++
		b.Grid[x][y] = BoardCellTypeHit
		if ship.Hits == ship.Size {
			b.sinkShip(ship)
		}
		return true
	}

	b.Grid[x][y] = BoardCellTypeMiss
	return false
}

// Стив что корабль убит
func (b *Board) sinkShip(ship *Ship) {
	for _, loc := range ship.Location {
		b.Grid[loc.X][loc.Y] = BoardCellTypeDeath
	}
}

// HasShipsLeft Првоеряем есть ли еще живые корабли на поле
func (b *Board) HasShipsLeft() bool {
	for _, ship := range b.Ships {
		if ship.Hits < ship.Size {
			return true
		}
	}
	return false
}

// Print Рисуем доску
func (b *Board) Print() {
	fmt.Print("  ")
	for i := 0; i < b.Size; i++ {
		fmt.Printf("%c ", 'A'+i)
	}
	fmt.Println()

	for y := 0; y < b.Size; y++ {
		fmt.Printf("%d ", y+1)
		for x := 0; x < b.Size; x++ {
			switch b.Grid[x][y] {
			case BoardCellTypeEmpty:
				if b.FindShip(x, y) != nil {
					fmt.Printf("* ")
				} else {
					fmt.Printf(". ")
				}
				break
			case BoardCellTypeMiss:
				fmt.Printf("O ")
				break
			case BoardCellTypeHit:
				fmt.Printf("X ")
				break
			case BoardCellTypeDeath:
				fmt.Printf("# ")
				break
			}

		}
		fmt.Println()
	}
}
