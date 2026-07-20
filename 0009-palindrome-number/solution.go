func isPalindrome(x int) bool {
	x_str := strconv.Itoa(x)
	for i := 0; i < len(x_str)/2; i++ {
		if x_str[i] != x_str[len(x_str)-1-i] {
			return false
		}
	}
	return true
}
