package hashing

func wordSubsets(words1 []string, words2 []string) []string {
	freq2 := [26]int{}

	for _, w := range words2 {
		freq := getFreq(w)

		for i := range freq {
			freq2[i] = max(freq2[i], freq[i])
		}
	}

	result := []string{}

Words1:
	for _, word := range words1 {
		freq := getFreq(word)

		for i := range freq2 {
			if freq[i] < freq2[i] {
				continue Words1
			}
		}

		result = append(result, word)
	}

	return result
}

func getFreq(s string) [26]int {
	freq := [26]int{}
	for i := range s {
		freq[s[i]-'a']++
	}
	return freq
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// https://leetcode.com/problems/word-subsets/
