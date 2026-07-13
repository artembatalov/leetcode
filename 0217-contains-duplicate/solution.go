func containsDuplicate(nums []int) bool {
	book := map[int]int{}
	for _, x := range nums {
		book[x] += 1
	}
	for _, value := range book {
		if value > 1 {
			return true
		}
	}
	return false
}
