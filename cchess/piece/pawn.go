package piece

// Pawn 兵
type Pawn struct {
	Piece
}

func (p *Pawn) CanEatDirect() bool {
	return true
}

func (p *Pawn) Ways(pos Pos, b bool) []Pos {
	//TODO implement me
	panic("implement me")
}

func (p *Pawn) Step() int {
	return 1
}

func NewPawn(piece Piece) *Pawn {
	return &Pawn{Piece: piece}
}

func (p *Pawn) CanEatDirectly() bool {
	//TODO implement me
	panic("implement me")
}

func (p *Pawn) Value() string {
	//TODO implement me
	return p.val
}

func (p *Pawn) IsCrossRiver() bool {
	//TODO implement me
	panic("implement me")
}

func (p *Pawn) GetStep() int {
	//TODO implement me
	return 1
}

func (p *Pawn) Range() []*Pos {
	//TODO implement me
	panic("implement me")
}

func (p *Pawn) CanCrossRiver() bool {
	return true
}

func (p *Pawn) CanMoveLine() bool {
	//TODO implement me
	return true
}

func (p *Pawn) CanMoveDiagonal() bool {
	//TODO implement me
	panic("implement me")
}

func (p *Pawn) CanCrossPiece() bool {
	return false
}

func (p *Pawn) CanTo(from, to Pos, isCrossRiver bool) bool {
	// 如果没过河，只能往前
	if p.Piece.Side() == Red && to.Y == from.Y+1 && to.X == from.X {
		return true
	}
	if p.Piece.Side() == Black && to.Y == from.Y-1 && to.X == from.X {
		return true
	}
	if isCrossRiver {
		// 已经过河了
		if to.Y == from.Y {
			if to.X == from.X+1 || to.X == from.X-1 {
				return true
			}
		}
	}
	return false
}

func (p *Pawn) CanMoveThrough() bool {
	return false
}
