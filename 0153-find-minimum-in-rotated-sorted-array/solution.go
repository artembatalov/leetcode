func findMin(nums []int) int {
	l := 0
	r := len(nums) - 1
	min := nums[0]
	for l <= r {
		mid := (r + l) / 2
		if nums[mid] < min {
			min = nums[mid]
		}
		if nums[l] <= nums[mid] {
			if nums[l] < min {
				min = nums[l]
			}
			l = mid + 1
		} else {
			if nums[mid+1] < min {
				min = nums[mid+1]
			}
			r = mid - 1
		}
	}
	return min
}
