package reconstructQueue

import "sort"

// SortPeople ...
type SortPeople [][]int

func (a SortPeople) Len() int      { return len(a) }
func (a SortPeople) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortPeople) Less(i, j int) bool {
	return a[i][0] > a[j][0] || (a[i][0] == a[j][0] && a[i][1] < a[j][1])
}

func reconstructQueue(people [][]int) [][]int {
	res := SortPeople(people)
	sort.Sort(res)
	for i, v := range res {
		for j := i; j > v[1]; j-- {
			res[j], res[j-1] = res[j-1], res[j]
		}
	}
	return res
}
