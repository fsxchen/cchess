package util

func MinInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func MaxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func SortTwoInt(x, y int) (minNum, maxNum int) {
	if x > y {
		return y, x
	} else {
		return x, y
	}
}
