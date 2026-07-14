func groupAnagrams(strs []string) [][]string {
	book := map[string][]string{}
	for _, s := range strs {
		key := []rune(s)
		slices.Sort(key)
		if _, ok := book[string(key)]; ok {
			book[string(key)] = append(book[string(key)], s)
		} else {
			book[string(key)] = []string{s}
		}
	}
	res := [][]string{}
	for _, val := range book {
		res = append(res, val)
	}
	return res
}
