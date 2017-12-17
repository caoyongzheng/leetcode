package leetcode

func maxSubArray(nums []int) int {
	max, sum := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if sum < 0 {
			sum = 0
		}
		if sum = sum + nums[i]; sum > max {
			max = sum
		}
	}
	return max
}
