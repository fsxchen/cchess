package game

import (
	"cchess/cchess/board"
	"cchess/cchess/piece"
	"cchess/util"
)

var CMD_NUM_RED_LIST = []string{"一", "二", "三", "四", "五", "六", "七", "八", "九"}
var CMD_NUM_BLACK_LIST = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

var RED_PIECE_LIST = []string{"帅", "仕", "相", "马", "车", "炮", "兵"}
var BLACK_PIECE_LIST = []string{"将", "士", "象", "馬", "車", "砲", "卒"}
var CMD_ACTION_LIST = []string{"进", "退", "平"}
var CMD_ACTION_STR_LIST = []string{"U", "D", "LR"}

// 移动命令解析，解析棋谱中的
// "兵七进一"
// 但是棋谱中有前中后，比如前卒进1这种，先不考虑了

func getActionStr(action string) (actionStr string) {
	index := util.FindArray(CMD_ACTION_LIST, action)
	if index == -1 {
		panic("错误的action")
	}
	return CMD_ACTION_STR_LIST[index]
}

func McmdParse(cmd string) (val string, fromCol int, direct string, toCmd int) {
	cmdList := []rune(cmd)
	if len(cmdList) != 4 {
		panic("不可用的命令")
	}
	// 其实这个指令是在不同的列上找到相应的棋子
	//fmt.Println(cmd, cmdList)
	fromPieceName := cmdList[0]
	fromPos := cmdList[1]
	directName := cmdList[2]
	toPos := cmdList[3]
	if util.InArray(RED_PIECE_LIST, string(fromPieceName)) &&
		util.InArray(CMD_NUM_RED_LIST, string(fromPos)) &&
		util.InArray(CMD_NUM_RED_LIST, string(toPos)) {
		// 处理标准的红方走的情况
		fromCol = board.WIDTH - 1 - util.FindArray(CMD_NUM_RED_LIST, string(fromPos))
		val = piece.NAME_FETCH_DICT[string(fromPieceName)]
		direct = getActionStr(string(directName))
		toCmd = util.FindArray(CMD_NUM_RED_LIST, string(toPos)) + 1
	}

	//if !util.InArray(CMD_ACTION_LIST, string(direct)) {
	//	panic("不可用")
	//}
	//
	//// 因为坐标系和棋谱的坐标系是反的
	//col := WIDTH - 1 - util.FindArray(CMD_NUM_RED_LIST, string(fromPos))
	//val := NAME_FENCH_DICT[string(fromPieceName)]

	// 上面步骤解析了两个操作
	return
}
