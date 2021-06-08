package main

import (
	"bufio"
	"fmt"
	"github.com/gookit/color"
	"os"
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
	board = createBoardFromFen("r1b1k1nr/p2p1pNp/n2B4/1p1NP2P/6P1/3P1Q2/P1P1K3/q5b1") // REPLACE STARTER BOARD FOR TESTING
	startGame(board)
}

func startGame(board *[64]int) {
	reader := bufio.NewReader(os.Stdin)
	currentPlayer := WHITE
	for {
		displayBoard(currentPlayer, board, []int{})
		fromSquare := getCommand("From ", reader)
		allowedMoves := getAllowedMoves(currentPlayer, board, fromSquare)
		fmt.Println(allowedMoves)
		displayBoard(currentPlayer, board, allowedMoves)
		toSquare := getCommand(fromSquare + " -> ? ", reader)
		fmt.Println(fromSquare, toSquare)
	}
}

func getCommand(inputtext string, reader *bufio.Reader) string {
	fmt.Print(inputtext)
	cmd, _ := reader.ReadString('\n')
	return strings.Replace(cmd, "\n", "", -1)
}

func createNewBoard() *[64]int {
	return createBoardFromFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR")
}

func getAllowedMoves(currentPlayer int, board *[64]int, fromSquare string) []int {
	position := []rune(fromSquare)
	file := position[0]-97
	rank := position[1]-49
	square := int(rank * 8 + file)

	uncoloredPiece := board[square] &^ currentPlayer
	if uncoloredPiece > 8 {
		return []int{}
	}

	return getMovementForPiece(board, square, uncoloredPiece)
}

func getMovementForPiece(board *[64]int, square int, piece int) []int {
	switch piece {
	case PAWN:
		return getPawnMovement(board, square)
	case KING:
		return getKingMovement(board, square)
	case KNIGHT:
		return getKnightMovement(board, square)
	case BISHOP:
		fmt.Println("I AM BISH")
		return getPossibleLinearMovement(board, square, []int{7, 9, -7, -9})
	case ROOK:
		return getPossibleLinearMovement(board, square, []int{1, 8, -1, -8})
	case QUEEN:
		return getPossibleLinearMovement(board, square, []int{7, 9, -7, -9, 1, 8, -1, -8})
	default:
		return []int{}
	}
}

func getPawnMovement(board *[64]int, square int) []int {
	return []int{}
}

func getKingMovement(board *[64]int, square int) []int {
	return []int{}
}

func getKnightMovement(board *[64]int, square int) []int {
	return []int{}
}

func getPossibleLinearMovement(board *[64]int, square int, directions []int) []int {
	var possibleMoves []int
	for _, d := range directions {
		checkedSquare := square + d
		for {
			fmt.Println("SQUARE")
			fmt.Println(checkedSquare)
			if (checkedSquare) < 0 || (checkedSquare) >= 64 || board[checkedSquare] != NONE {
				break
			}
			possibleMoves = append(possibleMoves, checkedSquare)
			checkedSquare += d
		}
	}
	return possibleMoves
}

func createBoardFromFen(fen string) *[64]int {
	board := [64]int{}
	pieceFromSymbol := map[string]int {
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

func displayBoard(side int, board *[64]int, possibleMovement []int) {
	if side == WHITE {
		for r := 7; r >= 0; r-- {
			fmt.Print(r+1)
			for f := 0; f < 8; f++ {
				printSquare(board, side, r, f, has(possibleMovement, r*8+f))
			}
			fmt.Println()
		}
		fmt.Println("  a  b  c  d  e  f  g  h")
	} else if side == BLACK {
		for r := 0; r < 8; r++ {
			fmt.Print(r+1)
			for f := 7; f >= 0; f-- {
				printSquare(board, side, r, f, has(possibleMovement, r*8+f))
			}
			fmt.Println()
		}
		fmt.Println("  h  g  f  e  d  c  b  a")
	}
}

func has(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func printSquare(board *[64]int, side int, rank int, file int, canMove bool) {
	fg := color.FgBlack
	if (board[rank*8+file]) < BLACK {
		fg = color.FgWhite
	}
	if canMove {
		color.New(fg, color.BgRed).Printf(" %c ", getUnicodePrintOfSquare(board[rank*8+file]))
	} else {
		color.New(fg, getBackgroundColor(side, rank, file)).Printf(" %c ", getUnicodePrintOfSquare(board[rank*8+file]))
	}
}

func getBackgroundColor(side int, rank int, file int) color.Color {
	bgColor := color.BgWhite
	if (side == WHITE && file % 2 == rank % 2) || (side == BLACK && file % 2 != rank % 2) {
		bgColor = color.BgGray
	}
	return bgColor
}

// TODO: Might simplify iconFromPiece later
func getUnicodePrintOfSquare(square int) rune {
	iconFromPiece := map[int]rune{
		NONE: ' ',
		BLACK | KING: '\u265A', BLACK | PAWN: '\u265F', BLACK | KNIGHT: '\u265E', BLACK | BISHOP: '\u265D', BLACK | ROOK: '\u265C', BLACK | QUEEN: '\u265B',
		WHITE | KING: '\u265A', WHITE | PAWN: '\u265F', WHITE | KNIGHT: '\u265E', WHITE | BISHOP: '\u265D', WHITE | ROOK: '\u265C', WHITE | QUEEN: '\u265B',
	}
	return iconFromPiece[square]
}