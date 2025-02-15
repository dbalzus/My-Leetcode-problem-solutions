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

	return
}
