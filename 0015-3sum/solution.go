func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	fmt.Println(nums)
	book := map[[3]int]bool{}
	for i := 0; i < len(nums); i++ {
		l := 0
		r := len(nums) - 1
		for l < r {
			if r == i {
				r -= 1
			}
			if l == i {
				l += 1
			}
			if r == l {
				break
			}
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				new := [3]int{}
				if i >= l && i <= r {
					new = [3]int{nums[l], nums[i], nums[r]}
				} else if i < l {
					new = [3]int{nums[i], nums[l], nums[r]}
				} else {
					new = [3]int{nums[l], nums[r], nums[i]}
				}
				_, ok := book[new]
				if !ok {
					book[new] = true
				}
				r -= 1
				l += 1
			} else if sum > 0 {
				r -= 1
			} else {
				l += 1
			}
		}
	}
	res := [][]int{}
	for key := range book {
		res = append(res, key[:])
	}
	return res
}
