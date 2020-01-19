package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	height := 6
	width := 4
	var grid [][]int = createGrid(height, width)

	reader := bufio.NewReader(os.Stdin)
	for {
		for i := 0; i < width; i++ {
			y := rand.Intn(height)
			x := rand.Intn(width)
			if grid[y][x] == 0 {
				grid[y][x] = 1
			}
		}

		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()
		printGrid(grid)
		fmt.Print("Move: ")
		move, _ := reader.ReadString('\n')

		switch move {
		case "j\n":
			moveVertical(grid, false)
		case "k\n":
			moveVertical(grid, true)
		case "h\n":
			moveHorizontal(grid, true)
		case "l\n":
			moveHorizontal(grid, false)
		}
	}
}

func moveVertical(grid [][]int, left bool) {
	var start, stop, move, reverseStop, reverseMove int
	var compare greaterOrLesserTest
	height := len(grid)
	width := len(grid[0])
	if left {
		start = 0
		stop = height
		reverseStop = start
		move = 1
		compare = func(a, b int) bool {
			return a < b
		}
	} else {
		start = height - 1
		stop = -1
		reverseStop = start
		move = -1
		compare = func(a, b int) bool {
			return a > b
		}
	}
	reverseMove = move * (-1)

	for i := 0; i < width; i++ {
		for j := start; compare(j, stop); j = j + move {
			// the currently examined cell's value
			current := grid[j][i]

			for reverse := j; compare(reverseStop, reverse); reverse = reverse + reverseMove {
				if grid[reverse+reverseMove][i] == 0 {
					grid[reverse+reverseMove][i] = current
					grid[reverse][i] = 0
				} else if grid[reverse+reverseMove][i] == current {
					current = current * 2
					grid[reverse+reverseMove][i] = current
					grid[reverse][i] = 0
					break
				} else {
					break
				}

			}
		}
	}
}

func moveHorizontal(grid [][]int, left bool) {
	var start, stop, move, reverseStop, reverseMove int
	var compare greaterOrLesserTest
	height := len(grid)
	width := len(grid[0])
	if left {
		start = 0
		stop = width
		reverseStop = start
		move = 1
		compare = func(a, b int) bool {
			return a < b
		}
	} else {
		start = width - 1
		stop = -1
		reverseStop = start
		move = -1
		compare = func(a, b int) bool {
			return a > b
		}
	}
	reverseMove = move * (-1)

	for i := 0; i < height; i++ {
		for j := start; compare(j, stop); j = j + move {
			// the currently examined cell's value
			current := grid[i][j]

			for reverse := j; compare(reverseStop, reverse); reverse = reverse + reverseMove {
				if grid[i][reverse+reverseMove] == 0 {
					grid[i][reverse+reverseMove] = current
					grid[i][reverse] = 0
				} else if grid[i][reverse+reverseMove] == current {
					current = current * 2
					grid[i][reverse+reverseMove] = current
					grid[i][reverse] = 0
					break
				} else {
					break
				}

			}
		}
	}
}

type greaterOrLesserTest func(a, b int) bool

// for simply creating a grid, n times n in dimensions
func createGrid(height, width int) [][]int {
	grid := make([][]int, height)
	for i := 0; i < height; i++ {
		grid[i] = make([]int, width)
	}
	return grid
}

// print grid to stdout
func printGrid(grid [][]int) {
	// save dimension (height/width) of the grid
	height := len(grid)
	width := len(grid[0])

	// make the top line
	makeHorizontalSeperator(width)

	// loop through the whole grid-matrix
	for y := 0; y < height; y++ {
		fmt.Print("|")
		for x := 0; x < width; x++ {
			// get the value of the current cell
			value := grid[y][x]
			// get the number of digits in the value
			digits := getLengthOfInt(value)

			// print whitespace to add padding to cells
			for q := 0; q < 11-digits; q++ {
				fmt.Print(" ")
			}

			// print the actual value
			fmt.Print(strconv.Itoa(value))
			fmt.Print("|")
		}

		// make a newline
		fmt.Println()
		// make a seperator between the lines
		makeHorizontalSeperator(width)
	}
}

// make the horizontal line
func makeHorizontalSeperator(width int) {
	fmt.Print("+")
	for i := 0; i < width; i++ {
		fmt.Print("-----------+")
	}
	fmt.Println()
}

// get number of digits in a integer
func getLengthOfInt(number int) int {
	return len(strconv.Itoa(number))
}
