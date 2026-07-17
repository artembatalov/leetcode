func maxProfit(prices []int) int {
	left := 0
	right := 0
	max := 0
	for right < len(prices) {
		if prices[left] > prices[right] {
			left = right
			right++
		} else {
			if prices[right]-prices[left] > max {
				max = prices[right] - prices[left]
			}
			right++
		}
	}
	return max
}
