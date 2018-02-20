package findCircleNum

func findCircleNum(M [][]int) int {
	circles, counts := 0, make(map[int]bool)
	for i := 0; i < len(M); i++ {
		counts[i] = true
	}

	for i := range counts {
		members := map[int]bool{i: true}
		arrM := map[int]bool{i: true}
		for len(arrM) != 0 {
			for m := range arrM {
				for j := 0; j < len(M); j++ {
					if _, ok := members[j]; M[m][j] == 1 && !ok {
						members[j] = true
						arrM[j] = true
					}
				}
				delete(arrM, m)
				delete(counts, m)
				break
			}
		}
		circles++
	}
	return circles
}
