package main

import (
	"github.com/gotk3/gotk3/gtk"
	"log"
	"strconv"
)

type gtkDisplayerController struct{}

func (g gtkDisplayerController) init() {}

func (g gtkDisplayerController) startGameLoop(grid [][]int) {
	// Initialize GTK without parsing any command line arguments.
	gtk.Init(nil)

	// Create a new toplevel window, set its title, and connect it to the
	// "destroy" signal to exit the GTK main loop when it is destroyed.
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle("Go-2048")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	// create a grid container
	gtkGrid, err := gtk.GridNew()
	if err != nil {
		log.Fatal("Unable to create grid layout:", err)
	}

	displayGrid(grid, *gtkGrid)

	// get dimensions
	height := len(grid)
	middleCell := int(len(grid[0]) / 2)

	// left
	leftButton, _ := gtk.ButtonNewWithLabel("Left")
	leftButton.Connect("clicked", func() {
		moveHorizontal(grid, true)
		displayGrid(grid, *gtkGrid)
		win.ShowAll()
	})
	gtkGrid.Attach(leftButton, middleCell-1, height+1, 1, 1)
	//right
	rightButton, _ := gtk.ButtonNewWithLabel("Right")
	rightButton.Connect("clicked", func() {
		moveHorizontal(grid, false)
		displayGrid(grid, *gtkGrid)
		win.ShowAll()
	})
	gtkGrid.Attach(rightButton, middleCell+1, height+1, 1, 1)
	//up
	upButton, _ := gtk.ButtonNewWithLabel("Up")
	upButton.Connect("clicked", func() {
		moveVertical(grid, true)
		displayGrid(grid, *gtkGrid)
		win.ShowAll()
	})
	gtkGrid.Attach(upButton, middleCell, height, 1, 1)
	//up
	downButton, _ := gtk.ButtonNewWithLabel("Down")
	downButton.Connect("clicked", func() {
		moveVertical(grid, false)
		displayGrid(grid, *gtkGrid)
		win.ShowAll()
	})
	gtkGrid.Attach(downButton, middleCell, height+2, 1, 1)

	gtkGrid.SetColumnHomogeneous(true)
	gtkGrid.SetRowHomogeneous(true)

	// Add the grid to the window.
	win.Add(gtkGrid)

	// Set the default window size.
	win.SetDefaultSize(800, 600)

	// Recursively show all widgets contained in this window.
	win.ShowAll()

	// Begin executing the GTK main loop.  This blocks until
	// gtk.MainQuit() is run.
	gtk.Main()
}

func displayGrid(grid [][]int, gtkGrid gtk.Grid) {
	// get dimensions
	height := len(grid)
	width := len(grid[0])

	// populate grid container with the grids values
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			current := grid[y][x]
			label, _ := gtk.LabelNew(strconv.Itoa(current))
			child, _ := gtkGrid.GetChildAt(x, y)
			child.Destroy()
			gtkGrid.Attach(label, x, y, 1, 1)
		}
	}
}
