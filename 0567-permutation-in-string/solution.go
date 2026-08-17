func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	var val int
	for i := 0; i < len(s1); i++ {
		val += int(s1[i])
	}
	var cur int
	for i := 0; i < len(s1); i++ {
		cur += int(s2[i])
	}
	for i := 0; i+len(s1) <= len(s2); i++ {
		if i != 0 {
			cur += int(s2[len(s1)+i-1])
			cur -= int(s2[i-1])
		}
		if val == cur {
			ch1 := []rune(s1)
			ch2 := []rune(s2[i:(len(s1) + i)])
			slices.Sort(ch1)
			slices.Sort(ch2)
			if slices.Compare(ch1, ch2) == 0 {
				return true
			}
		}
	}
	return false
}
