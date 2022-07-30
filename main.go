package main

import (
	"cchess/cchess/board"
	"cchess/cchess/piece"
)

func main() {
	bboard := board.BaseChessBoard{}
	FULL_INIT_FEN := "rnbakabnr/9/1c5c1/p1p1p1p1p/9/9/P1P1P1P1P/1C5C1/9/RNBAKABNR w - - 0 1"
	//FULL_INIT_FEN := "rn"
	bboard.FromFen(FULL_INIT_FEN)
	//bboard.PrintBoard()
	x := piece.NewPos(0, 0)
	y := piece.NewPos(0, 4)

	from := piece.NewPos(1, 0)
	to := piece.NewPos(1, 1)

	bboard.Move(x, y)
	bboard.Move(from, to)
	bboard.PrintBoard()
}
