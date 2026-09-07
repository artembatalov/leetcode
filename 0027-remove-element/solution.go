func removeElement(nums []int, val int) int {
	cnt := 0
	pointer := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			cnt += 1
		} else {
			nums[pointer] = nums[i]
			pointer += 1
		}
	}
	fmt.Println(nums)
	return len(nums) - cnt
}
