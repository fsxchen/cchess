package piece

// Rook 车
type Rook struct {
	Piece
}

func (r *Rook) CanCrossRiver() bool {
	//TODO implement me
	panic("implement me")
}

func (r *Rook) CanEatDirect() bool {
	return true
}

func (r *Rook) CanMoveDiagonal() bool {
	//TODO implement me
	panic("implement me")
}

func (r *Rook) Ways(pos Pos, b bool) []Pos {
	//TODO implement me
	panic("implement me")
}

func (r *Rook) Step() int {
	return 10
}

func (r *Rook) CanMoveLine() bool {
	return true
}

func NewRook(piece Piece) *Rook {
	return &Rook{Piece: piece}
}

func (r *Rook) CanTo(from, to Pos, isCrossRiver bool) bool {
	// 必须走直线
	if from.X == to.X || from.Y == to.Y {
		return true
	}
	return false
}

func (r *Rook) CanMoveThrough() bool {
	return false
}
