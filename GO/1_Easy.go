func twoSum(nums []int, target int) []int {
	numberIndexMap := make(map[int]int)

	for currentIndex, currentNumber := range nums {
		numberNeeded := target - currentNumber

		if neededIndex, found := numberIndexMap[numberNeeded]; found {
			return []int{neededIndex, currentIndex}
		}

		numberIndexMap[currentNumber] = currentIndex
	}

	return nil
}