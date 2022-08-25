package main

import (
	"cchess/cchess/game"
	"cchess/cchess/piece"
	"fmt"
)

func main() {
	g := game.NewGame()

	//from := piece.NewPos(0, 0)
	//to := piece.NewPos(0, 2)
	//
	//from1 := piece.NewPos(0, 2)
	//to1 := piece.NewPos(0, 0)
	//
	////bboard.Move(x, y)
	//bboard.Move(from, to)
	//bboard.Move(from1, to1)
	//bboard.PrintBoard()
	//fmt.Println(bboard.GetMoveWays(piece.NewPos(0, 3)))
	//g.Show()
	//r, p := g.GetPieceByName("R", 0)
	//fmt.Println(r, p, "xxx")
	//fmt.Println("-----")
	//g.Run("")
	//g.RunCmd("兵一进一")
	//g.RunCmd("兵一进一")
	g.RunCmd("車1进1")
	//g.RunCmd("兵一进一")
	//g.RunCmd("车一进四")
	//g.RunCmd("炮二进六")
	//g.RunCmd("兵一进二")
	//g.RunCmd("车一进二")
	//g.RunCmd("炮二进三")
	//g.RunCmd("车九进二")
	//g.RunCmd("马二进三")
	//fmt.Println(g.GetMoveWays(g.Board.GetFetchByPos(piece.Pos{7, 2})))
	//fmt.Println(g.GetMoveWays(g.Board.GetFetchByPos(piece.Pos{0, 9})))
	//fmt.Println(g.GetMoveWays(g.Board.GetFetchByPos(piece.Pos{1, 7})))
	//fmt.Println(g.GetMoveWays(g.Board.GetFetchByPos(piece.Pos{1, 9})))
	//fmt.Println(g.GetMoveWays(g.Board.GetFetchByPos(piece.Pos{4, 9})))
	//fmt.Println(g.GetMoveWays(g.Board.GetFetchByPos(piece.Pos{4, 0})))
	fmt.Println(g.GetMoveWays(g.Board.GetFetchByPos(piece.Pos{6, 0})))
	//g.MoveUp(piece.Pos{8, 3}, 1)
	//g.Show()

	//g.Move(piece.Pos{0, 0}, piece.Pos{0, 1})
	g.Show()
	//ways := g.GetMoveWays(piece.Pos{0, 3})

	//fmt.Println(ways)

}
