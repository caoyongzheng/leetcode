package leetcode

// Find All Numbers Disappeared in an Array

// Given an array of integers where 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.

// Find all the elements of [1, n] inclusive that do not appear in this array.

// Could you do it without extra space and in O(n) runtime? You may assume the returned list does not count as extra space.

// Example:

// Input:
// [4,3,2,7,8,2,3,1]

// Output:
// [5,6]

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func findDisappearedNumbers(nums []int) []int {
	for _, v := range nums {
		i := abs(v) - 1
		if nums[i] > 0 {
			nums[i] = -nums[i]
		}
	}
	r := []int{}
	for i := range nums {
		if nums[i] > 0 {
			r = append(r, i+1)
		}
	}
	return r
}
