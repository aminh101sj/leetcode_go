package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums2Index := 0

	// nothing to add from nums2 or both is empty
	if n == 0 {
		return
	}

	tmpNums := make([]int, m)
	for i := 0; i < m; i++ {
		tmpNums[i] = nums1[i]
	}

	nums1Index := 0
	for i := 0; i < len(nums1); i++ {
		if nums2Index >= n && nums1Index >= m {
			continue
		} else if nums2Index >= n {
			nums1[i] = tmpNums[nums1Index]
			nums1Index += 1
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
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)
	nums1 = []int{0}
	nums2 = []int{1}
	merge(nums1, 0, nums2, 1)
	fmt.Println(nums1)
	nums1 = []int{2, 0}
	nums2 = []int{1}
	merge(nums1, 1, nums2, 1)
	fmt.Println(nums1)
}
