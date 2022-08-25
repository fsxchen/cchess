package piece

// 相
type Bishop struct {
	Piece
}

func (b *Bishop) CanEatDirect() bool {
	return true
}

func (b *Bishop) Ways(pos Pos, isCrossRiver bool) (result []Pos) {
	result = make([]Pos, 0)
	if isCrossRiver {
		return
	}
	if b.Side() == Black {
		if pos.Y-2 > 4 && pos.X-2 >= 0 {
			result = append(result, NewPos(pos.X-2, pos.Y-2))
		}
		if pos.Y-2 > 4 && pos.X+2 < WIDTH {
			result = append(result, NewPos(pos.X+2, pos.Y-2))
		}
		if pos.Y+2 > 4 && pos.Y+2 < HEIGH && pos.X-2 > 0 {
			result = append(result, NewPos(pos.X-2, pos.Y+2))
		}
		if pos.Y+2 > 4 && pos.Y+2 < HEIGH && pos.X+2 < HEIGH {
			result = append(result, NewPos(pos.X+2, pos.Y+2))
		}

	} else if b.Side() == Red {
		if pos.Y-2 >= 0 && pos.Y-2 < 5 && pos.X-2 >= 0 {
			result = append(result, NewPos(pos.X-2, pos.Y-2))
		}
		if pos.Y-2 >= 0 && pos.Y-2 < 5 && pos.X+2 < WIDTH {
			result = append(result, NewPos(pos.X+2, pos.Y-2))
		}
		if pos.Y+2 < 5 && pos.Y+2 < HEIGH && pos.X-2 > 0 {
			result = append(result, NewPos(pos.X-2, pos.Y+2))
		}
		if pos.Y+2 < 5 && pos.Y+2 < HEIGH && pos.X+2 < HEIGH {
			result = append(result, NewPos(pos.X+2, pos.Y+2))
		}
	}
	return
}

func (b *Bishop) Step() int {
	return 4
}

func NewBishop(piece Piece) *Bishop {
	return &Bishop{Piece: piece}
}

func (b *Bishop) Value() string {
	return b.Piece.Value()
}

func (b *Bishop) GetStep() int {
	return 4
}

func (b *Bishop) Range() []*Pos {
	//TODO implement me
	panic("implement me")
}

func (b *Bishop) CanCrossRiver() bool {
	return false
}

func (b *Bishop) CanMoveLine() bool {
	return false
}

func (b *Bishop) CanMoveDiagonal() bool {
	return true
}

func (b *Bishop) CanCrossPiece() bool {
	//TODO implement me
	panic("implement me")
}

func (b *Bishop) CanTo(from, to Pos, isCrossRiver bool) bool {
	if b.Value() == "b" && to.Y > 4 {
		if to.Y == from.Y+2 || to.Y == from.Y-2 {
			if to.X == from.X+2 || to.X == from.X-2 {
				return true
			}
		}
	} else if b.Value() == "B" && to.Y < 5 {
		if to.Y == from.Y+2 || to.Y == from.Y-2 {
			if to.X == from.X+2 || to.X == from.X-2 {
				return true
			}
		}
	}
	return false
}
