func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		res[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		res[i] *= nums[i-1] * res[i-1]
	}
	suffix := 1
	for i := len(nums) - 2; i >= 0; i-- {
		suffix *= nums[i+1]
		res[i] *= suffix
	}
	return res
}
