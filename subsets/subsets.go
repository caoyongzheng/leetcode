package subsets

func subsets(nums []int) [][]int {
	r := [][]int{[]int{}}
	for i := 1; i <= len(nums); i++ {
		r = fillSubSets([]int{}, i, nums, r, 0)
	}
	return r
}

func fillSubSets(arr []int, l int, nums []int, r [][]int, start int) [][]int {
	if len(arr) == l {
		return append(r, arr)
	}
	for i := start; i < len(nums); i++ {
		r = fillSubSets(append(arr, nums[i]), l, nums, r, i+1)
	}
	return r
}
