package board

import (
	"cchess/cchess/piece"
	"cchess/util"
	"fmt"
	"strconv"
	"strings"
)

/**

9 砗──碼──象──士──将──士──象──碼──砗
  │   │   │   │ ＼│ ／│   │   │   │
8 ├───┼───┼───┼───┼───┼───┼───┼───┤
  │   │　 │   │ ／│ ＼│   │   │   │
7 ├───砲──┼───┼───┼───┼───┼───砲──┤
  │   │　 │　 │　 │   │   │   │   │
6 卒──┼───卒──┼───卒──┼───卒──┼───卒
  │　 │　 │   │   │   │   │   │   │
5 ├───┴───┴───┴───┴───┴───┴───┴───┤
  │　                             │
4 ├───┬───┬───┬───┬───┬───┬───┬───┤
  │　 │　 │   │   │   │　 │　 │　 │
3 兵──┼───兵──┼───兵──┼───兵──┼───兵
  │   │　 │　 │　 │   │   │   │   │
2 ├───炮──┼───┼───┼───┼───┼───炮──┤
  │   │   │   │ ＼│ ／│　 │　 │　 │
1 ├───┼───┼───┼───┼───┼───┼───┼───┤
  │   │　 │   │ ／│ ＼│　 │   │   │
0 车──马──相──仕──帅──仕──相──马──车
*/

//const FULL_INIT_FEN = "rnbakabnr/9/1c5c1/p1p1p1p1p/9/9/P1P1P1P1P/1C5C1/9/RNBAKABNR w - - 0 1"
const FULL_INIT_FEN = "rn"

var _TEXT_BOARD = []string{
	//" 1   2   3   4   5   6   7   8   9",
	"9 ┌───┬───┬───┬───┬───┬───┬───┬───┐ ",
	"  │   │   │   │ ＼│ ／│   │   │   │ ",
	"8 ├───┼───┼───┼───┼───┼───┼───┼───┤ ",
	"  │   │　 │   │ ／│ ＼│   │   │   │ ",
	"7 ├───┼───┼───┼───┼───┼───┼───┼───┤ ",
	"  │   │　 │　 │　 │   │   │   │   │ ",
	"6 ├───┼───┼───┼───┼───┼───┼───┼───┤ ",
	"  │　 │　 │   │   │   │   │   │   │ ",
	"5 ├───┴───┴───┴───┴───┴───┴───┴───┤ ",
	"  │　                             │ ",
	"4 ├───┬───┬───┬───┬───┬───┬───┬───┤ ",
	"  │　 │　 │   │   │   │　 │　 │　 │ ",
	"3 ├───┼───┼───┼───┼───┼───┼───┼───┤ ",
	"  │   │　 │　 │　 │   │   │   │   │ ",
	"2 ├───┼───┼───┼───┼───┼───┼───┼───┤ ",
	"  │   │   │   │ ＼│ ／│　 │　 │　 │ ",
	"1 ├───┼───┼───┼───┼───┼───┼───┼───┤ ",
	"  │   │　 │   │ ／│ ＼│　 │   │   │ ",
	"0 └───┴───┴───┴───┴───┴───┴───┴───┘ ",
	"   ",
	"  a   b   c   d   e   f   g   h   i ",
	"  0   1   2   3   4   5   6   7   8 ",
	//"九  八  七  六  五  四  三  二  一"
}

var (
	NAME_FENCH_DICT = map[string]string{
		"帅": "K",
		"将": "k",
		"仕": "A",
		"士": "a",
		"相": "B",
		"象": "b",
		"马": "n",
		"车": "r",
		"炮": "c",
		"兵": "P",
		"卒": "p",
	}
)

type BaseChessBoard struct {
	// 用来存放棋子的地方
	board [10][9]piece.IPiece
}

// TextView 在文本试图下的棋盘，也就是把board对应带文本界面
func (bb *BaseChessBoard) TextView() []string {
	//boardStr := make([]string, len(_TEXT_BOARD))
	boardStr := util.StringArrayCopy(_TEXT_BOARD)
	y := 0
	// 效率有点低
	for _, line := range bb.board {
		x := 8
		// 这里的line是数组，不必担心不能使用len取长度的问题
		for i := len(line); i > 0; i-- {
			// 从后面往前处理，是因为一个棋子，在图上占用了2个格子，从后往前不用担心数据处理的问题了
			ch := line[i-1]
			if ch != nil {
				pos := piece.NewPos(x, y)
				textPos := pos.ToViewPos()
				// 找到第y列，更新其中的字符串
				//boardStr[textPos.y]
				// 先把找到的行转换成[]rune
				line := []rune(boardStr[textPos.Y])
				line[textPos.X] = []rune(ch.ShowName())[0]
				line = append(line[:textPos.X+1], line[textPos.X+2:]...)
				boardStr[textPos.Y] = string(line)
			}
			x--
		}
		y++
	}
	return boardStr
}

// PrintBoard 打印棋盘
func (bb *BaseChessBoard) PrintBoard() {
	//boardStr := _TEXT_BOARD[:]
	for _, line := range bb.TextView() {
		fmt.Println(line)
	}
}

// 初始化
func (bb *BaseChessBoard) FromFen(fen string) bool {
	chSet := []string{"k", "a", "b", "n", "r", "c", "p"}
	numSet := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

	x := 0
	y := 9
	fenLen := len(fen)
	// 这里是解析初始化的格式文档
	for i := 0; i < fenLen; i++ {
		ch := string(fen[i])

		if ch == " " {
			break
		} else if ch == "/" {
			// /是换行符的含义
			y -= 1
			x = 0
			if y < 0 {
				break
			}
		} else if util.InArray(numSet, ch) {
			ich, _ := strconv.Atoi(ch)
			x += ich
			if x > 8 {
				x = 8
			}
		} else if util.InArray(chSet, strings.ToLower(ch)) {
			if x <= 8 {
				bb.PutFench(ch, piece.Pos{x, y})
				x += 1
			}
		} else {
			return false
		}

	}
	//fmt.Println("初始化了....")
	return true
}

// PutFench 在board中，加入一个棋子
func (bb *BaseChessBoard) PutFench(ch string, p piece.Pos) {
	bb.board[p.Y][p.X] = piece.NewPiece(ch)
}

// 清空棋盘
func (bb *BaseChessBoard) clear() {
	for y := 0; y < 10; y++ {
		for x := 0; x < 9; x++ {
			bb.board[y][x] = nil
		}
	}
}

func (bb *BaseChessBoard) Move(from, to piece.Pos) {
	//  移动一个棋子到另外一个地方
	fromPiece := bb.GetPieceByPos(from)
	if fromPiece == nil {
		fmt.Errorf("error")
		return
	}
	bb.movePiece(from, to)
}

// GetPieceByPos 根据坐标获取棋子
func (bb *BaseChessBoard) GetPieceByPos(pos piece.Pos) piece.IPiece {
	return bb.board[pos.Y][pos.X]
}

// 移动内部的棋子
func (bb *BaseChessBoard) movePiece(from, to piece.Pos) piece.IPiece {
	fench := bb.board[from.Y][from.X]
	bb.board[to.Y][to.X] = fench
	bb.board[from.Y][from.X] = nil
	return fench
}

type ChessBoard struct {
	BaseChessBoard
}
