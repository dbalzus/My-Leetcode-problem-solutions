package main

import (
	"fmt"
)

func main() {
	//var nums = []int{2, 11, 10, 1, 3}
	var nums = []int{1, 1, 2, 4, 9}
	//var nums = []int{999999999, 999999999, 999999999}
	//var k int = 10
	var k int = 20
	//var k int = 1000000000

	i := minOperations(nums, k) //doesn't work properly

	fmt.Printf("iterations.: %d\n", i)
}

func minOperations(nums []int, k int) int {
	var x int
	var y int
	var j int
	var xUsed = false
	var yUsed = false
	var iterations int
	var tempSlice []int
	for {
		if iterations >= 1 {
			nums = nil
			nums = append(nums, tempSlice...)
			tempSlice = nil
			xUsed = false
			yUsed = false
			x = 0
			y = 0
			j = 0
		}

		if len(nums) < 2 {
			break
		}

		for i := 0; i < len(nums); i++ {
			if nums[i] < k {
				if xUsed == false {
					x = nums[i]
					xUsed = true
				} else if yUsed == false {
					y = nums[i]
					yUsed = true
				} else if nums[i] < x || nums[i] < y {
					switch max(x, y) {
					case x:
						tempSlice = append(tempSlice, x)
						j++
						x = nums[i]

					case y:
						tempSlice = append(tempSlice, y)
						j++
						y = nums[i]
					}
				} else {
					tempSlice = append(tempSlice, nums[i])
					j++
				}

			} else {
				tempSlice = append(tempSlice, nums[i])
				j++
			}
		}
		if j >= len(nums) {
			break
		}

		if x == 0 || y == 0 {
			iterations++
			break
		}
		tempSlice = append(tempSlice, min(x, y)*2+max(x, y))
		iterations++
	}
	return iterations
}
