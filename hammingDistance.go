package leetcode

func hammingDistance(x, y int) int {
	s, c := x^y, 0
	for ; s != 0; s /= 2 {
		if s%2 == 1 {
			c++
		}
	}
	return c
}
