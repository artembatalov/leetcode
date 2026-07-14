func twoSum(numbers []int, target int) []int {
	l := 0
	r := len(numbers) - 1
	for l < len(numbers)-1 && r > 0 {
		if numbers[l]+numbers[r] == target {
			return []int{l + 1, r + 1}
		} else if numbers[l]+numbers[r] > target {
			r -= 1
		} else {
			l += 1
		}
	}
	return []int{-1, -1}
}
