package screen

import (
	"fmt"
	"os"
	"os/signal"
)

type Cell struct {
	Occupied bool
	Type     CharType
}

type CharType string

const (
	Piece              CharType = "piece"
	Interface          CharType = "interface"
	PermeableInterface CharType = "permeableInterface"
	Config             CharType = "config"
	Clear              CharType = ""
)

const (
	MAIN_BLOCK_WIDTH_OFFSET  = 10
	MAIN_BLOCK_HEIGHT_OFFSET = 10
)

// ANSI styles
const (
	bold    = "\x1b[1m"
	cyan    = "\x1b[36m"
	Reset   = "\x1b[0m" // \x1b[0m
	redPink = "\x1b[38;5;125m"
)

const (
	V = "│"
	H = "─"

	TL = "┌"
	BL = "└"
	TR = "┐"
	BR = "┘"
)

func ComposeANSI(str string, styles ...string) string {
	res := ""
	for _, stl := range styles {
		res += stl
	}
	res += str + Reset
	return res
}

func DrawMultiChars(stdout *os.File, charType CharType, str string, row int, col int, matrix [][]Cell) {
	for i := range str {
		DrawChar(stdout, charType, string(str[i]), row, col+i, matrix)
	}
}

func ClearScreen() {
	fmt.Fprint(os.Stdout, "\x1b[?25l") // hide cursor
	fmt.Fprint(os.Stdout, "\x1b[2J")   // clear screen

	fmt.Fprint(os.Stdout, "") // clear screen - again
}

func MoveCursor(row, col int) {
	fmt.Fprintf(os.Stdout, "\x1b[%d;%dH", row, col)
}

func DrawChar(stdout *os.File, charType CharType, char string, row int, col int, matrix [][]Cell) error {
	MoveCursor(row, col)
	fmt.Fprintf(stdout, char)

	occupied := true
	if charType == Clear {
		occupied = false
	}

	if matrix != nil {
		matrix[row][col] = Cell{Occupied: occupied, Type: charType}
	}

	return nil
}

// func DrawHoldPiece(piece)

func SetupRestore() func() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	cleanupDone := make(chan struct{}, 1)
	go func() {
		<-c
		restoreScreen()
		os.Exit(0)
	}()

	return func() {
		restoreScreen()
		close(cleanupDone)
	}
}

// fmt.Fprint(os.Stdout, "\x1b[?25h")
// fmt.Fprint(os.Stdout, "\x1b[?12h")
func restoreScreen() {
	fmt.Fprint(os.Stdout, "\x1b[0m")   // Reset attributes
	fmt.Fprint(os.Stdout, "\x1b[2J")   // clear screen
	fmt.Fprint(os.Stdout, "\x1b[H")    // move to home
	fmt.Fprint(os.Stdout, "\x1b[?25h") // show cursor
}
