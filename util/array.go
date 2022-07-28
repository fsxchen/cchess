package util

func InArray(arr []string, ele string) bool {
	for _, arrEle := range arr {
		if ele == arrEle {
			return true
		}
	}
	return false
}
