package countSubstrings

func countSubstrings(s string) (c int) {
	for i := range s {
		for j := len(s) - 1; j >= i; j-- {
			c++
			// whether start i and end j is Palindromic
			for k := 0; k <= (j-i)/2; k++ {
				if s[i+k] != s[j-k] {
					c--
					break
				}
			}
		}
	}
	return
}
