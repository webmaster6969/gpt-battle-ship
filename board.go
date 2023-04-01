package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Board struct {
	Size  int
	Ships []*Ship
	Grid  [][]string
}

// NewBoard Создаем доску
func NewBoard(size int, shipCounts []int) *Board {
	rand.Seed(time.Now().UnixNano())

	// Создаем доску
	b := &Board{
		Size: size,
		Grid: make([][]string, size),
	}

	// Ставим каждой клетки изначальное значение
	for i := 0; i < size; i++ {
		b.Grid[i] = make([]string, size)
		for j := 0; j < size; j++ {
			b.Grid[i][j] = "."
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
					ship.Location = append(ship.Location, [2]int{x + i, y})
					b.Grid[x+i][y] = "*"
				} else {
					b.Grid[x][y+i] = "*"
					ship.Location = append(ship.Location, [2]int{x, y + i})
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

func (b *Board) isPointCell(_x, _y int) bool {
	if _x >= b.Size || _y >= b.Size || _x < 0 || _y < 0 {
		return false
	}

	check := make([][]int, b.Size)
	for i := range check {
		check[i] = make([]int, b.Size)
	}

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

	for y := 0; y < b.Size; y++ {
		for x := 0; x < b.Size; x++ {
			//fmt.Printf("%s ", b.Grid[x][y])
			if check[x][y] == 1 && b.Grid[x][y] != "." {
				return false
			}
		}
	}

	return true
}

func (b *Board) Shoot(x, y int) bool {
	if b.Grid[x][y] == "X" || b.Grid[x][y] == "O" {
		return false
	}

	for _, ship := range b.Ships {
		for _, loc := range ship.Location {
			if loc[0] == x && loc[1] == y {
				ship.Hits = append(ship.Hits, loc)
				b.Grid[x][y] = "X"
				if len(ship.Hits) == ship.Size {
					b.sinkShip(ship)
				}
				return true
			}
		}
	}

	b.Grid[x][y] = "O"
	return false
}

// Стив что корабль убит
func (b *Board) sinkShip(ship *Ship) {
	for _, loc := range ship.Location {
		b.Grid[loc[0]][loc[1]] = "#"
	}
}

// HasShipsLeft Првоеряем есть ли еще живые корабли на поле
func (b *Board) HasShipsLeft() bool {
	for _, ship := range b.Ships {
		if len(ship.Hits) < ship.Size {
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
			fmt.Printf("%s ", b.Grid[x][y])
		}
		fmt.Println()
	}
}
