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

type cliDisplayer struct{}

func (c cliDisplayer) init() {}

// the main game loop which handles input and display
func (c cliDisplayer) startGameLoop(grid [][]int) {
	// set the wanted displayer

	// seed random with the current time
	rand.Seed(time.Now().UTC().UnixNano())

	// create a reader
	reader := bufio.NewReader(os.Stdin)
	for {
		createNewTiles(grid)
		// clear the screen
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()

		// show the map
		displayGrid(grid)

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

// print grid to stdout
func displayGrid(grid [][]int) {
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
