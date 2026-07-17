func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	vol := (r - l) * min(height[r], height[l])
	for l <= r {
		if (r-l)*min(height[r], height[l]) > vol {
			vol = (r - l) * min(height[r], height[l])
		}
		if height[l] > height[r] {
			r--
		} else if height[l] < height[r] {
			l++
		} else {
			r--
			l++
		}
	}
	return vol
}
