package util

func InArray(arr []string, ele string) bool {
	for _, arrEle := range arr {
		if ele == arrEle {
			return true
		}
	}
	return false
}

func FindArray(arr []string, ele string) int {
	for index, arrEle := range arr {
		if ele == arrEle {
			return index
		}
	}
	return -1
}

// StringArrayCopy 拷贝字符串切片
func StringArrayCopy(src []string) (dst []string) {
	dst = make([]string, 0)
	for _, srcStr := range src {
		dst = append(dst, srcStr)
	}
	return
}
