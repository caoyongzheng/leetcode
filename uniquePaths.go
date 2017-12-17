package leetcode

// A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).

// The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).

// How many possible unique paths are there?

// Above is a 3 x 7 grid. How many possible unique paths are there?

// Note: m and n will be at most 100.

func uniquePaths(m int, n int) int {
	if m > n {
		return selector(m-1, m+n-2)
	}
	return selector(n-1, m+n-2)
}

func selector(k, n int) int {
	if k == 0 || k == n {
		return 1
	}
	return selector(k, n-1) * n / (n - k)
}
