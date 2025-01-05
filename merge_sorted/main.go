package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	fmt.Println("Calling merge")
	nums2Index := 0

	/*
		Need to handle with nums1 is empty
		Need to think about if nums2 is always smaller
		Need to think about when nums1 is always smaller
	*/

	// nothing to add from nums2 or both is empty
	if n == 0 || m == 0 {
		return
	}

	tmpNums := make([]int, m)
	for i := 0; i < m; i++ {
		tmpNums[i] = nums1[i]
	}
	fmt.Println(tmpNums)

	nums1Index := 0
	for i, _ := range nums1 {
		if nums2Index >= n {
			continue
		} else if nums1Index >= m {
			nums1[i] = nums2[nums2Index]
			nums2Index += 1
		} else if tmpNums[nums1Index] > nums2[nums2Index] {
			nums1[i] = nums2[nums2Index]
			nums2Index += 1
		} else {
			nums1[i] = tmpNums[nums1Index]
			nums1Index += 1
		}
	}
	fmt.Println(nums1)
}

func main() {
	fmt.Println("Hello World")
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)
}
