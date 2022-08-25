package piece

// Advisor 士
type Advisor struct {
	Piece
}

func (a *Advisor) CanCrossRiver() bool {
	//TODO implement me
	panic("implement me")
}

func (a *Advisor) CanEatDirect() bool {
	return true
}

func (a *Advisor) CanMoveDiagonal() bool {
	//TODO implement me
	panic("implement me")
}

func (a *Advisor) Ways(pos Pos, b bool) []Pos {
	//TODO implement me
	panic("implement me")
}

func (a *Advisor) Step() int {
	//TODO implement me
	panic("implement me")
}

func (a *Advisor) CanMoveLine() bool {
	return false
}

func NewAdvisor(piece Piece) *Advisor {
	return &Advisor{Piece: piece}
}

func (a *Advisor) CanTo(from, to Pos, isCrossRiver bool) bool {
	return true
}

func (a *Advisor) ShowName() string {
	return a.Piece.ShowName()
}

func (a *Advisor) GetValue() string {
	return a.Piece.val
}

//func (a *Advisor) CanEatDirectly() bool {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (a *Advisor) Value() string {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (a *Advisor) IsCrossRiver() bool {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (a *Advisor) GetStep() int {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (a *Advisor) Range() []*Pos {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (a *Advisor) CanCrossRiver() bool {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (a *Advisor) CanMoveLine() bool {
//	//TODO implement me
//	panic("implement me")
//}
//
