package main

import (
	"fmt"
	"github.com/gookit/color"
	"strings"
	"unicode"
)

const NONE = 0
const KING = 1
const PAWN = 2
const KNIGHT = 3
const BISHOP = 4
const ROOK = 5
const QUEEN = 6
const WHITE = 8
const BLACK = 16

func main() {
	fmt.Println("Starting Derpfish")
	board := createNewBoard()
	displayBoard(WHITE, board)
}

func createNewBoard() *[64]int {
	return createBoardFromFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR")
}

func createBoardFromFen(fen string) *[64]int {
	board := [64]int{}
	pieceFromSymbol := map[string]int{
		"k": BLACK | KING, "p": BLACK | PAWN, "n": BLACK | KNIGHT, "b": BLACK | BISHOP, "r": BLACK | ROOK, "q": BLACK | QUEEN,
		"K": WHITE | KING, "P": WHITE | PAWN, "N": WHITE | KNIGHT, "B": WHITE | BISHOP, "R": WHITE | ROOK, "Q": WHITE | QUEEN,
	}

	fenBoard := strings.Split(fen, " ")[0]
	file := 0
	rank := 7

	for _, s := range fenBoard {
		if s == 47 {
			file = 0
			rank--
		} else if unicode.IsDigit(s) {
			file = file + int(s-'0')
		} else {
			ss := string(s)
			board[rank * 8 + file] = pieceFromSymbol[ss]
			file++
		}
	}
	return &board
}

func getBackgroundColor(side int, rank int, file int) color.Color {
	bgColor := color.BgWhite
	if (side == WHITE && file % 2 == rank % 2) || (side == BLACK && file % 2 != rank % 2) {
		bgColor = color.BgGray
	}
	return bgColor
}

func displayBoard(side int, board *[64]int) {
	if side == WHITE {
		for r := 7; r >= 0; r-- {
			for f := 0; f < 8; f++ {
				color.New(color.FgBlack, getBackgroundColor(side, r, f)).Printf(" %c ", getUnicodePrintOfSquare(board[r*8+f]))
			}
			fmt.Println()
		}
	} else if side == BLACK {
		for r := 0; r < 8; r++ {
			for f := 7; f >= 0; f-- {
				color.New(color.FgBlack, getBackgroundColor(side, r, f)).Printf(" %c ", getUnicodePrintOfSquare(board[r*8+f]))
			}
			fmt.Println()
		}
	}
}

func getUnicodePrintOfSquare(square int) rune {
	iconFromPiece := map[int]rune{
		NONE: ' ',
		BLACK | KING: '\u265A', BLACK | PAWN: '\u265F', BLACK | KNIGHT: '\u265E', BLACK | BISHOP: '\u265D', BLACK | ROOK: '\u265C', BLACK | QUEEN: '\u265B',
		WHITE | KING: '\u2654', WHITE | PAWN: '\u2659', WHITE | KNIGHT: '\u2658', WHITE | BISHOP: '\u2657', WHITE | ROOK: '\u2656', WHITE | QUEEN: '\u2655',
	}
	return iconFromPiece[square]
}