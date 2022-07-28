package board

import (
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
const FULL_INIT_FEN = "r"

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
	FENCH_NAME_DICT = map[string]string{
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
	board [10][9]string
}

// TextView 在文本试图下的棋盘，也就是把board对应带文本界面
func (bb *BaseChessBoard) TextView() []string {
	boardStr := _TEXT_BOARD[:]

	y := 0
	// 效率有点低
	for _, line := range bb.board {
		x := 0
		for _, ch := range line {
			if ch != "" && ch != " " {
				fmt.Println(ch)
				pos := Pos{x, y}
				textPos := pos.ToViewPos()
				fmt.Println(textPos.y, textPos.x, "....")
				// 找到第y列，更新其中的字符串
				//boardStr[textPos.y]
				newTextLine := boardStr[textPos.y][:textPos.x-1] + fench_to_txt_name(ch) + boardStr[textPos.y][textPos.x+3:]
				boardStr[textPos.y] = newTextLine
			}
			x++
		}
		y++
	}
	return boardStr
}

func fench_to_txt_name(ch string) string {
	return FENCH_NAME_DICT[ch]
}

// PrintBoard 打印棋盘
func (bb *BaseChessBoard) PrintBoard() {
	//boardStr := _TEXT_BOARD[:]
	fmt.Println(bb.board)
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
				bb.PutFench(ch, Pos{x, y})
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
func (bb *BaseChessBoard) PutFench(ch string, p Pos) {
	bb.board[p.y][p.x] = ch
}

// 清空棋盘
func (bb *BaseChessBoard) clear() {
	for y := 0; y < 10; y++ {
		for x := 0; x < 9; x++ {
			bb.board[y][x] = " "
		}
	}
}

type ChessBoard struct {
	BaseChessBoard
}
