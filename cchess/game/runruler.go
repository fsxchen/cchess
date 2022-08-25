package game

import "cchess/cchess/piece"

// 有些规则是在走的时候检查的

// 是否马不可以走,别了
func (g *Game) CanNotHrun(from, to piece.Pos) bool {
	fromFentch := g.Board.GetFetchByPos(from)
	var nextFetch *Fetch
	if fromFentch == nil || !g.CanTo(fromFentch, to) {
		panic("错误了")
	}
	if to.X-from.X == 2 {
		nextFetch = g.Board.GetFetchByPos(piece.NewPos(from.X+1, from.Y))
	}
	if to.X-from.X == -2 {
		nextFetch = g.Board.GetFetchByPos(piece.NewPos(from.X-1, from.Y))
	}
	if to.Y-from.Y == 2 {
		nextFetch = g.Board.GetFetchByPos(piece.NewPos(from.X, from.Y+1))
	}
	if to.Y-from.Y == -2 {
		nextFetch = g.Board.GetFetchByPos(piece.NewPos(from.X, from.Y-1))
	}
	return nextFetch != nil
}
