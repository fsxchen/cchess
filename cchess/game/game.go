package game

import (
	"cchess/cchess/board"
	"cchess/cchess/piece"
	"cchess/util"
	"errors"
	"fmt"
)

type Game struct {
	RAte  []*piece.Piece
	BAte  []*piece.Piece
	Board FetchBoard
}

func NewGame() *Game {
	board := *NewFetchBoard()
	FULL_INIT_FEN := "rnbakabnr/9/1c5c1/p1p1p1p1p/9/9/P1P1P1P1P/1C5C1/9/RNBAKABNR w - - 0 1"
	//FULL_INIT_FEN := "rn"
	board.FromFen(FULL_INIT_FEN)

	RAte := make([]*piece.Piece, 0)
	BAte := make([]*piece.Piece, 0)
	return &Game{Board: board, RAte: RAte, BAte: BAte}
}

func (g *Game) MoveUp(fetch *Fetch, step int) {
	// 一个棋子前进
	var toY int
	var target piece.Pos
	from := fetch.pos
	if fetch.piece.CanMoveLine() {
		if fetch.piece.Side() == piece.Red {
			toY = from.Y + step
		} else if fetch.piece.Side() == piece.Black {
			toY = from.Y - step
		}
		target = piece.Pos{X: from.X, Y: toY}
	} else if fetch.CanMoveDiagonal() {
		// 如果走的不是直线，那么，此时的step表示表示的就是列了
		mayPosLis := g.GetMoveWays(fetch)
		for _, mayPos := range mayPosLis {
			var targetCol int
			if fetch.Side() == piece.Red {
				targetCol = 8 - step + 1
				if mayPos.X == targetCol && mayPos.Y > from.Y {
					target = mayPos
				}

			} else if fetch.Side() == piece.Black {
				if mayPos.X == step && mayPos.Y < from.Y {
					target = mayPos
				}
			}
		}
	}
	g.Move(fetch, target)
}

func (g *Game) RunCmd(cmd string) {
	// 执行一个棋谱的语句
	val, col, direct, toCmd := McmdParse(cmd)
	f := g.GetFetchByCmd(val, col)
	if f == nil {
		fmt.Println("错误")
		return
	}
	switch direct {
	case "U":
		g.MoveUp(f, toCmd)
	case "LR":
		g.MoveLR(f, toCmd)
	case "D":
		g.MoveDown(f, toCmd)
	}
}

// 获取某行两个棋子中间的棋子
func (g *Game) getMidFetch(from, to piece.Pos) (res []*Fetch) {
	// 如果是行，那么Y是必须是相同的
	res = make([]*Fetch, 0)
	if from.Y == to.Y && from.X != to.X {
		s, e := util.SortTwoInt(from.X, to.X)
		for i := s + 1; i < e; i++ {
			f := g.Board.GetFetchByPos(piece.Pos{X: i, Y: from.Y})
			if f != nil {
				res = append(res, f)
			}
		}
	}
	if from.X == to.X && from.Y != to.Y {
		s, e := util.SortTwoInt(from.Y, to.Y)
		//for _, fetch := range g.FetchList {
		//	if fetch.pos.X == from.X && fetch.pos.Y > s && fetch.pos.Y < e {
		//		res = append(res, fetch)
		//	}
		//}
		for i := s + 1; i < e; i++ {
			f := g.Board.GetFetchByPos(piece.Pos{X: from.X, Y: i})
			if f != nil {
				res = append(res, f)
			}
		}
	}
	return
}

func (g *Game) Move(fromFetch *Fetch, to piece.Pos) error {
	//  移动一个棋子到另外一个地方
	// 移动的逻辑如下：
	// 1. 判断本地是否有
	// 2. 是否符合走子规范
	// 3. 目的点情况判断。如果没有：
	//					1. 路径规范判断
	//                 如果有：
	//                  1. 是否吃子
	//                  2. 路径判断
	if !fromFetch.piece.CanTo(fromFetch.pos, to, false) {
		return errors.New("走子规则不允许")
	}

	if fromFetch.pos.X == to.X && fromFetch.pos.Y == to.Y {
		return errors.New("error! 落子无效")
	}

	toFetch := g.Board.GetFetchByPos(to)

	if toFetch == nil {
		// 走子逻辑，走子的逻辑在于，如果是走直线的子，不能有阻挡
		if !g.CanTo(fromFetch, to) {
			return errors.New("走子错误")
		}
		// 构造固定的区域，从左边下到右边上

	} else {
		if toFetch.piece.Side() == fromFetch.piece.Side() {
			return errors.New("error 错误，两边一样")
		} else {

			//	// 吃子的逻辑
			//	// 除了炮，走子和吃子的逻辑不一样
			//	if fromPiece.CanEatDirectly() && fromPiece.CanMoveLine() {
			//		// 能直线行走，直线吃的棋子，车/兵
			//		if fromPiece.CanTo(from, to, fromPiece.IsCrossRiver()) {
			//
			//		}
			//	} else {
			//		g.AddAteList(to)
			//	}
			//
		}
	}
	//g.Board.Move(fromFetch.pos, to)
	g.Board.Move(fromFetch.pos, to)
	fromFetch.pos = to
	return nil
}

func (g *Game) GetFetchByCmd(val string, col int) *Fetch {
	for i := 0; i < 10; i++ {
		f := g.Board.GetFetchByIndex(i, col)
		if f != nil && val == f.piece.Value() {
			return f
		}
	}
	return nil
}

func (g *Game) Show() {
	g.Board.PrintBoard()
}

