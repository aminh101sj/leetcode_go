package main

import "fmt"

func removeDuplicates(nums []int) int {
	l := 0
	for r := 1; r < len(nums); r++ {
		if nums[l] != nums[r] {
			l++
			nums[l] = nums[r]
		}
	}
	return l + 1
}

func main() {
	nums := []int{1, 1, 2}
	k := removeDuplicates(nums)
	fmt.Println(k, nums)
}
