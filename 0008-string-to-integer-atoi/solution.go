func myAtoi(s string) int {
	var res int64
	var started bool = false
	var negative bool = false
	for i := 0; i < len(s); i++ {
		digit := s[i]
		if digit == ' ' && !started {
			continue
		}
		if (digit == ' ' && started) || (digit == '+' && started) || (digit == '-' && started) {
			break
		}
		if digit == '+' {
			started = true
			continue
		}
		if digit == '-' {
			started = true
			negative = true
			continue
		}
		if digit >= '0' && digit <= '9' {
			started = true
			res = (res*10 + int64(rune(digit)-rune('0')))
			if !negative && res > math.MaxInt32 {
				return math.MaxInt32
			}
			if negative && -1*res < math.MinInt32 {
				return math.MinInt32
			}
			continue
		}
		break
	}
	if negative {
		res *= -1
	}
	return int(res)
}
