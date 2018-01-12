package leetcode

// 438. Find All Anagrams in a String

// Given a string s and a non-empty string p, find all the start indices of p's anagrams in s.

// Strings consists of lowercase English letters only and the length of both strings s and p will not be larger than 20,100.

// The order of output does not matter.

// Example 1:

// Input:
// s: "cbaebabacd" p: "abc"

// Output:
// [0, 6]

// Explanation:
// The substring with start index = 0 is "cba", which is an anagram of "abc".
// The substring with start index = 6 is "bac", which is an anagram of "abc".
// Example 2:

// Input:
// s: "abab" p: "ab"

// Output:
// [0, 1, 2]

// Explanation:
// The substring with start index = 0 is "ab", which is an anagram of "ab".
// The substring with start index = 1 is "ba", which is an anagram of "ab".
// The substring with start index = 2 is "ab", which is an anagram of "ab".

func findAnagrams(s string, p string) (r []int) {
	sMap, pMap := make(map[byte]int), make(map[byte]int)
	for i := range p {
		pMap[p[i]]++
	}
	for i := range s {
		sMap[s[i]]++
		if i >= len(p) {
			sMap[s[i-len(p)]]--
			if sMap[s[i-len(p)]] <= 0 {
				delete(sMap, s[i-len(p)])
			}
		}
		if i >= len(p)-1 && isAnagram(sMap, pMap) {
			r = append(r, i-len(p)+1)
		}
	}
	return
}

func isAnagram(sMap map[byte]int, pMap map[byte]int) bool {
	if len(sMap) != len(pMap) {
		return false
	}
	for k, v := range sMap {
		if pMap[k] != v {
			return false
		}
	}
	return true
}
