package countBits

func countBits(num int) []int {
	result := make([]int, num+1)
	for i := 0; i <= num; i++ {
		result[i] = result[i>>1] + i&1
	}
	return result
}
