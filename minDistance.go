package leetcode

import (
	"strings"
)

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
	if strings.Contains(w2, w1) {
		return len(w2) - len(w1)
	}
	for w1[0] == w2[0] {
		w1, w2 = w1[1:], w2[1:]
	}
	if strings.Contains(w2, w1) {
		return len(w2) - len(w1)
	}
	r := minDistance(w1, w2[1:]) + 1
	r1 := minDistance(w1[1:], w2[1:]) + 1
	if r > r1 {
		r = r1
	}
	r2 := minDistance(w1[1:], w2) + 1
	if r > r2 {
		r = r2
	}
	return r
}
