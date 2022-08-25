package board

//import (
//	"cchess/cchess/piece"
//	"cchess/util"
//	"fmt"
//	"strconv"
//	"strings"
//)
//
const WIDTH = 9
const HEIGH = 10

//
//var _TEXT_BOARD = []string{
//	//" 1   2   3   4   5   6   7   8   9",
//	"9 ┌───┬───┬───┬───┬───┬───┬───┬───┐ ",
//	"  │   │   │   │ ＼│ ／│   │   │   │ ",
//	"8 ├───┼───┼───┼───┼───┼───┼───┼───┤ ",
//	"  │   │　 │   │ ／│ ＼│   │   │   │ ",
//	"7 ├───┼───┼───┼───┼───┼───┼───┼───┤ ",
//	"  │   │　 │　 │　 │   │   │   │   │ ",
//	"6 ├───┼───┼───┼───┼───┼───┼───┼───┤ ",
//	"  │　 │　 │   │   │   │   │   │   │ ",
//	"5 ├───┴───┴───┴───┴───┴───┴───┴───┤ ",
//	"  │　                             │ ",
//	"4 ├───┬───┬───┬───┬───┬───┬───┬───┤ ",
//	"  │　 │　 │   │   │   │　 │　 │　 │ ",
//	"3 ├───┼───┼───┼───┼───┼───┼───┼───┤ ",
//	"  │   │　 │　 │　 │   │   │   │   │ ",
//	"2 ├───┼───┼───┼───┼───┼───┼───┼───┤ ",
//	"  │   │   │   │ ＼│ ／│　 │　 │　 │ ",
//	"1 ├───┼───┼───┼───┼───┼───┼───┼───┤ ",
//	"  │   │　 │   │ ／│ ＼│　 │   │   │ ",
//	"0 └───┴───┴───┴───┴───┴───┴───┴───┘ ",
//	"   ",
//	"  a   b   c   d   e   f   g   h   i ",
//	"  0   1   2   3   4   5   6   7   8 ",
//	//"九  八  七  六  五  四  三  二  一"
//}
//
//type BaseChessBoard struct {
//	// 用来存放棋子的地方,是一个10X9的数组
//	// 每个数组中 ，存放的是棋子结构的地址/nil
//	Board [10][9]piece.IPiece
//}
//
//// TextView 在文本试图下的棋盘，也就是把board对应带文本界面
//func (b *BaseChessBoard) TextView() []string {
//	//boardStr := make([]string, len(_TEXT_BOARD))
//	boardStr := util.StringArrayCopy(_TEXT_BOARD)
//	y := 0
//	// 效率有点低
//	for _, line := range b.Board {
//		x := 8
//		// 这里的line是数组，不必担心不能使用len取长度的问题
//		for i := len(line); i > 0; i-- {
//			// 从后面往前处理，是因为一个棋子，在图上占用了2个格子，从后往前不用担心数据处理的问题了
//			ch := line[i-1]
//			if ch != nil {
//				pos := piece.NewPos(x, y)
//				textPos := pos.ToViewPos()
//				// 找到第y列，更新其中的字符串
//				//boardStr[textPos.y]
//				// 先把找到的行转换成[]rune
//				line := []rune(boardStr[textPos.Y])
//				line[textPos.X] = []rune(ch.ShowName())[0]
//				line = append(line[:textPos.X+1], line[textPos.X+2:]...)
//				boardStr[textPos.Y] = string(line)
//			}
//			x--
//		}
//		y++
//	}
//	return boardStr
//}
//
//// PrintBoard 打印棋盘
//func (b *BaseChessBoard) PrintBoard() {
//	//boardStr := _TEXT_BOARD[:]
//	for _, line := range b.TextView() {
//		fmt.Println(line)
//	}
//}
//
//// 初始化
//func (b *BaseChessBoard) FromFen(fen string) bool {
//	chSet := []string{"k", "a", "b", "n", "r", "c", "p"}
//	numSet := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
//
//	x := 0
//	y := 9
//	fenLen := len(fen)
//	// 这里是解析初始化的格式文档
//	for i := 0; i < fenLen; i++ {
//		ch := string(fen[i])
//
//		if ch == " " {
//			break
//		} else if ch == "/" {
//			// /是换行符的含义
//			y -= 1
//			x = 0
//			if y < 0 {
//				break
//			}
//		} else if util.InArray(numSet, ch) {
//			ich, _ := strconv.Atoi(ch)
//			x += ich
//			if x > 8 {
//				x = 8
//			}
//		} else if util.InArray(chSet, strings.ToLower(ch)) {
//			if x <= 8 {
//				b.PutFench(ch, piece.Pos{x, y})
//				x += 1
//			}
//		} else {
//			return false
//		}
//
//	}
//	//fmt.Println("初始化了....")
//	return true
//}
//
//// PutFench 在board中，加入一个棋子
//func (b *BaseChessBoard) PutFench(ch string, p piece.Pos) {
//	b.Board[p.Y][p.X] = piece.CreatePieceFactory(ch)
//}
//
//// 清空棋盘
//func (b *BaseChessBoard) clear() {
//	for y := 0; y < 10; y++ {
//		for x := 0; x < 9; x++ {
//			b.Board[y][x] = nil
//		}
//	}
//}
//
//// GetPieceByPos 根据坐标获取棋子
//func (b *BaseChessBoard) GetPieceByPos(pos piece.Pos) piece.IPiece {
//	return b.Board[pos.Y][pos.X]
//}
//
//func (b *BaseChessBoard) Move(from, to piece.Pos) error {
//	//  移动一个棋子到另外一个地方
//	// 移动的逻辑如下：
//	// 1. 判断本地是否有
//	// 2. 是否符合走子规范
//	// 3. 目的点情况判断。如果没有：
//	//					1. 路径规范判断
//	//                 如果有：
//	//                  1. 是否吃子
//	//                  2. 路径判断
//	if from.X == to.X && from.Y == to.Y {
//		panic("Error ! 落子无效")
//	}
//
//	fromPiece := b.GetFetchByPos(from)
//	if fromPiece == nil {
//		fmt.Errorf("error")
//		panic("error")
//	}
//	// 判断是否符合棋子的规则
//	//if !fromPiece.CanTo(from, to, false) {
//	//	fmt.Errorf("direct error")
//	//	return
//	//}
//
//	toPiece := b.GetPieceByPos(to)
//
//	if toPiece == nil {
//		// 走子逻辑，走子的逻辑在于，如果是走直线的子，不能有阻挡
//
//		// 构造固定的区域，从左边下到右边上
//		var moveAreaLD piece.Pos
//		var moveAreaRU piece.Pos
//
//		moveAreaLD.X = util.MinInt(from.X, to.X)
//		moveAreaLD.Y = util.MinInt(from.Y, to.Y)
//
//		moveAreaRU.X = util.MaxInt(from.X, to.X)
//		moveAreaRU.Y = util.MaxInt(from.Y, to.Y)
//
//		// 找到区域中间的棋子
//		midPiece := make([]piece.IPiece, 0)
//		if moveAreaRU.X == moveAreaLD.X && moveAreaLD.Y < moveAreaRU.Y {
//			for i := moveAreaLD.Y + 1; i < moveAreaRU.Y; i++ {
//				if nil != b.Board[i][moveAreaLD.X] {
//					midPiece = append(midPiece, b.Board[i][moveAreaLD.X])
//				}
//			}
//		}
//
//		if moveAreaRU.Y == moveAreaLD.Y && moveAreaLD.X < moveAreaRU.X {
//			for i := moveAreaLD.X + 1; i < moveAreaRU.X; i++ {
//				if nil != b.Board[moveAreaLD.Y][i] {
//					midPiece = append(midPiece, b.Board[moveAreaLD.Y][i])
//				}
//			}
//		}
//
//		//if fromPiece.CanMoveLine() && len(midPiece) > 0 {
//		//	fmt.Println("error!")
//		//	return
//		//}
//
//		//if !fromPiece.CanMoveThrough() && len(midPiece) > 0 {
//		//	fmt.Errorf("error")
//		//	return
//		//}
//
//	} else {
//		//if toPiece.(piece.Piece).Side() == fromPiece.(piece.Piece).Side() {
//		//	fmt.Errorf("error")
//		//	return
//		//} else {
//		//	// 吃子的逻辑
//		//	// 除了炮，走子和吃子的逻辑不一样
//		//	if fromPiece.CanEatDirectly() && fromPiece.CanMoveLine() {
//		//		// 能直线行走，直线吃的棋子，车/兵
//		//		if fromPiece.CanTo(from, to, fromPiece.IsCrossRiver()) {
//		//
//		//		}
//		//	} else {
//		//		g.AddAteList(to)
//		//	}
//		//
//		//}
//	}
//
//	b.mvoePiece(from, to)
//	return nil
//}
//
//// 移动内部的棋子
//func (b *BaseChessBoard) mvoePiece(from, to piece.Pos) piece.IPiece {
//	fench := b.Board[from.Y][from.X]
//	b.Board[to.Y][to.X] = fench
//	b.Board[from.Y][from.X] = nil
//	return fench
//}
//
//// 吃掉另外一个子
//func (b *BaseChessBoard) EatPiece(from, to piece.Pos) piece.IPiece {
//	fench := b.Board[from.Y][from.X]
//	b.Board[to.Y][to.X] = fench
//	b.Board[from.Y][from.X] = nil
//	return fench
//}
//
//// 获取棋盘是哪方
//func (bb *BaseChessBoard) getSideByPos(pos piece.Pos) piece.ChessSide {
//	y := pos.Y
//	if y <= 4 {
//		return piece.Red
//	} else {
//	}
//	return piece.Black
//}
//
//type ChessBoard struct {
//	BaseChessBoard
//}
//
//// GetEatWays 所有可以吃的子
//func (bb *BaseChessBoard) GetEatWays(p piece.Pos) []piece.Pos {
//	result := make([]piece.Pos, 0) // 接收结果
//	//curPiece := bb.GetPieceByPos(p)
//	//if curPiece == nil {
//	//	return result
//	//}
//	//if curPiece.CanMoveLine() {
//	//	// 走直线，判断上下左右是否有路
//	//	// 判断上方
//	//	// 例如，输入Y= 【0， 1】
//	//	// 吃范围内遇到的第一个棋子
//	//	for i := p.Y + 1; i < 10; i++ {
//	//		if bb.Board[i][p.X] == nil {
//	//			result = append(result, piece.Pos{p.X, i})
//	//		} else {
//	//			break
//	//		}
//	//	}
//	//}
//	return result
//}
//
//func (b *BaseChessBoard) GetFetchByPos(from piece.Pos) piece.IPiece {
//	x, y := from.ToArrayIndex()
//	return b.Board[x][y]
//}
