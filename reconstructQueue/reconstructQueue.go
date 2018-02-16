package reconstructQueue

func reconstructQueue(people [][]int) [][]int {
	result := make([][]int, len(people))
	frontLeftList := make([]int, len(people))
	for i, v := range people {
		frontLeftList[i] = v[1]
	}

	for i := 0; i < len(people); i++ {
		// find the first people
		first := -1
		for j, v := range frontLeftList {
			if v == 0 && (first == -1 || people[first][0] > people[j][0]) {
				first = j
			}
		}

		result[i] = people[first]

		for j, v := range frontLeftList {
			if v != -1 && result[i][0] >= people[j][0] {
				frontLeftList[j]--
			}
		}
	}

	return result
}
