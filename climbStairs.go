package leetcode

func climbStairs(n int) (sum int) {
	l := n / 2
	for i := 0; i <= l; i++ {
		sum += selector(i, n-i)
	}
	return
}
