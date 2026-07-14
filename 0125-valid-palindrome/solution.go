func isPalindrome(s string) bool {
	l := 0
	r := len(s) - 1
	isValidLetter := func(r byte) bool {
		if 'a' <= r && r <= 'z' {
			return true
		}
		if 'A' <= r && r <= 'Z' {
			return true
		}
		if '0' <= r && r <= '9' {
			return true
		}
		return false
	}
	downLetter := func(r byte) byte {
		if 'A' <= r && r <= 'Z' {
			return 'a' + (r - 'A')
		}
		return r
	}
	for l < r {
		for !isValidLetter(s[l]) && l < len(s)-1 {
			l += 1
		}
		for !isValidLetter(s[r]) && r > 0 {
			r -= 1
		}
		if r == 0 && l == len(s)-1 {
			return true
		}
		if downLetter(s[l]) != downLetter(s[r]) {
			return false
		}
		l += 1
		r -= 1
	}
	return true
}
