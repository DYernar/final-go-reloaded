package student

func Compare(a, b string) int {
	if a == b {
		return 0
	} else if a < b {
		return -1
	} else {
		return 1
	}
}

func CompareBack(b, a string) int {
	if a == b {
		return 0
	} else if a < b {
		return -1
	} else {
		return 1
	}
}
