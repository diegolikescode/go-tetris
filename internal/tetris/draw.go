package tetris

import (
	"math"
	"os"

	"github.com/diegolikescode/go-tetris/internal/screen"
)

// REFERENCE: https://www.reddit.com/media?url=https%3A%2F%2Fexternal-preview.redd.it%2FrMJqW-zViDxoKTkODFO00VnEyKKWQiC0CnCuLvwbxc8.jpg%3Fauto%3Dwebp%26s%3D758c506e60811db7a78b107d53c61e1868f23221

// NOTE: every row has to occupy 2 cells because of the 2/1 line-height to width ratio

const (
	teeStyle       string = "\x1b[48;5;129m"
	lStyle         string = "\x1b[48;5;208m"
	lInvertedStyle string = "\x1b[48;5;27m"
	zStyle         string = "\x1b[101m"
	zInvertedStyle string = "\x1b[48;5;46m"
	squareStyle    string = "\x1b[48;5;226m"
	columnStyle    string = "\x1b[48;5;51m"
)

type PieceType string

const (
	tee       PieceType = "tee"
	l         PieceType = "l"
	lInverted PieceType = "lInverted"
	z         PieceType = "z"
	zInverted PieceType = "zInverted"
	square    PieceType = "square"
	column    PieceType = "column"
)

type Positioning struct {
	Row    int
	Column int
}

type Piece struct {
	Position []Positioning
	Name     PieceType
	Style    string
}

func (p *Piece) ClearPiece(matrix [][]screen.Cell) {
	for _, p := range p.Position {
		screen.DrawChar(os.Stdout, screen.Clear, screen.ComposeANSI(" ", screen.Reset), p.Row, p.Column, matrix)
	}
}

func (p *Piece) MoveLeft(matrix [][]screen.Cell) {
	leftPoint64 := float64(math.MaxInt)
	leftiestPoint := Positioning{}
	for i := range p.Position {
		newLeftPoint64 := math.Min(leftPoint64, float64(p.Position[i].Column))
		if leftPoint64 != newLeftPoint64 {
			leftiestPoint.Row = p.Position[i].Row
			leftiestPoint.Column = p.Position[i].Column

			leftPoint64 = newLeftPoint64
		}
	}

	if !matrix[leftiestPoint.Row][leftiestPoint.Column-1].Occupied && !matrix[leftiestPoint.Row][leftiestPoint.Column-2].Occupied {
		p.ClearPiece(matrix)
		for i := range p.Position {
			screen.DrawChar(os.Stdout, screen.Clear, screen.ComposeANSI(" ", p.Style), p.Position[i].Row, p.Position[i].Column-2, matrix)
			p.Position[i].Column = p.Position[i].Column - 2
		}
	}
}

func (p *Piece) MoveRight(matrix [][]screen.Cell) {
	rightPoint64 := 0.0
	rightiestPoint := Positioning{}
	for i := range p.Position {
		newRightPoint64 := math.Max(rightPoint64, float64(p.Position[i].Column))
		if rightPoint64 != newRightPoint64 {
			rightiestPoint.Row = p.Position[i].Row
			rightiestPoint.Column = p.Position[i].Column
		}
	}

	if !matrix[rightiestPoint.Row][rightiestPoint.Column+1].Occupied && !matrix[rightiestPoint.Row][rightiestPoint.Column+2].Occupied {
		p.ClearPiece(matrix)
		for i := range p.Position {
			screen.DrawChar(os.Stdout, screen.Clear, screen.ComposeANSI(" ", p.Style), p.Position[i].Row, p.Position[i].Column+2, matrix)
			p.Position[i].Column = p.Position[i].Column + 2
		}
	}
}

func (p *Piece) MoveDown(matrix [][]screen.Cell) {
	lowPoint64 := 0.0
	lowestPoint := Positioning{}
	for i := range p.Position {
		newLowPoint64 := math.Max(lowPoint64, float64(p.Position[i].Row))
		if lowPoint64 != newLowPoint64 {
			lowestPoint.Row = p.Position[i].Row
			lowestPoint.Column = p.Position[i].Column
		}
	}

	if !matrix[lowestPoint.Row+1][lowestPoint.Column].Occupied {
		p.ClearPiece(matrix)
		for i := range p.Position {
			screen.DrawChar(os.Stdout, screen.Clear, screen.ComposeANSI(" ", p.Style), p.Position[i].Row+1, p.Position[i].Column, matrix)
			p.Position[i].Row = p.Position[i].Row + 1
		}
	}
}

