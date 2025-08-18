package screen

import (
	"fmt"
	"os"
	"os/signal"
)

type Cell struct {
	Occupied bool
	Type     string
}

func ClearScreen() {
	fmt.Fprint(os.Stdout, "\x1b[?25l") // hide cursor
	fmt.Fprint(os.Stdout, "\x1b[2J")   // clear screen

	fmt.Fprint(os.Stdout, "") // clear screen
}

func MoveCursor(row, col int) {
	fmt.Fprintf(os.Stdout, "\x1b[%d;%dH", row, col)
}

func DrawChar(stdout *os.File, char string, row, col int, matrix [][]Cell) error {
	MoveCursor(row, col)
	fmt.Fprintf(stdout, char)

	// matrix[row][col] =

	return nil
}

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
	fmt.Fprint(os.Stdout, "\x1b[0m")   // reset attributes
	fmt.Fprint(os.Stdout, "\x1b[2J")   // clear screen
	fmt.Fprint(os.Stdout, "\x1b[H")    // move to home
	fmt.Fprint(os.Stdout, "\x1b[?25h") // show cursor
}
