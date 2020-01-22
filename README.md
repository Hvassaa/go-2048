# Go-2048

A small clone of the game 2048. I made this mostly to try out the Go programming language and the GTK toolkit.

Go-2048 might behave a bit different to the orignial, since it has been a while since I played the original and I did not look its' source code.

## Building

### Prerequisite

For Ubuntu 19.10:

    golang-go
    golang-github-gotk3-gotk3-dev

### Compiling 

In the folder simply run 

    go build

## Running

To run the build binary, simply run it by

    ./go-2048

Or without building 

   go run 2048.go cliDisplayer.go gtkDisplayerController.go

The program can optionally take either 1 or 3 parameters. The first parameter is how to display the game, either in STDOUT or as a GTK3 window (standard is gtk):

    ./go-2048 cli
    ./go-2048 gtk

You can also add the wanted height and width of the map (standard is 4 by 4):

    ./go-2048 gtk 4 4

## Acknowledgement

The original version of 2048 can be found [here](https://github.com/gabrielecirulli/2048). 