// TODO: fix
func (p *Piece) SmashDown(matrix [][]screen.Cell) {
	lowPoint64 := 0.0
	lowestPoint := Positioning{}

	for i := range p.Position {
		newLowPoint64 := math.Max(lowPoint64, float64(p.Position[i].Row))
		if lowPoint64 != newLowPoint64 {
			lowestPoint.Row = p.Position[i].Row
			lowestPoint.Column = p.Position[i].Column
		}
	}

	heighestPointOffset := lowestPoint.Row - p.Position[0].Row

	smashingPoint := Positioning{}
	for i := range len(matrix) {
		row := matrix[i]

		if row[lowestPoint.Column].Occupied && row[lowestPoint.Column].Type != screen.PermeableInterface {
			smashingPoint.Row = i - heighestPointOffset
			smashingPoint.Column = lowestPoint.Column
			break
		}
	}

	if !matrix[lowestPoint.Row+1][lowestPoint.Column].Occupied {
		p.ClearPiece(matrix)
		for i := range p.Position {
			row := smashingPoint.Row - heighestPointOffset
			screen.DrawChar(os.Stdout, screen.Clear, screen.ComposeANSI(" ", p.Style), row, p.Position[i].Column, matrix)
			p.Position[i].Row = smashingPoint.Row
		}
	}
}

func DrawTee(stdout *os.File, row, col int, matrix [][]screen.Cell) Piece {
	p := Piece{Position: []Positioning{}}
	p.Name = tee
	p.Style = teeStyle

	p.Position = append(p.Position, Positioning{Row: row, Column: col})
	p.Position = append(p.Position, Positioning{Row: row, Column: col + 1})
	p.Position = append(p.Position, Positioning{Row: row, Column: col + 2})
	p.Position = append(p.Position, Positioning{Row: row, Column: col + 3})
	p.Position = append(p.Position, Positioning{Row: row, Column: col + 4})
	p.Position = append(p.Position, Positioning{Row: row, Column: col + 5})
	p.Position = append(p.Position, Positioning{Row: row + 1, Column: col + 2})
	p.Position = append(p.Position, Positioning{Row: row + 1, Column: col + 3})

	for _, p := range p.Position {
		screen.DrawChar(stdout, screen.Piece, screen.ComposeANSI(" ", teeStyle), p.Row, p.Column, matrix)
	}

	return p
}

func DrawL(stdout *os.File, row, col int, matrix [][]screen.Cell) Piece {
	p := Piece{Position: []Positioning{}}
	p.Name = l
	p.Style = lStyle

	p.Position = append(p.Position, Positioning{Row: row, Column: col})
	p.Position = append(p.Position, Positioning{Row: row, Column: col + 1})

	p.Position = append(p.Position, Positioning{Row: row + 1, Column: col})
	p.Position = append(p.Position, Positioning{Row: row + 1, Column: col + 1})

	p.Position = append(p.Position, Positioning{Row: row + 2, Column: col})
	p.Position = append(p.Position, Positioning{Row: row + 2, Column: col + 1})

	p.Position = append(p.Position, Positioning{Row: row + 2, Column: col + 2})
	p.Position = append(p.Position, Positioning{Row: row + 2, Column: col + 3})

	for _, p := range p.Position {
		screen.DrawChar(stdout, screen.Piece, screen.ComposeANSI(" ", lStyle), p.Row, p.Column, matrix)
	}

	return p
}

func DrawLInverted(stdout *os.File, row, col int, matrix [][]screen.Cell) Piece {
	p := Piece{}
	p.Name = lInverted
	p.Style = lInvertedStyle

	p.Position = append(p.Position, Positioning{Row: row, Column: col})
	p.Position = append(p.Position, Positioning{Row: row, Column: col + 1})

	p.Position = append(p.Position, Positioning{Row: row + 1, Column: col})
	p.Position = append(p.Position, Positioning{Row: row + 1, Column: col + 1})

	p.Position = append(p.Position, Positioning{Row: row + 2, Column: col})
	p.Position = append(p.Position, Positioning{Row: row + 2, Column: col + 1})

	p.Position = append(p.Position, Positioning{Row: row + 2, Column: col - 1})
	p.Position = append(p.Position, Positioning{Row: row + 2, Column: col - 2})

	for _, p := range p.Position {
		screen.DrawChar(stdout, screen.Piece, screen.ComposeANSI(" ", lInvertedStyle), p.Row, p.Column, matrix)
	}

	return p
}

