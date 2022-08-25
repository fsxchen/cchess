package piece

const WIDTH = 9
const HEIGH = 10

// 马
type Knight struct {
	Piece
}

func (k *Knight) CanEatDirect() bool {
	return true
}

func (k *Knight) Ways(pos Pos, b bool) (res []Pos) {
	//TODO implement me
	// 正常来说，做多有8个点可以走
	// 左边
	res = make([]Pos, 0)
	if pos.X+2 < WIDTH && pos.Y+1 < HEIGH {
		res = append(res, NewPos(pos.X+2, pos.Y+1))
	}
	if pos.X+2 < WIDTH && pos.Y-1 >= 0 {
		res = append(res, NewPos(pos.X+2, pos.Y-1))
	}
	if pos.X-2 >= 0 && pos.Y+1 < HEIGH {
		res = append(res, NewPos(pos.X-2, pos.Y+1))
	}
	if pos.X-2 >= 0 && pos.Y-1 >= 0 {
		res = append(res, NewPos(pos.X-2, pos.Y-1))
	}
	if pos.X+1 < WIDTH && pos.Y+2 < HEIGH {
		res = append(res, NewPos(pos.X+1, pos.Y+2))
	}
	if pos.X+1 < WIDTH && pos.Y-2 >= 0 {
		res = append(res, NewPos(pos.X+1, pos.Y-2))
	}
	if pos.X-1 >= 0 && pos.Y+2 < HEIGH {
		res = append(res, NewPos(pos.X-1, pos.Y+2))
	}
	if pos.X-1 >= 0 && pos.Y-2 >= 0 {
		res = append(res, NewPos(pos.X-1, pos.Y-2))
	}

	return
}

func (k *Knight) Step() int {
	//TODO implement me
	return 3
}

func NewKnight(piece Piece) *Knight {
	return &Knight{Piece: piece}
}

func (k *Knight) CanEatDirectly() bool {
	//TODO implement me
	panic("implement me")
}

func (k *Knight) Value() string {
	//TODO implement me
	return k.Piece.val
}

func (k *Knight) IsCrossRiver() bool {
	//TODO implement me
	panic("implement me")
}

func (k *Knight) GetStep() int {
	return 3
}

func (k *Knight) Range() []*Pos {
	//TODO implement me
	panic("implement me")
}

func (k *Knight) CanCrossRiver() bool {
	return true
}

func (k *Knight) CanMoveLine() bool {
	return false
}

func (k *Knight) CanMoveDiagonal() bool {
	return true
}

func (k *Knight) CanCrossPiece() bool {
	//TODO implement me
	return true
}

func (k *Knight) CanTo(from, to Pos, isCrossRiver bool) bool {
	// 只能区分坐标上是否合理
	//
	if from.X == to.X+2 || from.X == to.X-2 {
		if from.Y+1 == to.Y || from.Y-1 == to.Y {
			return true
		}
	} else if from.X == to.X+1 || from.X == to.X-1 {
		if from.Y+2 == to.Y || from.Y-2 == to.Y {
			return true
		}
	}
	return false
}

func (a *Knight) CanMoveThrough() bool {
	return false
}
