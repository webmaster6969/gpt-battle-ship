package main

import (
	"strconv"
	"strings"
)

func main() {
	Play("c1", "c2", 10, []int{2})
}

/*func main() {
	b := NewBoard(10, []int{1, 2, 3, 4})

	reader := bufio.NewReader(os.Stdin)

	for b.HasShipsLeft() {
		b.Print()

		fmt.Println("Enter coordinate (e.g. A5): ")
		text, _ := reader.ReadString('\n')

		x, y := parseCoordinate(text)
		hit := b.Shoot(x, y)

		if hit {
			fmt.Println("Hit!")
		} else {
			fmt.Println("Miss!")
		}
	}

	fmt.Println("Game over!")
}*/

func parseCoordinate(text string) (int, int) {
	text = strings.TrimSpace(text)
	if len(text) < 2 {
		return -1, -1
	}

	x := int(text[0] - 'A')
	y, err := strconv.Atoi(text[1:])
	if err != nil {
		return -1, -1
	}
	y--

	return x, y
}
