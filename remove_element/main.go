package main

import "fmt"

// Function will return the number of elements not equal to val
func removeElement(nums []int, val int) int {
	i := 0
	for _, v := range nums {
		if v != val {
			nums[i] = v
			i++
		}
	}
	return i
}

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	// expects [2,2,_,_]
	k := removeElement(nums, val)
	fmt.Println(k, nums)

}
