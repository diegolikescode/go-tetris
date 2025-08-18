package tetris

import (
	"os"

	"github.com/diegolikescode/go-tetris/internal/screen"
)

// REFERENCE: https://www.reddit.com/media?url=https%3A%2F%2Fexternal-preview.redd.it%2FrMJqW-zViDxoKTkODFO00VnEyKKWQiC0CnCuLvwbxc8.jpg%3Fauto%3Dwebp%26s%3D758c506e60811db7a78b107d53c61e1868f23221

func DrawTee(stdout *os.File, row, col int, matrix [][]int) {
	screen.DrawChar(stdout, "\x1b[101m ", row, col, matrix)
	screen.DrawChar(stdout, "\x1b[101m ", row, col+1, matrix)
	screen.DrawChar(stdout, "\x1b[101m ", row, col+2, matrix)
	screen.DrawChar(stdout, "\x1b[101m ", row+1, col+1, matrix)
}

func DrawL(stdout *os.File, row, col int, matrix [][]int) {
	screen.DrawChar(stdout, "\x1b[101m ", row, col, matrix)
	screen.DrawChar(stdout, "\x1b[101m ", row+1, col, matrix)
	screen.DrawChar(stdout, "\x1b[101m ", row+2, col, matrix)
	screen.DrawChar(stdout, "\x1b[101m ", row+2, col+1, matrix)
}

func DrawLInverted(stdout *os.File, row, col int, matrix [][]int) {
	screen.DrawChar(stdout, "\x1b[101m ", row, col, matrix)
	screen.DrawChar(stdout, "\x1b[101m ", row+1, col, matrix)
	screen.DrawChar(stdout, "\x1b[101m ", row+2, col, matrix)
	screen.DrawChar(stdout, "\x1b[101m ", row+2, col-1, matrix)
}

func DrawZ(stdout *os.File, row, col int, matrix [][]int) {
	screen.DrawChar(stdout, "\x1b[101m ", row, col, matrix)
	screen.DrawChar(stdout, "\x1b[101m ", row, col+1, matrix)
	screen.DrawChar(stdout, "\x1b[101m ", row+1, col+1, matrix)
	screen.DrawChar(stdout, "\x1b[101m ", row+1, col+2, matrix)
}

func DrawZInverted(stdout *os.File, row, col int, matrix [][]int) {
	screen.DrawChar(stdout, "\x1b[101m ", row, col, matrix)
	screen.DrawChar(stdout, "\x1b[101m ", row, col-1, matrix)
	screen.DrawChar(stdout, "\x1b[101m ", row+1, col-1, matrix)
	screen.DrawChar(stdout, "\x1b[101m ", row+1, col-2, matrix)
}

func DrawSquare(stdout *os.File, row, col int, matrix [][]int) {
	screen.DrawChar(stdout, "\x1b[101m ", row, col, matrix)
	screen.DrawChar(stdout, "\x1b[101m ", row, col+1, matrix)
	screen.DrawChar(stdout, "\x1b[101m ", row, col+2, matrix)
	screen.DrawChar(stdout, "\x1b[101m ", row, col+3, matrix)
	screen.DrawChar(stdout, "\x1b[101m ", row+1, col, matrix)
	screen.DrawChar(stdout, "\x1b[101m ", row+1, col+1, matrix)
	screen.DrawChar(stdout, "\x1b[101m ", row+1, col+2, matrix)
	screen.DrawChar(stdout, "\x1b[101m ", row+1, col+3, matrix)
}

func DrawColumn(stdout *os.File, row, col int, matrix [][]int) {
	screen.DrawChar(stdout, "\x1b[101m ", row, col, matrix)
	screen.DrawChar(stdout, "\x1b[101m ", row+1, col, matrix)
	screen.DrawChar(stdout, "\x1b[101m ", row+2, col, matrix)
	screen.DrawChar(stdout, "\x1b[101m ", row+3, col, matrix)
}
