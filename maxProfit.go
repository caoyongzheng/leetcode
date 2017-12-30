package leetcode

func maxProfit(prices []int) (ma int) {
	if len(prices) == 0 {
		return 0
	}
	mi := prices[0]
	for i := 1; i < len(prices); i++ {
		ma = max(ma, prices[i]-mi)
		mi = min(prices[i], mi)
	}
	return
}
