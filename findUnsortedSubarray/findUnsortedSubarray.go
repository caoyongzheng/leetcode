package findUnsortedSubarray

func findUnsortedSubarray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	startIndex := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[startIndex] {
			startIndex = findMid(nums, 0, startIndex, nums[i])
		} else if i == startIndex+1 {
			startIndex = i
		}
	}

	endIndex, maxValue := startIndex, nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxValue {
			maxValue = nums[i]
		} else if nums[i] < maxValue {
			endIndex = i
		}
	}

	return endIndex - startIndex + min(1, endIndex-startIndex)
}

func findMid(nums []int, startIndex, endIndex int, v int) int {
	s, e, m := startIndex, endIndex, (startIndex+endIndex)/2
	for s != e {
		if nums[m] < v {
			s, m = m+1, (m+e+1)/2
		} else if nums[m] > v {
			m, e = (s+m)/2, m
		} else {
			return min(m+1, endIndex)
		}
	}
	return m
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
