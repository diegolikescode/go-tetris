package tetris

import (
	"math/rand"
	"os"
	"time"

	"github.com/diegolikescode/go-tetris/internal/screen"
)

type DrawPiece func(*os.File, int, int, [][]screen.Cell) Piece

type InputAction string

const (
	Quit       InputAction = "q"
	Left       InputAction = "a"
	Right      InputAction = "d"
	TurnPiece  InputAction = "w"
	SmashDown  InputAction = " "
	Accelerate InputAction = "s"
	NoAction   InputAction = ""
)

func ReadUserAction(p *Piece, matrix [][]screen.Cell) bool {
	buf := make([]byte, 1)
	n, err := os.Stdin.Read(buf)
	if err != nil {
		os.Exit(1)
	}

	bufStr := string(buf[:n])
	screen.DrawChar(os.Stdout, screen.Interface, bufStr, 5, 5, matrix)

	switch bufStr {
	case "a", "A":
		p.MoveLeft(matrix)
		return true

	case "d", "D":
		p.MoveRight(matrix)
		return true

	case "s", "S":
		p.MoveDown(matrix)
		return true

	case " ":
		p.SmashDown(matrix)
		return true

	case "q", "Q":
		return false
	}

	return true
}

var (
	piecesSequence [1_000_000]int
	sequenceIndex  int
)

func prepSequence() {
	prev := -1
	for i := range piecesSequence {
		random := rand.Intn(7)
		for prev == random {
			random = rand.Intn(7)
		}

		prev = random
		piecesSequence[i] = random
	}
}

func SelectPiece(drawFuncArray []DrawPiece, row, col int, matrix [][]screen.Cell) Piece {
	sequenceIndex++
	return drawFuncArray[piecesSequence[sequenceIndex]](os.Stdout, row, col, matrix)
}

func CurrentMovingPiece(p *Piece, row, col int, matrix [][]screen.Cell) {
	keep := true
	counter := 0
	for keep {
		if !p.CanGoDown(matrix) {
			keep = false
			break
		}

		screen.DrawMainBlock(os.Stdout, screen.MiddleRowMain, screen.MiddleColMain, matrix)
		p.MoveDown(matrix)

		counter++
		time.Sleep(400 * time.Millisecond)
	}
}

func MainLoop(matrix [][]screen.Cell) {
	prepSequence()

	keep := true

	firstRow := screen.StartMainRow - 3
	firstCol := screen.StartMainCol - 1

	funcArray := []DrawPiece{
		DrawTee, DrawL, DrawLInverted, DrawZ, DrawZInverted, DrawSquare, DrawColumn,
	}

	var currPiece Piece
	go func() {
		for keep {
			currPiece = SelectPiece(funcArray, firstRow, firstCol, matrix)
			CurrentMovingPiece(&currPiece, firstRow, firstCol, matrix)

		}
	}()

	for keep {
		keep = ReadUserAction(&currPiece, matrix)
		if !keep {
			break
		}
	}
}
