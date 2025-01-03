package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	fmt.Println("Calling merge")
}

func main() {
	fmt.Println("Hello World")
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}
