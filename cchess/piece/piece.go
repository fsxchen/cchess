package piece

import "cchess/util"

var FENCH_NAME_DICT = map[string]string{
	"K": "帅",
	"k": "将",
	"A": "仕",
	"a": "士",
	"B": "相",
	"b": "象",
	"N": "马",
	"n": "馬",
	"R": "车",
	"r": "車",
	"C": "炮",
	"c": "砲",
	"P": "兵",
	"p": "卒",
}

var NAME_FETCH_DICT = map[string]string{
	"帅": "K",
	"将": "k",
	"仕": "A",
	"士": "a",
	"相": "B",
	"象": "b",
	"马": "N",
	"馬": "n",
	"车": "R",
	"車": "r",
	"炮": "C",
	"砲": "c",
	"兵": "P",
	"卒": "p",
}

var PIECE_LIST = []string{"帅", "将", "仕", "士", "相", "象", "马", "馬”", "车", "車", "炮", "砲", "兵", "卒"}
var RED_PIECE_LIST = []string{"K", "A", "B", "N", "R", "C", "P"}
var BLACK_PIECE_LIST = []string{"k", "a", "b", "n", "r", "c", "p"}

type (
	IPiece interface {
		CanTo(Pos, Pos, bool) bool
		ShowName() string
		Value() string
		Side() ChessSide
		CanMoveLine() bool
		CanMoveDiagonal() bool
		CanEatDirect() bool
		CanCrossRiver() bool
		Step() int
		Ways(Pos, bool) []Pos
	}
)

// Area 因为所有涉及的区域都为举行区域
type Area struct {
	LD Pos
	RU Pos
}

type ChessSide byte

const (
	Red ChessSide = iota
	Black
)

type Piece struct {
	val  string    // 值
	side ChessSide // 那一方

	// 1. 判断棋子是否有固定范围
	// 2. 是否可以过河
	// 3. 是否可以走直线
	// 4. 是否可以走斜线
	// 5. 是否有特殊规则
	// 6. 是否可以穿越
}

func (p *Piece) Side() ChessSide {
	return p.side
}

func (p *Piece) Value() string {
	return p.val
}

func NewPiece(val string) *Piece {
	var side ChessSide
	if util.InArray(RED_PIECE_LIST, val) {
		side = Red
	} else if util.InArray(BLACK_PIECE_LIST, val) {
		side = Black
	} else {
		panic("不可用的初始化值")
	}
	return &Piece{val: val, side: side}
}

func (p *Piece) ShowName() string {
	return FENCH_NAME_DICT[p.val]
}

func CreatePieceFactory(val string) IPiece {
	piece := NewPiece(val)
	switch val {
	case "a", "A":
		return NewAdvisor(*piece)
	case "b", "B":
		return NewBishop(*piece)
	case "c", "C":
		return NewCannon(*piece)
	case "k", "K":
		return NewKing(*piece)
	case "n", "N":
		return NewKnight(*piece)
	case "p", "P":
		return NewPawn(*piece)
	case "r", "R":
		return NewRook(*piece)
	default:
		panic("错误字符")
	}
}
