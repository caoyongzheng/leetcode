package leetcode

// House Robber

// You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security system connected and it will automatically contact the police if two adjacent houses were broken into on the same night.

// Given a list of non-negative integers representing the amount of money of each house, determine the maximum amount of money you can rob tonight without alerting the police.

func rob(nums []int) int {
	m := make(map[int]int, len(nums))
	for i := range nums {
		robN(nums, i+1, m)
	}
	return m[len(nums)]
}

func robN(nums []int, n int, m map[int]int) int {
	m[n] = max(m[n-1], m[n-2]+nums[n-1])
	return m[n]
}
