package leetcode

// Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.

// For example, given nums = [0, 1, 0, 3, 12], after calling your function, nums should be [1, 3, 12, 0, 0].

// Note:
// You must do this in-place without making a copy of the array.
// Minimize the total number of operations.
// Credits:
// Special thanks to @jianchao.li.fighter for adding this problem and creating all test cases.

func moveZeroes(nums []int) {
	for i, j := 0, 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if i != j {
				nums[j], nums[i] = nums[i], 0
			}
			j++
		}
	}
}
