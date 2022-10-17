package backtracking

type DataRef struct {
	Lookup map[rune]string
	Digits []rune
	Ans    []string
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	lookup := map[rune]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "prqs",
		'8': "tuv",
		'9': "wxyz",
	}
	// store in struct to avoid copying data on each recurse
	data := &DataRef{
		Lookup: lookup,
		Digits: []rune(digits),
		Ans:    []string{},
	}

	getCombinations(0, "", data)
	return data.Ans
}

func getCombinations(start int, combo string, data *DataRef) {
	if start > len(data.Digits) {
		data.Ans = append(data.Ans, combo)
		return
	}

	char := data.Digits[start]
	values := data.Lookup[char]

	for _, ch := range values {
		tmp := combo + string(ch)
		getCombinations(start+1, tmp, data)
	}

}
