package piece

type IPiece interface {
	ShowName() string
}

type Piece struct {
	val  string // 值
	side byte   // 那一方
}

var FENCH_NAME_DICT = map[string]string{
	"K": "帅",
	"k": "将",
	"A": "仕",
	"a": "士",
	"B": "相",
	"b": "象",
	"N": "马",
	"n": "马",
	"R": "车",
	"r": "车",
	"C": "炮",
	"c": "炮",
	"P": "兵",
	"p": "卒",
}

func NewPiece(val string) IPiece {
	// 创建不同的棋子
	switch val {
	case "r":
		return &Rook{Piece{val: val, side: 'r'}}
	}
	return &Piece{val: val, side: 'r'}
}

func (p *Piece) ShowName() string {
	return FENCH_NAME_DICT[p.val]
}
