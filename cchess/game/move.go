package game

import (
	"cchess/cchess/piece"
	"fmt"
)

type (
	Mover interface {
		MoveUp(int) error
		MoveDown(int) error
		MoveLeft(int) error
		MoveRight(int) error
	}
)

type Fetch struct {
	piece piece.IPiece
	pos   piece.Pos
}

func NewFetch(piece piece.IPiece, pos piece.Pos) *Fetch {
	return &Fetch{piece: piece, pos: pos}
}

func (f *Fetch) Side() piece.ChessSide {
	return f.piece.Side()
}

func (f Fetch) MoveUp(step int) error {
	//f.piece.CanTo()
	from := f.pos
	var toPos piece.Pos
	if f.piece.CanMoveLine() {
		if f.piece.Side() == piece.Red {
			// 红色方往上
			toPos = piece.Pos{X: from.X, Y: from.Y + step}
		}
		if f.piece.Side() == piece.Black {
			// 红色方往上
			toPos = piece.Pos{X: from.X, Y: from.Y - step}
		}
	}
	f.Move(toPos)
	fmt.Println(from)
	//if f.piece.CanTo()
	return nil
}

func (f Fetch) MoveDown(step int) error {
	return nil
}

func (f Fetch) MoveRight(step int) error {
	return nil
}

func (f Fetch) MoveLeft(step int) error {
	return nil
}

func (f Fetch) Move(to piece.Pos) {
	if f.piece.CanTo(f.pos, to, false) {

	}
}

func (f *Fetch) CanTo(to piece.Pos, isCrossRiver bool) bool {
	return f.piece.CanTo(f.pos, to, isCrossRiver)
}

func (f *Fetch) Ways() []piece.Pos {
	return f.piece.Ways(f.pos, false)
}

func (f *Fetch) CanMoveDiagonal() bool {
	return f.piece.CanMoveDiagonal()
}
