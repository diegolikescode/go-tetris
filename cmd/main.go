package main

import (
	"os"

	"github.com/diegolikescode/go-tetris/internal/screen"
	"github.com/diegolikescode/go-tetris/internal/tetris"
	"golang.org/x/term"
)

var ScreenMatrix = [][]int{}

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

	ScreenMatrix = make([][]int, h)
	for i := range ScreenMatrix {
		ScreenMatrix[i] = make([]int, w)
	}

	row := h / 2
	col := w / 2

	tetris.DrawTee(os.Stdout, row, col, ScreenMatrix)
	tetris.DrawL(os.Stdout, row+5, col, ScreenMatrix)
	tetris.DrawLInverted(os.Stdout, row-5, col-6, ScreenMatrix)
	tetris.DrawZ(os.Stdout, row+10, col+10, ScreenMatrix)
	tetris.DrawZInverted(os.Stdout, 5, 5, ScreenMatrix)
	tetris.DrawSquare(os.Stdout, row-20, col-20, ScreenMatrix)
	tetris.DrawColumn(os.Stdout, col-10, col-10, ScreenMatrix)
	screen.MoveCursor(0, 1)
	os.Stdout.Sync()

	buf := make([]byte, 1)
	os.Stdin.Read(buf) // this thing makes it read next user input, and buf will store it
	restoreCb := screen.SetupRestore()
	restoreCb()
}
