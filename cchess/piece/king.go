package piece

// 王
type King struct {
	Piece
}

func (k *King) CanEatDirect() bool {
	return true
}

func (k *King) Ways(pos Pos, b bool) []Pos {
	//TODO implement me
	panic("implement me")
}

func (k *King) Step() int {
	//TODO implement me
	return 1
}

func NewKing(piece Piece) *King {
	return &King{Piece: piece}
}

func (k *King) CanEatDirectly() bool {
	//TODO implement me
	panic("implement me")
}

func (k *King) Range() []*Pos {
	//TODO implement me
	panic("implement me")
}

func (k *King) CanCrossRiver() bool {
	return false
}

func (k *King) CanMoveLine() bool {
	return true
}

func (k *King) CanMoveDiagonal() bool {
	//TODO implement me
	panic("implement me")
}

func (k *King) CanCrossPiece() bool {
	//TODO implement me
	panic("implement me")
}

func (k *King) GetStep() int {
	//TODO implement me
	panic("implement me")
}

func (k *King) IsCrossRiver() bool {
	//TODO implement me
	panic("implement me")
}

func (k *King) CanTo(from, to Pos, isCrossRiver bool) bool {
	return true
}

func (k *King) CanMoveThrough() bool {
	return false
}
