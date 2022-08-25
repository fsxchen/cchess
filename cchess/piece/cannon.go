package piece

// Cannon 炮
type Cannon struct {
	Piece
}

func (c *Cannon) CanEatDirect() bool {
	return false
}

func (c *Cannon) Ways(pos Pos, b bool) []Pos {
	//TODO implement me
	panic("implement me")
}

func (c *Cannon) Step() int {
	return 10
}

func NewCannon(piece Piece) *Cannon {
	return &Cannon{Piece: piece}
}

func (c *Cannon) CanEatDirectly() bool {
	return false
}

func (c *Cannon) IsCrossRiver() bool {
	//TODO implement me
	panic("implement me")
}

func (c *Cannon) Range() []*Pos {
	//TODO implement me
	panic("implement me")
}

func (c *Cannon) CanCrossRiver() bool {
	//TODO implement me
	panic("implement me")
}

func (c *Cannon) CanMoveLine() bool {
	//TODO implement me
	return true
}

func (c *Cannon) CanMoveDiagonal() bool {
	//TODO implement me
	panic("implement me")
}

func (c *Cannon) CanCrossPiece() bool {
	//TODO implement me
	panic("implement me")
}

func (c *Cannon) CanTo(from, to Pos, isCrossRiver bool) bool {
	return true
}

func (c *Cannon) CanMoveThrough() bool {
	return false
}
