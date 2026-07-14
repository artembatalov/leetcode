func topKFrequent(nums []int, k int) []int {
	type number struct {
		i   int
		num int
	}
	book := make([]number, 25000)
	for _, num := range nums {
		book[num+10000].i += 1
		book[num+10000].num = num
	}
	sort.Slice(book, func(i, j int) bool {
		return book[i].i > book[j].i
	})
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = book[i].num
	}
	return res
}
