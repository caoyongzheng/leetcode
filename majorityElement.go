package leetcode

// Given an array of size n, find the majority element. The majority element is the element that appears more than ⌊ n/2 ⌋ times.

// You may assume that the array is non-empty and the majority element always exist in the array.

func majorityElement(nums []int) int {
	r, c := 0, 0
	for _, v := range nums {
		if c == 0 {
			r = v
		}
		if r == v {
			c++
		} else {
			c--
		}
	}
	return r
}
