package main

func main() {

	maximumCount([]int{2, 3, 5, 7, 11, 13})
	return
}

func maximumCount(nums []int) int {
	var pos = 0
	var neg = 0

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			pos += 1
		}
		if nums[i] < 0 {
			neg += 1
		}
	}

	return max(neg, pos)
}
