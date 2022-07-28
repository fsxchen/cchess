package main

import (
	"cchess/cchess/board"
)

func main() {
	board := board.BaseChessBoard{}
	//FULL_INIT_FEN := "rnbakabnr/9/1c5c1/p1p1p1p1p/9/9/P1P1P1P1P/1C5C1/9/RNBAKABNR w - - 0 1"
	FULL_INIT_FEN := "rn"
	board.FromFen(FULL_INIT_FEN)
	board.PrintBoard()
}
