package leetcode

// Given two words word1 and word2, find the minimum number of steps required to convert word1 to word2. (each operation is counted as 1 step.)

// You have the following 3 operations permitted on a word:

// a) Insert a character
// b) Delete a character
// c) Replace a character

func min(a int, others ...int) int {
	r := a
	for _, v := range others {
		if r > v {
			r = v
		}
	}
	return r
}

func minDistance(word1 string, word2 string) int {
	state := make([][]int, len(word1)+1)
	for i := range state {
		state[i] = make([]int, len(word2)+1)
	}
	for i := 0; i <= len(word1); i++ {
		for j := 0; j <= len(word2); j++ {
			if i == 0 {
				state[0][j] = j
			} else if j == 0 {
				state[i][0] = i
			} else if word1[i-1] == word2[j-1] {
				state[i][j] = state[i-1][j-1]
			} else {
				state[i][j] = min(state[i-1][j-1], state[i][j-1], state[i-1][j]) + 1
			}
		}
	}
	return state[len(word1)][len(word2)]
}
