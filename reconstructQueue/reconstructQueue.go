package reconstructQueue

func reconstructQueue(people [][]int) [][]int {
	frontLeftList := make([]int, len(people))
	for i, v := range people {
		frontLeftList[i] = v[1]
	}

	for i := 0; i < len(people); i++ {
		first := -1
		for j := i; j < len(people); j++ {
			if frontLeftList[j] == 0 && (first == -1 || people[first][0] > people[j][0]) {
				first = j
			}
		}

		people[i], people[first] = people[first], people[i]
		frontLeftList[i], frontLeftList[first] = frontLeftList[first], frontLeftList[i]

		for j := i + 1; j < len(people); j++ {
			if frontLeftList[j] != -1 && people[i][0] >= people[j][0] {
				frontLeftList[j]--
			}
		}
	}

	return people
}
