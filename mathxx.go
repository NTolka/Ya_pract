package yapract

func Add(a, b int) int {
	return a + b
}

func Contains(x []string, a string) bool {
	for _, v := range x {
		if a == v {
			return true
		}
	}
	return false
}
