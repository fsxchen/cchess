package game

import (
	"cchess/cchess/piece"
	"cchess/util"
	"fmt"
	"strconv"
	"strings"
)

type FetchBoard struct {
	board [10][9]*Fetch
}

func (b *FetchBoard) SetFetch(fetch Fetch) {
	x, y := fetch.pos.ToArrayIndex()
	b.board[x][y] = &fetch
}

// 根据某行/某列获取一个Fetch
func (b *FetchBoard) GetFetchByIndex(row int, col int) *Fetch {
	return b.board[row][col]
}

func (b *FetchBoard) GetFetchByPos(pos piece.Pos) *Fetch {
	return b.board[pos.Y][pos.X]
}

// 把一个棋子从位置上的一个地方移动到另一个地方
func (b *FetchBoard) Move(from piece.Pos, to piece.Pos) {
	toX, toY := to.ToArrayIndex()
	fromX, fromY := from.ToArrayIndex()
	b.board[toX][toY] = b.board[fromX][fromY]
	b.board[fromX][fromY] = nil
}

func (b *FetchBoard) TextView() []string {
	//boardStr := make([]string, len(_TEXT_BOARD))
	boardStr := util.StringArrayCopy(_TEXT_BOARD)
	y := 0
	// 效率有点低
	for _, line := range b.board {
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
				line[textPos.X] = []rune(ch.piece.ShowName())[0]
				line = append(line[:textPos.X+1], line[textPos.X+2:]...)
				boardStr[textPos.Y] = string(line)
			}
			x--
		}
		y++
	}
	return boardStr
}

func (b *FetchBoard) PrintBoard() {
	for _, line := range b.TextView() {
		fmt.Println(line)
	}
}

func (b *FetchBoard) FromFen(fen string) bool {
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
				b.PutFench(ch, piece.Pos{x, y})
				x += 1
			}
		} else {
			return false
		}
	}
	return true
}

func (b *FetchBoard) PutFench(ch string, pos piece.Pos) {
	newPiece := piece.CreatePieceFactory(ch)

	b.SetFetch(*NewFetch(newPiece, pos))
}

// GetPosSide 获取一个坐标是位于那一边
func (b *FetchBoard) GetPosSide(pos piece.Pos) piece.ChessSide {
	if pos.Y > 4 {
		return piece.Black
	} else {
		return piece.Red
	}
}

func NewFetchBoard() *FetchBoard {
	var res = [10][9]*Fetch{}
	return &FetchBoard{res}
}
