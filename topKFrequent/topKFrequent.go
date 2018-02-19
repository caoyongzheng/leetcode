package topKFrequent

func topKFrequent(nums []int, k int) (r []int) {
	m := make(map[int]*[2]int)
	for _, v := range nums {
		mv, ok := m[v]
		if !ok {
			r = append(r, v)
			m[v] = &[2]int{len(r) - 1, 1}
			continue
		}

		mv[1]++

		for mv[0] != 0 {
			pmv := m[r[mv[0]-1]]
			if pmv[1] >= mv[1] {
				break
			}
			r[mv[0]], r[pmv[0]] = r[pmv[0]], r[mv[0]]
			mv[0]--
			pmv[0]++
		}
	}
	r = r[0:k]
	return
}
