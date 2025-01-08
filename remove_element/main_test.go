package main

import (
	"slices"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	val := 3
	expectNums := []int{2, 2}
	// expects [2,2,_,_]
	k := removeElement(nums, val)

	if k != 2 {
		t.Errorf("Got wrong length: %q", k)
	}
	slices.Sort(nums)
	for i := 0; i < 2; i++ {
		if nums[i] != expectNums[i] {
			t.Error("Wrong array expected vs actual", expectNums, nums)
		}
	}
}
