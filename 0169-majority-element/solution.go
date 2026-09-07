func majorityElement(nums []int) int {
	major := nums[0]
	cnt := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == major {
			cnt += 1
		} else {
			cnt -= 1
			if cnt == 0 {
				major = nums[i]
				cnt = 1
			}
		}
	}
	return major
}
