package board

type Pos struct {
	x, y int
}

func (p *Pos) ToViewPos() Pos {
	return Pos{(4 * p.x) + 2, (9 - p.y) * 2}
}
