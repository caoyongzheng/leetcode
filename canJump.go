package leetcode

/*
	Given an array of non-negative integers, you are initially positioned at the first index of the array.

	Each element in the array represents your maximum jump length at that position.

	Determine if you are able to reach the last index.

	For example:
	A = [2,3,1,1,4], return true.

	A = [3,2,1,0,4], return false.
*/

func canJump(nums []int) bool {
	return canJump2(0, nums)
}

func canJump2(position int, nums []int) bool {
	if position == len(nums) {
		return true
	}
	if nums[position] == 0 {
		return false
	}
	if nums[position] >= len(nums)-position-1 {
		return true
	}
	for i := 1; i <= nums[position]; i++ {
		if canJump2(position+i, nums) {
			return true
		}
	}
	return false
}
