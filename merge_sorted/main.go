package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	fmt.Println("Calling merge")
	s2 := 0
	tmp := 0

	/*
		Need to handle with nums1 is empty
		Need to think about if nums2 is always smaller
		Need to think about when nums1 is always smaller
	*/

	if n == 0 {
		return
	}

	if m == 0 {
		nums1 = nums2
		return
	}

	for i, val := range nums1 {
		fmt.Println(i, val)

		if tmp > nums2[s2] && s2 < n {
			if val > nums2[s2] {
				tmp = nums1[i]
				nums1[i] = nums2[s2]
				nums2[s2] = tmp
			}
		} else if s2 < n {
			if val > nums2[s2] {
				tmp = val
				nums[i] = nums2[s2]
				s2 += 1
			} else if val > tmp && tmp != 0 {

			}
		} else {

		}
	}
}

func main() {
	fmt.Println("Hello World")
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}
