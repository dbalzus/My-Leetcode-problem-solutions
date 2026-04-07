func divideArray(nums []int) bool {
	counter := make(map[int]int)

	for _, num := range nums {
		counter[num]++
	}

	for _, count := range counter {
		if count%2 != 0 {
			return false
		}
	}
	return true
}