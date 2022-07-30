package piece

type Pos struct {
	X, Y int
}

func NewPos(x, y int) Pos {
	return Pos{x, y}
}

func (p *Pos) ToViewPos() Pos {
	return Pos{(4 * p.X) + 2, (9 - p.Y) * 2}
}
