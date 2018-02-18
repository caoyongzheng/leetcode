package countSubstrings

func countSubstrings(s string) (c int) {
	for i := range s {
		for j, o := 0, min(i, len(s)-i-1); j <= o && s[i-j] == s[i+j]; j++ {
			c++
		}
		for j, o := 0, min(i, len(s)-i-2); j <= o && s[i-j] == s[i+j+1]; j++ {
			c++
		}
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
