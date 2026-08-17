func lengthOfLongestSubstring(s string) int {
	l, r := 0, 0
	res := 0
	for i := range s {
		for strings.Contains(s[l:r], string(s[i])) {
			l += 1
		}
		r += 1
		if r-l > res {
			res = r - l
		}
	}
	return res
}
