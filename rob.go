package leetcode

// House Robber

// You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security system connected and it will automatically contact the police if two adjacent houses were broken into on the same night.

// Given a list of non-negative integers representing the amount of money of each house, determine the maximum amount of money you can rob tonight without alerting the police.

func rob(nums []int) int {
	t1, t2 := 0, 0
	for i := range nums {
		t2, t1 = max(t2, t1+nums[i]), t2
	}
	return t2
}
