package main

func main() {
	Play("c1", "c2", 10, []int{4, 3, 3, 2, 2, 2, 1, 1, 1, 1})
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
