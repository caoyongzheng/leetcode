package leetcode

import "strings"

// Given two words word1 and word2, find the minimum number of steps required to convert word1 to word2. (each operation is counted as 1 step.)

// You have the following 3 operations permitted on a word:

// a) Insert a character
// b) Delete a character
// c) Replace a character

func minDistance(word1 string, word2 string) int {
	w1, w2 := word1, word2
	if len(word1) > len(word2) {
		w1, w2 = w2, w1
	}
	if len(w1) == 0 {
		return len(w2)
	}
	if len(w1) == 1 {
		i := strings.Index(w2, string(w1[0]))
		if i != -1 {
			return len(w2) - 1
		}
		return len(w2)
	}
	i1 := strings.Index(w2, string(w1[0]))
	i2 := strings.LastIndex(w2, string(w1[1]))
	if i1 == -1 {
		if i2 == 0 {
			return minDistance(w1[1:], w2) + 1
		}
		return minDistance(w1[1:], w2[1:]) + 1
	} else if i1 == 0 {
		return minDistance(w1[1:], w2[1:])
	} else {
		if i1 <= len(w2)-len(w1) {
			if i1 > i2 {
				return minDistance(w1[1:], w2[1:]) + 1
			}
			return minDistance(w1, w2[1:]) + 1
		}
		if i2 == 0 {
			return minDistance(w1[1:], w2) + 1
		}
		return minDistance(w1[1:], w2[1:]) + 1
	}
}
