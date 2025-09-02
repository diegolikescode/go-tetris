package main

import (
	"fmt"
	"os"

	"github.com/diegolikescode/go-tetris/internal/screen"
	"github.com/diegolikescode/go-tetris/internal/tetris"
	"golang.org/x/term"
)

var ScreenMatrix = [][]screen.Cell{}

func main() {
	stdInFd := int(os.Stdin.Fd())
	stdOutFd := int(os.Stdout.Fd())

	oldState, err := term.MakeRaw(stdInFd)
	if err != nil {
		panic(err)
	}
	defer term.Restore(stdInFd, oldState)

	screen.ClearScreen()
	w, h, err := term.GetSize(stdOutFd)
	if err != nil || w <= 0 || h <= 0 {
		w, h = 80, 40
	}

	ScreenMatrix = make([][]screen.Cell, h)
	for i := range ScreenMatrix {
		ScreenMatrix[i] = make([]screen.Cell, w)
	}

	row := h / 2
	col := w / 2

	screen.DrawMainBlock(os.Stdout, row, col, ScreenMatrix)
	screen.DrawHoldBlock(os.Stdout, row, col, ScreenMatrix)
	screen.DrawCommandTexts(os.Stdout, row, col, ScreenMatrix)

	drawSize(h, w)
	screen.MoveCursor(0, 1)
	os.Stdout.Sync()

	tetris.MainLoop(ScreenMatrix)

	// var buf []byte
	// buf = make([]byte, 1)
	// os.Stdin.Read(buf)

	restoreCb := screen.SetupRestore()
	restoreCb()
}

func drawSize(h, w int) {
	screen.MoveCursor(0, 0)
	fmt.Fprint(os.Stdout, "size: ", w, "x", h)
}