func DrawZ(stdout *os.File, row, col int, matrix [][]screen.Cell) Piece {
	p := Piece{}
	p.Name = z
	p.Style = zStyle

	p.Position = append(p.Position, Positioning{Row: row, Column: col})
	p.Position = append(p.Position, Positioning{Row: row, Column: col + 1})
	p.Position = append(p.Position, Positioning{Row: row, Column: col + 2})
	p.Position = append(p.Position, Positioning{Row: row, Column: col + 3})

	p.Position = append(p.Position, Positioning{Row: row + 1, Column: col + 2})
	p.Position = append(p.Position, Positioning{Row: row + 1, Column: col + 3})
	p.Position = append(p.Position, Positioning{Row: row + 1, Column: col + 4})
	p.Position = append(p.Position, Positioning{Row: row + 1, Column: col + 5})

	for _, p := range p.Position {
		screen.DrawChar(stdout, screen.Piece, screen.ComposeANSI(" ", zStyle), p.Row, p.Column, matrix)
	}

	return p
}

func DrawZInverted(stdout *os.File, row, col int, matrix [][]screen.Cell) Piece {
	p := Piece{}
	p.Name = zInverted
	p.Style = zInvertedStyle

	p.Position = append(p.Position, Positioning{Row: row, Column: col})
	p.Position = append(p.Position, Positioning{Row: row, Column: col - 1})
	p.Position = append(p.Position, Positioning{Row: row, Column: col - 2})
	p.Position = append(p.Position, Positioning{Row: row, Column: col - 3})

	p.Position = append(p.Position, Positioning{Row: row + 1, Column: col - 2})
	p.Position = append(p.Position, Positioning{Row: row + 1, Column: col - 3})
	p.Position = append(p.Position, Positioning{Row: row + 1, Column: col - 4})
	p.Position = append(p.Position, Positioning{Row: row + 1, Column: col - 5})

	for _, p := range p.Position {
		screen.DrawChar(stdout, screen.Piece, screen.ComposeANSI(" ", zInvertedStyle), p.Row, p.Column, matrix)
	}

	return p
}

func DrawSquare(stdout *os.File, row, col int, matrix [][]screen.Cell) Piece {
	p := Piece{}
	p.Name = square
	p.Style = squareStyle

	p.Position = append(p.Position, Positioning{Row: row, Column: col})
	p.Position = append(p.Position, Positioning{Row: row, Column: col + 1})
	p.Position = append(p.Position, Positioning{Row: row, Column: col + 2})
	p.Position = append(p.Position, Positioning{Row: row, Column: col + 3})

	p.Position = append(p.Position, Positioning{Row: row + 1, Column: col})
	p.Position = append(p.Position, Positioning{Row: row + 1, Column: col + 1})
	p.Position = append(p.Position, Positioning{Row: row + 1, Column: col + 2})
	p.Position = append(p.Position, Positioning{Row: row + 1, Column: col + 3})

	for _, p := range p.Position {
		screen.DrawChar(stdout, screen.Piece, screen.ComposeANSI(" ", squareStyle), p.Row, p.Column, matrix)
	}

	return p
}

func DrawColumn(stdout *os.File, row, col int, matrix [][]screen.Cell) Piece {
	p := Piece{}
	p.Name = column
	p.Style = columnStyle

	p.Position = append(p.Position, Positioning{Row: row, Column: col})
	p.Position = append(p.Position, Positioning{Row: row + 1, Column: col})
	p.Position = append(p.Position, Positioning{Row: row + 2, Column: col})
	p.Position = append(p.Position, Positioning{Row: row + 3, Column: col})

	p.Position = append(p.Position, Positioning{Row: row, Column: col + 1})
	p.Position = append(p.Position, Positioning{Row: row + 1, Column: col + 1})
	p.Position = append(p.Position, Positioning{Row: row + 2, Column: col + 1})
	p.Position = append(p.Position, Positioning{Row: row + 3, Column: col + 1})

	for _, p := range p.Position {
		screen.DrawChar(stdout, screen.Piece, screen.ComposeANSI(" ", columnStyle), p.Row, p.Column, matrix)
	}

	return p
}
