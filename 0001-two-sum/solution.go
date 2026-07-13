func twoSum(nums []int, target int) []int {
	for i, a := range nums {
		for j, b := range nums {
			if (i != j) && (a+b == target) {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}