// GetMoveWays 获取一个棋子在当前棋盘上所有可以移动的方向
func (g *Game) GetMoveWays(f *Fetch) []piece.Pos {
	if f == nil {
		panic("不可用的棋子")
	}
	result := make([]piece.Pos, 0) // 接收结果
	//bb := &g.Board

	if f.piece.CanMoveLine() {
		// 走直线，判断上下左右是否有路
		// 判断上方
		// 例如，输入Y= 【0， 1】
		step := f.piece.Step()
		ulimt := util.MinInt(f.pos.Y+step+1, board.HEIGH)
		var nf *Fetch
		uMid := 0 // 记录数量
		for i := f.pos.Y + 1; i < ulimt; i++ {
			nf = g.Board.GetFetchByPos(piece.Pos{X: f.pos.X, Y: i})
			if nf == nil && g.CanTo(f, piece.Pos{X: f.pos.X, Y: i}) && uMid == 0 {
				result = append(result, piece.Pos{X: f.pos.X, Y: i})
			} else {
				uMid++
				if g.CanEat(f, nf) {
					result = append(result, nf.pos)
				}
			}
		}
		//	// 判断下方地址
		dlimt := util.MaxInt(f.pos.Y-step, 0)
		dMid := 0
		for i := f.pos.Y - 1; i >= dlimt; i-- {
			nf = g.Board.GetFetchByPos(piece.Pos{X: f.pos.X, Y: i})
			if nf == nil && g.CanTo(f, piece.Pos{X: f.pos.X, Y: i}) && dMid == 0 {
				result = append(result, piece.Pos{X: f.pos.X, Y: i})
			} else {
				dMid++
				if g.CanEat(f, nf) {
					result = append(result, nf.pos)
				}
			}
		}
		// 判断左边 ex [0, 1]
		lMid := 0
		llimt := util.MaxInt(f.pos.X-step, 0)
		for i := f.pos.X - 1; i >= llimt; i-- {
			nf = g.Board.GetFetchByPos(piece.Pos{X: i, Y: f.pos.Y})
			if nf == nil && g.CanTo(f, piece.Pos{X: i, Y: f.pos.Y}) && lMid == 0 {
				result = append(result, piece.Pos{X: i, Y: f.pos.Y})
			} else {
				lMid++
				if g.CanEat(f, nf) {
					result = append(result, nf.pos)
				}
			}
		}
		// 判断右边
		rMid := 0
		rlimt := util.MinInt(f.pos.X+step+1, board.WIDTH)
		for i := f.pos.X + 1; i < rlimt; i++ {
			nf = g.Board.GetFetchByPos(piece.Pos{X: i, Y: f.pos.Y})
			if nf == nil && g.CanTo(f, piece.Pos{X: i, Y: f.pos.Y}) && rMid == 0 {
				result = append(result, piece.Pos{X: i, Y: f.pos.Y})
			} else {
				rMid++
				if g.CanEat(f, nf) {
					result = append(result, nf.pos)
				}
			}
		}
	} else if f.CanMoveDiagonal() {
		mayBePos := f.piece.Ways(f.pos, false)
		for _, mayPos := range mayBePos {
			targetF := g.Board.GetFetchByPos(mayPos)
			if targetF != nil && targetF.Side() == f.Side() {
				continue
			}
			if f.piece.Value() == "n" || f.piece.Value() == "N" {
				if !g.CanNotHrun(f.pos, mayPos) {
					result = append(result, mayPos)
				}
			}
			if f.piece.Value() == "B" || f.piece.Value() == "b" {
				// TODO 判断是否填心
				if !g.TainXin(f.pos, mayPos) {
					result = append(result, mayPos)
				}

			}
		}
	}

	return result
}

func (g *Game) AddAteList(to piece.Pos) {
	//toPiece := g.GetFetchByPos(to)
	//if toPiece.GetSide() == piece.Black {
	//	g.RAte = append(g.RAte, toPiece)
	//} else {
	//	g.BAte = append(g.BAte, toPiece)
	//}
}

func (g *Game) MoveLR(f *Fetch, cmd int) {

}

func (g *Game) MoveDown(f *Fetch, cmd int) {

}

func (g *Game) CanEat(f *Fetch, nf *Fetch) bool {
	if nf == nil {
		return false
	}
	if f.pos.X == nf.pos.X && f.pos.Y == nf.pos.Y {
		panic("Eat error!")
	}
	if f.Side() == nf.Side() {
		return false
	}
	mid := g.getMidFetch(f.pos, nf.pos)
	if f.piece.CanEatDirect() {
		if g.CanTo(f, nf.pos) {
			return true
		}
		return false
	} else {
		if len(mid) == 1 {
			return true
		}
	}
	return false
}

// CanTo 判断一个棋子是否可以走到灵位一个地方
func (g *Game) CanTo(fetch *Fetch, pos piece.Pos) bool {
	mid := g.getMidFetch(fetch.pos, pos)
	// 如果走的是直线，那么中间不能有间隔
	if fetch.piece.CanMoveLine() && len(mid) > 0 {
		return false
	}
	// 如果不能过河，但是目的地过河
	isCrossRiver := g.Board.GetPosSide(pos) != fetch.Side()
	if !fetch.piece.CanCrossRiver() && isCrossRiver {
		return false
	}
	return fetch.CanTo(pos, isCrossRiver)
}

func (g *Game) TainXin(pos piece.Pos, pos2 piece.Pos) bool {
	midX := (pos.X + pos2.X) / 2
	midY := (pos.Y + pos2.Y) / 2
	mid := piece.NewPos(midX, midY)
	f := g.Board.GetFetchByPos(mid)
	return f != nil
}
