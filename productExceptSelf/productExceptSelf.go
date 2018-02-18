package productExceptSelf

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))

	result[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		result[i] = result[i+1] * nums[i+1]
	}

	for i, temp := 1, 1; i < len(nums); i++ {
		temp *= nums[i-1]
		result[i] *= temp
	}

	return result
}
