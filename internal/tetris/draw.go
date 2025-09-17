package tetris

import (
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

func (p *Piece) MovePiece(startCell Positioning, matrix [][]screen.Cell) {
	relativePositions := []Positioning{}
	for i := range p.Position {
		relativePositions[i].Row = p.Position[i].Row - startCell.Row
		relativePositions[i].Column = p.Position[i].Column - startCell.Column
	}

	p.ClearPiece(matrix)
	for i := range p.Position {
		screen.DrawChar(os.Stdout, screen.Clear, screen.ComposeANSI(" ", p.Style), p.Position[i].Row-relativePositions[i].Row, p.Position[i].Column-relativePositions[i].Row, matrix)
		p.Position[i].Column = p.Position[i].Column - 2
	}
}

func (p *Piece) MoveLeft(matrix [][]screen.Cell) {
	if p.CanGoLeft(matrix) {
		p.ClearPiece(matrix)
		for i := range p.Position {
			screen.DrawChar(os.Stdout, screen.Clear, screen.ComposeANSI(" ", p.Style), p.Position[i].Row, p.Position[i].Column-2, matrix)
			p.Position[i].Column = p.Position[i].Column - 2
		}
	}
}

func (p *Piece) MoveRight(matrix [][]screen.Cell) {
	if p.CanGoRight(matrix) {
		p.ClearPiece(matrix)
		for i := range p.Position {
			screen.DrawChar(os.Stdout, screen.Clear, screen.ComposeANSI(" ", p.Style), p.Position[i].Row, p.Position[i].Column+2, matrix)
			p.Position[i].Column = p.Position[i].Column + 2
		}
	}
}

/*
XXXXXX
  XX
[0,0],[0,1]|[0,2],[0,3]|[0,4],[0,5]
           |[1,2],[1,3]|
*/

func (p *Piece) MoveDown(matrix [][]screen.Cell) {
	if p.CanGoDown(matrix) {
		p.ClearPiece(matrix)
		for i := range p.Position {
			screen.DrawChar(os.Stdout, screen.Piece, screen.ComposeANSI(" ", p.Style), p.Position[i].Row+1, p.Position[i].Column, matrix)
			p.Position[i].Row = p.Position[i].Row + 1
		}
	}
}

/*
XXXXXX
  XX
[0,0],[0,1]|[0,2],[0,3]|[0,4],[0,5]
	       |[1,2],[1,3]|
*/

func (p *Piece) CanGoLeft(matrix [][]screen.Cell) bool {
	// map: row -> col
	positionsDict := make(map[int]int)

	for _, pos := range p.Position {
		if positionsDict[pos.Row] == 0 || positionsDict[pos.Row] > pos.Column {
			positionsDict[pos.Row] = pos.Column
		}
	}

	canMove := true
	for k, v := range positionsDict {
		if (matrix[k][v-1].Occupied || matrix[k][v-2].Occupied) && (matrix[k][v-1].Type != screen.PermeableInterface || matrix[k][v-2].Type != screen.PermeableInterface) {
			return false
		}
	}
	return canMove
}

func (p *Piece) CanGoRight(matrix [][]screen.Cell) bool {
	// map: row -> col
	positionsDict := make(map[int]int)

	for _, pos := range p.Position {
		if positionsDict[pos.Row] == 0 || positionsDict[pos.Row] < pos.Column {
			positionsDict[pos.Row] = pos.Column
		}
	}

	canMove := true
	for k, v := range positionsDict {
		if (matrix[k][v+1].Occupied || matrix[k][v+2].Occupied) && (matrix[k][v+1].Type != screen.PermeableInterface || matrix[k][v+2].Type != screen.PermeableInterface) {
			return false
		}
	}
	return canMove
}

func (p *Piece) CanGoDown(matrix [][]screen.Cell) bool {
	// map: column -> row
	positionsDict := make(map[int]int)

	for _, pos := range p.Position {
		if positionsDict[pos.Column] == 0 || positionsDict[pos.Column] < pos.Row {
			positionsDict[pos.Column] = pos.Row
		}
	}

	canMove := true
	for k, v := range positionsDict {
		if matrix[v+1][k].Occupied && matrix[v+1][k].Type != screen.PermeableInterface {
			return false
		}
	}
	return canMove
}

// @todo
func (p *Piece) SmashDown(matrix [][]screen.Cell) {
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
