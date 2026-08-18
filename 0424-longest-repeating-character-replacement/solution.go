func isValid(window [26]int, k int, length int) bool {
	key_number := 0
	for j := 0; j < 26; j++ {
		if window[j] > key_number {
			key_number = window[j]
		}
	}
	return (k+key_number >= length)
}

func characterReplacement(s string, k int) int {
	window := [26]int{}
	l, r, max_lenght, length := 0, 0, 0, 0
	for i := 0; i < len(s); i++ {
		r += 1
		length += 1
		window[int(s[i])-int('A')] += 1
		for !isValid(window, k, length) {
			window[int(s[l])-int('A')] -= 1
			length -= 1
			l += 1
		}
		if length > max_lenght {
			max_lenght = length
		}
	}
	return max_lenght
}
