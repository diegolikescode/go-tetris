package main

import (
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

	tetris.DrawTee(os.Stdout, row, col, ScreenMatrix)
	tetris.DrawL(os.Stdout, row+5, col, ScreenMatrix)
	tetris.DrawLInverted(os.Stdout, row-10, col-10, ScreenMatrix)
	tetris.DrawZ(os.Stdout, row+10, col+10, ScreenMatrix)
	tetris.DrawZInverted(os.Stdout, 10, 10, ScreenMatrix)
	tetris.DrawSquare(os.Stdout, row-20, col-20, ScreenMatrix)
	tetris.DrawColumn(os.Stdout, col-10, col-10, ScreenMatrix)
	screen.MoveCursor(0, 1)
	os.Stdout.Sync()

	buf := make([]byte, 1)
	os.Stdin.Read(buf) // this thing makes it read next user input, and buf will store it
	restoreCb := screen.SetupRestore()
	restoreCb()
}
