package tetris

import (
	"math"
	"os"
	"time"

	"github.com/diegolikescode/go-tetris/internal/screen"
)

type DrawPiece func(*os.File, int, int, [][]screen.Cell) Piece

type InputAction string

const (
	Quit       InputAction = "q" // just for testing
	Left       InputAction = "a"
	Right      InputAction = "d"
	TurnPiece  InputAction = "w"
	SmashDown  InputAction = " "
	Accelerate InputAction = "s"
	NoAction   InputAction = ""
)

func ReadUserAction(p Piece, matrix [][]screen.Cell) bool {
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

func SelectPiece(drawFuncArray []DrawPiece, row, col int, matrix [][]screen.Cell) Piece {
	// r := rand.Intn(6)
	return drawFuncArray[1](os.Stdout, row, col, matrix)
}

func CurrentMovingPiece(p Piece, row, col int, matrix [][]screen.Cell) {
	keep := true
	counter := 0
	for keep {
		lowPoint64 := 0.0
		for _, pos := range p.Position {
			lowPoint64 = math.Max(lowPoint64, float64(pos.Row))
		}

		lowestPoint := int(lowPoint64)
		// if matrix[lowestPoint+2][col].Occupied && matrix[lowestPoint][col].Type != screen.PermeableInterface {
		if matrix[lowestPoint+1][col].Occupied {
			keep = false
			break
		}
		time.Sleep(400 * time.Millisecond)

		screen.DrawMainBlock(os.Stdout, screen.MiddleRowMain, screen.MiddleColMain, matrix)
		p.MoveDown(matrix)

		counter++
	}
}

func MainLoop(matrix [][]screen.Cell) {
	keep := true

	firstRow := screen.StartMainRow
	firstCol := screen.StartMainCol

	funcArray := []DrawPiece{
		DrawTee, DrawL, DrawLInverted, DrawZ, DrawZInverted, DrawSquare, DrawColumn,
	}

	currPiece := SelectPiece(funcArray, firstRow, firstCol, matrix)
	go func() { CurrentMovingPiece(currPiece, firstRow, firstCol, matrix) }()

	for keep {
		keep = ReadUserAction(currPiece, matrix)
		if !keep {
			break
		}
	}
}
