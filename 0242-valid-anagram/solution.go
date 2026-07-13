func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	s_book := map[rune]int{}
	t_book := map[rune]int{}
	for _, v := range s {
		s_book[v] += 1
	}
	for _, v := range t {
		t_book[v] += 1
	}
	for k, v := range s_book {
		if t_book[k] != v {
			return false
		}
	}
	return true
}
