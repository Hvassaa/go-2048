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
	// the dimensions for the grid / map
	height := 6
	width := 4

	// create the grid
	var grid [][]int = createGrid(height, width)

	// the game loop
	gameLoop(grid)
}

// the main game loop which handles input and display
func gameLoop(grid [][]int) {
	// get dimensions
	height := len(grid)
	width := len(grid[0])

	// seed random with the current time
	rand.Seed(time.Now().UTC().UnixNano())

	// create a reader
	reader := bufio.NewReader(os.Stdin)
	for {
		// randomly add some new 1's every round
		for i := 0; i < width; i++ {
			y := rand.Intn(height)
			x := rand.Intn(width)
			if grid[y][x] == 0 {
				grid[y][x] = 1
			}
		}

		// clear the screen
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()

		// show the map
		printGrid(grid)

		// get movement from stdin and move accordinlgy
		fmt.Print("Move: ")
		move, _ := reader.ReadString('\n')
		// use vim keybinding for movement
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

// slightly overengineered method to move both up and down
func moveVertical(grid [][]int, up bool) {
	// generic values to bet set depending on direction
	var start, stop, move, reverseStop, reverseMove int

	// generic function to compare two values
	// for use in the boolean test in the for-loops
	var compare greaterOrLesserTest
	height := len(grid)
	width := len(grid[0])

	// set values depending movement direction
	if up {
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

// see moveVertical
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

// a type for use in the move methods, for generic comparations
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
