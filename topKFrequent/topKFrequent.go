package topKFrequent

func topKFrequent(nums []int, k int) []int {
	counts := make(map[int]int)
	for _, v := range nums {
		counts[v]++
	}

	freqs := make([][]int, len(nums)+1)
	for k, v := range counts {
		freqs[v] = append(freqs[v], k)
	}

	r := make([]int, 0, k)
	for i := len(freqs) - 1; i >= 0; i-- {
		for _, v := range freqs[i] {
			r = append(r, v)
			if len(r) == k {
				return r
			}
		}
	}

	return r
}
