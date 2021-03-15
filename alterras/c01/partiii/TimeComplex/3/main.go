func liner2(n int, m int) int {
	var result int = 0
	for i := 0; i < n; i++ {
		result += 1
	}
	for j := 0; j < n; j++ {
		result += 1
	}
	return result
}