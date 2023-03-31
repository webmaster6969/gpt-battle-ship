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
	for i, count := range shipCounts {
		for j := 0; j < count; j++ {
			size := i + 1
			ship := b.placeShip(size)
			b.Ships = append(b.Ships, ship)
		}
	}

	return b
}

// Пытаемся поставить корабль в нужную точку
func (b *Board) placeShip(size int) *Ship {
	for {
		ship := &Ship{
			Size: size,
		}

		x := rand.Intn(b.Size)
		y := rand.Intn(b.Size)
		vertical := rand.Intn(2) == 0

		if b.canPlaceShip(ship, x, y, vertical) {
			for i := 0; i < size; i++ {
				if vertical {
					ship.Location = append(ship.Location, [2]int{x + i, y})
				} else {
					ship.Location = append(ship.Location, [2]int{x, y + i})
				}
			}
			return ship
		}
	}
}

func (b *Board) canPlaceShip(ship *Ship, x, y int, vertical bool) bool {
	for i := 0; i < ship.Size; i++ {
		if x+i >= b.Size || y+i >= b.Size {
			return false
		}

		if vertical {
			if b.Grid[x+i][y] != "." {
				return false
			}
		} else {
			if b.Grid[x][y+i] != "." {
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

func (b *Board) sinkShip(ship *Ship) {
	for _, loc := range ship.Location {
		b.Grid[loc[0]][loc[1]] = "#"
	}
}

func (b *Board) HasShipsLeft() bool {
	for _, ship := range b.Ships {
		if len(ship.Hits) < ship.Size {
			return true
		}
	}
	return false
}

func (b *Board) Print() {
	fmt.Print("  ")
	for i := 0; i < b.Size; i++ {
		fmt.Printf("%c ", 'A'+i)
	}
	fmt.Println()

	for i := 0; i < b.Size; i++ {
		fmt.Printf("%d ", i)
		for j := 0; j < b.Size; j++ {
			fmt.Printf("%s ", b.Grid[i][j])
		}
		fmt.Println()
	}
}
