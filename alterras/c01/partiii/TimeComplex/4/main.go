func logarithmic(n int) int {
	var result int = 0
	for n > 1 {
		n /= 2
		result += 1
	}
	return result

}
// log(32) - > 5