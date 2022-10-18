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
		'7': "qprs",
		'8': "tuv",
		'9': "wxyz",
	}

	data := &DataRef{
		Lookup: lookup,
		Digits: []rune(digits),
		Ans:    []string{},
	}

	backtrack(0, "", data)
	return data.Ans
}

func backtrack(start int, str string, data *DataRef) {
	if start >= len(data.Digits) {
		data.Ans = append(data.Ans, str)
		return
	}

	digit := data.Digits[start]
	chars := data.Lookup[digit]

	for _, ch := range chars {
		tmp := str + string(ch)
		backtrack(start+1, tmp, data)
	}
}

// https://leetcode.com/problems/letter-combinations-of-a-phone-number/
