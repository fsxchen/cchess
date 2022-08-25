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

// ToArrayIndex 坐标和数组索引是有区别的
func (p *Pos) ToArrayIndex() (x, y int) {
	x = p.Y
	y = p.X
	return
}

func (p *Pos) Equal(pos Pos) bool {
	if p.X == pos.X && p.Y == pos.Y {
		return true
	}
	return false
}

func IndexToPos(x, y int) Pos {
	return Pos{y, x}
}

func (p *Pos) nextU() *Pos {
	if p.Y+1 >= HEIGH {
		return nil
	}
	n := NewPos(p.X, p.Y+1)
	return &n
}

func (p *Pos) nextD() *Pos {
	if p.Y-1 < 0 {
		return nil
	}
	n := NewPos(p.X, p.Y-1)
	return &n
}

func (p *Pos) nextL() *Pos {
	if p.X-1 < 0 {
		return nil
	}
	n := NewPos(p.X-1, p.Y)
	return &n
}

func (p *Pos) nextR() *Pos {
	if p.X+1 >= WIDTH {
		return nil
	}
	n := NewPos(p.X-1, p.Y)
	return &n
}
