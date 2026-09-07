func getConcatenation(nums []int) []int {
	res := make([]int, len(nums)*2)
	n := len(nums)
	for i := 0; i < n; i++ {
		res[i] = nums[i]
		res[i+n] = nums[i]
	}
	return res
}
