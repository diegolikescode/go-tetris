package tetris

import (
	"os"

	"github.com/diegolikescode/go-tetris/internal/screen"
)

// REFERENCE: https://www.reddit.com/media?url=https%3A%2F%2Fexternal-preview.redd.it%2FrMJqW-zViDxoKTkODFO00VnEyKKWQiC0CnCuLvwbxc8.jpg%3Fauto%3Dwebp%26s%3D758c506e60811db7a78b107d53c61e1868f23221

// NOTE: every row has to occupy 2 cells because of things like line-height, which is unever

func DrawTee(stdout *os.File, row, col int, matrix [][]screen.Cell) {
	screen.DrawChar(stdout, "\x1b[48;5;129m ", row, col, matrix)
	screen.DrawChar(stdout, "\x1b[38;5;129m ", row, col+1, matrix)
	screen.DrawChar(stdout, "\x1b[38;5;129m ", row, col+2, matrix)
	screen.DrawChar(stdout, "\x1b[38;5;129m ", row, col+3, matrix)
	screen.DrawChar(stdout, "\x1b[38;5;129m ", row, col+4, matrix)
	screen.DrawChar(stdout, "\x1b[38;5;129m ", row, col+5, matrix)

	screen.DrawChar(stdout, "\x1b[38;5;129m ", row+1, col+2, matrix)
	screen.DrawChar(stdout, "\x1b[38;5;129m ", row+1, col+3, matrix)
}

func DrawL(stdout *os.File, row, col int, matrix [][]screen.Cell) {
	screen.DrawChar(stdout, "\x1b[48;5;208m ", row, col, matrix)
	screen.DrawChar(stdout, "\x1b[48;5;208m ", row+1, col, matrix)
	screen.DrawChar(stdout, "\x1b[48;5;208m ", row+2, col, matrix)
	screen.DrawChar(stdout, "\x1b[48;5;208m ", row+2, col+1, matrix)

	screen.DrawChar(stdout, "\x1b[48;5;208m ", row, col+1, matrix)
	screen.DrawChar(stdout, "\x1b[48;5;208m ", row+1, col+1, matrix)
	screen.DrawChar(stdout, "\x1b[48;5;208m ", row+2, col+1, matrix)
	screen.DrawChar(stdout, "\x1b[48;5;208m ", row+2, col+2, matrix)
	screen.DrawChar(stdout, "\x1b[48;5;208m ", row+2, col+3, matrix)
}

func DrawLInverted(stdout *os.File, row, col int, matrix [][]screen.Cell) {
	screen.DrawChar(stdout, "\x1b[48;5;27m ", row, col, matrix)
	screen.DrawChar(stdout, "\x1b[48;5;27m ", row+1, col, matrix)
	screen.DrawChar(stdout, "\x1b[48;5;27m ", row+2, col, matrix)
	screen.DrawChar(stdout, "\x1b[48;5;27m ", row+2, col-1, matrix)

	screen.DrawChar(stdout, "\x1b[48;5;27m ", row, col+1, matrix)
	screen.DrawChar(stdout, "\x1b[48;5;27m ", row+1, col+1, matrix)
	screen.DrawChar(stdout, "\x1b[48;5;27m ", row+2, col+1, matrix)
	screen.DrawChar(stdout, "\x1b[48;5;27m ", row+2, col-1, matrix)
	screen.DrawChar(stdout, "\x1b[48;5;27m ", row+2, col-2, matrix)
}

func DrawZ(stdout *os.File, row, col int, matrix [][]screen.Cell) {
	screen.DrawChar(stdout, "\x1b[101m ", row, col, matrix)
	screen.DrawChar(stdout, "\x1b[101m ", row, col+1, matrix)
	screen.DrawChar(stdout, "\x1b[101m ", row, col+2, matrix)
	screen.DrawChar(stdout, "\x1b[101m ", row, col+3, matrix)

	screen.DrawChar(stdout, "\x1b[101m ", row+1, col+2, matrix)
	screen.DrawChar(stdout, "\x1b[101m ", row+1, col+3, matrix)
	screen.DrawChar(stdout, "\x1b[101m ", row+1, col+4, matrix)
	screen.DrawChar(stdout, "\x1b[101m ", row+1, col+5, matrix)
}

func DrawZInverted(stdout *os.File, row, col int, matrix [][]screen.Cell) {
	screen.DrawChar(stdout, "\x1b[48;5;46m ", row, col, matrix)
	screen.DrawChar(stdout, "\x1b[48;5;46m ", row, col-1, matrix)
	screen.DrawChar(stdout, "\x1b[48;5;46m ", row, col-2, matrix)
	screen.DrawChar(stdout, "\x1b[48;5;46m ", row, col-3, matrix)

	screen.DrawChar(stdout, "\x1b[48;5;46m ", row+1, col-2, matrix)
	screen.DrawChar(stdout, "\x1b[48;5;46m ", row+1, col-3, matrix)
	screen.DrawChar(stdout, "\x1b[48;5;46m ", row+1, col-4, matrix)
	screen.DrawChar(stdout, "\x1b[48;5;46m ", row+1, col-5, matrix)
}

func DrawSquare(stdout *os.File, row, col int, matrix [][]screen.Cell) {
	screen.DrawChar(stdout, "\x1b[48;5;226m ", row, col, matrix)
	screen.DrawChar(stdout, "\x1b[48;5;226m ", row, col+1, matrix)
	screen.DrawChar(stdout, "\x1b[48;5;226m ", row, col+2, matrix)
	screen.DrawChar(stdout, "\x1b[48;5;226m ", row, col+3, matrix)
	screen.DrawChar(stdout, "\x1b[48;5;226m ", row+1, col, matrix)
	screen.DrawChar(stdout, "\x1b[48;5;226m ", row+1, col+1, matrix)
	screen.DrawChar(stdout, "\x1b[48;5;226m ", row+1, col+2, matrix)
	screen.DrawChar(stdout, "\x1b[48;5;226m ", row+1, col+3, matrix)
}

func DrawColumn(stdout *os.File, row, col int, matrix [][]screen.Cell) {
	screen.DrawChar(stdout, "\x1b[48;5;51m ", row, col, matrix)
	screen.DrawChar(stdout, "\x1b[48;5;51m ", row+1, col, matrix)
	screen.DrawChar(stdout, "\x1b[48;5;51m ", row+2, col, matrix)
	screen.DrawChar(stdout, "\x1b[48;5;51m ", row+3, col, matrix)

	screen.DrawChar(stdout, "\x1b[48;5;51m ", row, col+1, matrix)
	screen.DrawChar(stdout, "\x1b[48;5;51m ", row+1, col+1, matrix)
	screen.DrawChar(stdout, "\x1b[48;5;51m ", row+2, col+1, matrix)
	screen.DrawChar(stdout, "\x1b[48;5;51m ", row+3, col+1, matrix)
}
