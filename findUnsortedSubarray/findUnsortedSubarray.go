package findUnsortedSubarray

func findUnsortedSubarray(nums []int) int {
	startIndex, endIndex, hasFindStartIndex, hasFindEndIndex := 0, 0, false, false
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				hasFindStartIndex = true
				break
			}
		}
		if hasFindStartIndex {
			startIndex = i
			break
		}
	}

	for i := len(nums) - 1; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if nums[i] < nums[j] {
				hasFindEndIndex = true
				break
			}
		}
		if hasFindEndIndex {
			endIndex = i
			break
		}
	}

	if endIndex == startIndex {
		return 0
	}

	return endIndex - startIndex + 1
}
