package main

import (
	"math/rand"
)

func main() {
	// the dimensions for the grid / map
	height := 4
	width := 4

	// create the grid
	var grid [][]int = createGrid(height, width)
	createNewTiles(grid)

	//var displayerController DisplayerControllerInterface = cliDisplayer{}
	var displayerController DisplayerControllerInterface = gtkDisplayerController{}

	// the game loop
	displayerController.init()
	displayerController.startGameLoop(grid)
}

type DisplayerControllerInterface interface {
	init()
	startGameLoop(grid [][]int)
}

// randomly add some new 1's to the grid
func createNewTiles(grid [][]int) {
	// get dimensions
	height := len(grid)
	width := len(grid[0])

	for i := 0; i < width; i++ {
		y := rand.Intn(height)
		x := rand.Intn(width)
		if grid[y][x] == 0 {
			grid[y][x] = 1
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
	createNewTiles(grid)
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
	createNewTiles(grid)
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
