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
	dimension := 4
	var grid [][]int = createGrid(dimension)

	reader := bufio.NewReader(os.Stdin)
	for {
		for i := 0; i < dimension; i++ {
			y := rand.Intn(dimension)
			x := rand.Intn(dimension)
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
	dimension := len(grid)
	if left {
		start = 0
		stop = dimension
		reverseStop = start
		move = 1
		compare = func(a, b int) bool {
			return a < b
		}
	} else {
		start = dimension - 1
		stop = -1
		reverseStop = start
		move = -1
		compare = func(a, b int) bool {
			return a > b
		}
	}
	reverseMove = move * (-1)

	for i := 0; i < dimension; i++ {
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
	dimension := len(grid)
	if left {
		start = 0
		stop = dimension
		reverseStop = start
		move = 1
		compare = func(a, b int) bool {
			return a < b
		}
	} else {
		start = dimension - 1
		stop = -1
		reverseStop = start
		move = -1
		compare = func(a, b int) bool {
			return a > b
		}
	}
	reverseMove = move * (-1)

	for i := 0; i < dimension; i++ {
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

func moveLeft(grid [][]int) {
	dimension := len(grid)
	for i := 0; i < dimension; i++ {
		for j := 1; j < dimension; j++ {
			// the currently examined cell's value
			current := grid[i][j]
			// boolean indicating if cells have been merged
			// since each cell can only merge once per move
			// (i.e. 1 | 1 | 1 | 1 <<becomes<< 2 | 2 | 0 | 0)
			isMerged := false

			for reverse := j; reverse > 0; reverse-- {
				if grid[i][reverse-1] == 0 {
					grid[i][reverse-1] = current
					grid[i][reverse] = 0
				} else if grid[i][reverse-1] == current && !isMerged {
					current = current * 2
					grid[i][reverse-1] = current
					grid[i][reverse] = 0
					isMerged = true
				} else {
					break
				}

			}
		}
	}
}

// for simply creating a grid, n times n in dimensions
func createGrid(n int) [][]int {
	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, n)
	}
	return grid
}

// print grid to stdout
func printGrid(grid [][]int) {
	// save dimension (height/width) of the grid
	dimension := len(grid)

	// make the top line
	makeHorizontalSeperator(dimension)

	// loop through the whole grid-matrix
	for y := 0; y < dimension; y++ {
		fmt.Print("|")
		for x := 0; x < dimension; x++ {
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
		makeHorizontalSeperator(dimension)
	}
}

// make the horizontal line
func makeHorizontalSeperator(dimension int) {
	fmt.Print("+")
	for i := 0; i < dimension; i++ {
		fmt.Print("-----------+")
	}
	fmt.Println()
}

// get number of digits in a integer
func getLengthOfInt(number int) int {
	return len(strconv.Itoa(number))
}
