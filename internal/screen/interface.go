package screen

import (
	"os"
)

var (
	MiddleRowHold int = 0
	MiddleColHold int = 0

	MiddleRowMain int = 0
	MiddleColMain int = 0
)

var (
	StartMainRow int = 0
	StartMainCol int = 0
)

var (
	QUIT_TEXT       string = "Q     - Quit" // just for testing
	LEFT_TEXT       string = "A     - Move Left"
	RIGHT_TEXT      string = "D     - Move Right"
	ROLL_TEXT       string = "W     - Turn Piece"
	ACCELERATE_TEXT string = "S     - Accelerate Down"
	HOLD_TEXT       string = "F     - Hold!"
	SMASH_TEXT      string = "Space - Smash Down!"

	COMMAND_TEXTS = []string{QUIT_TEXT, LEFT_TEXT, RIGHT_TEXT, ROLL_TEXT, ACCELERATE_TEXT, SMASH_TEXT}
)

func DrawMainBlock(stdout *os.File, middleRow, middleColumn int, matrix [][]Cell) {
	StartMainRow = middleRow - MAIN_BLOCK_HEIGHT_OFFSET
	StartMainCol = middleColumn

	MiddleRowMain = middleRow
	MiddleColMain = middleColumn

	var c int
	for c = 0; c < 9; c++ {
		DrawChar(stdout, Interface, ComposeANSI(V, Reset), middleRow+c, middleColumn+MAIN_BLOCK_WIDTH_OFFSET, matrix)
		DrawChar(stdout, Interface, ComposeANSI(V, Reset), middleRow+c, middleColumn-MAIN_BLOCK_WIDTH_OFFSET, matrix)

		DrawChar(stdout, Interface, ComposeANSI(V, Reset), middleRow-c-1, middleColumn-MAIN_BLOCK_WIDTH_OFFSET, matrix)
		DrawChar(stdout, Interface, ComposeANSI(V, Reset), middleRow-c-1, middleColumn+MAIN_BLOCK_WIDTH_OFFSET, matrix)
	}

	for c = 0; c < MAIN_BLOCK_WIDTH_OFFSET; c++ {
		DrawChar(stdout, PermeableInterface, ComposeANSI(H, Reset), middleRow-MAIN_BLOCK_HEIGHT_OFFSET, middleColumn+c, matrix)
		DrawChar(stdout, PermeableInterface, ComposeANSI(H, Reset), middleRow-MAIN_BLOCK_HEIGHT_OFFSET, middleColumn-c, matrix)

		DrawChar(stdout, Interface, ComposeANSI(H, Reset), middleRow+MAIN_BLOCK_HEIGHT_OFFSET-1, middleColumn-c, matrix)
		DrawChar(stdout, Interface, ComposeANSI(H, Reset), middleRow+MAIN_BLOCK_HEIGHT_OFFSET-1, middleColumn+c, matrix)

		// DrawChar(stdout, Interface, ComposeANSI("▀", Reset), middleRow+MAIN_BLOCK_HEIGHT_OFFSET-1, middleColumn-c, matrix)
		// DrawChar(stdout, Interface, ComposeANSI("▀", Reset), middleRow+MAIN_BLOCK_HEIGHT_OFFSET-1, middleColumn+c, matrix)
	}

	DrawChar(stdout, Interface, ComposeANSI(TR, Reset), middleRow-MAIN_BLOCK_HEIGHT_OFFSET, middleColumn+c, matrix)
	DrawChar(stdout, Interface, ComposeANSI(TL, Reset), middleRow-MAIN_BLOCK_HEIGHT_OFFSET, middleColumn-c, matrix)
	DrawChar(stdout, Interface, ComposeANSI(BL, Reset), middleRow+MAIN_BLOCK_HEIGHT_OFFSET-1, middleColumn-c, matrix)
	DrawChar(stdout, Interface, ComposeANSI(BR, Reset), middleRow+MAIN_BLOCK_HEIGHT_OFFSET-1, middleColumn+c, matrix)

	// DrawChar(stdout, Interface, ComposeANSI("▀", Reset), middleRow+MAIN_BLOCK_HEIGHT_OFFSET-1, middleColumn-c, matrix)
	// DrawChar(stdout, Interface, ComposeANSI("▀", Reset), middleRow+MAIN_BLOCK_HEIGHT_OFFSET-1, middleColumn+c, matrix)
}

func DrawHoldBlock(stdout *os.File, middleRow, middleCol int, matrix [][]Cell) {
	mainCol := middleCol - MAIN_BLOCK_WIDTH_OFFSET - 1
	mainRow := middleRow - MAIN_BLOCK_HEIGHT_OFFSET + 1

	LEFT_COL := 14

	MiddleColHold = mainCol - LEFT_COL/4 - 4
	MiddleRowHold = mainRow + LEFT_COL/4 - 1

	var c int
	for c = 0; c < LEFT_COL; c++ {
		DrawChar(stdout, Interface, ComposeANSI(H, Reset), mainRow, mainCol-c, matrix)
		DrawChar(stdout, Interface, ComposeANSI(H, Reset), mainRow+LEFT_COL/2, mainCol-c, matrix)

		// because I only need half of it
		if c < LEFT_COL/2 {
			DrawChar(stdout, Interface, ComposeANSI(V, Reset), mainRow+c+1, mainCol-LEFT_COL, matrix)
		}
	}
	DrawChar(stdout, Interface, ComposeANSI("┌", ""), mainRow, mainCol-LEFT_COL, matrix)
	DrawChar(stdout, Interface, ComposeANSI("└", ""), mainRow+c/2, mainCol-LEFT_COL, matrix)

	DrawChar(stdout, Interface, ComposeANSI("[", bold, redPink), mainRow, mainCol-LEFT_COL/2-2, matrix)
	DrawChar(stdout, Interface, ComposeANSI("H", bold, redPink), mainRow, mainCol-LEFT_COL/2-1, matrix)
	DrawChar(stdout, Interface, ComposeANSI("O", bold, redPink), mainRow, mainCol-LEFT_COL/2, matrix)
	DrawChar(stdout, Interface, ComposeANSI("L", bold, redPink), mainRow, mainCol-LEFT_COL/2+1, matrix)
	DrawChar(stdout, Interface, ComposeANSI("D", bold, redPink), mainRow, mainCol-LEFT_COL/2+2, matrix)
	DrawChar(stdout, Interface, ComposeANSI("]", bold, redPink), mainRow, mainCol-LEFT_COL/2+3, matrix)
}

func DrawCommandTexts(stdout *os.File, startRow, startCol int, matrix [][]Cell) {
	for i := range COMMAND_TEXTS {
		DrawMultiChars(stdout,
			Interface,
			COMMAND_TEXTS[i],
			startRow+3+i,
			startCol-MAIN_BLOCK_WIDTH_OFFSET-25,
			matrix,
		)
	}
}
