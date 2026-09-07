func longestCommonPrefix(strs []string) string {
	minlen := len(strs[0])
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < minlen {
			minlen = len(strs[i])
		}
	}
	for i := 0; i < minlen; i++ {
		for j := 0; j < len(strs); j++ {
			if strs[j][i] != strs[0][i] {
				return strs[0][0:i]
			}
		}
	}
	return strs[0][0:minlen]
}

//func main() {
//	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}) == "")
//	fmt.Println(longestCommonPrefix([]string{"dog", "dog", "dog"}) == "dog")
//	fmt.Println(longestCommonPrefix([]string{"dogy", "dog", "dog"}) == "dog")
//	fmt.Println(longestCommonPrefix([]string{"dog", "dog", ""}) == "")
//	fmt.Println(longestCommonPrefix([]string{"paparazi"}) == "paparazi")
//}

