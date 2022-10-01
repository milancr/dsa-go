package stackandqueue

import (
	"strings"
)

func simplifyPath(path string) string {
	stack := []string{}

	str := strings.Split(path, "/")

	for i := 0; i < len(str); i++ {
		if str[i] == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if str[i] == "." || str[i] == "" {
			continue
		} else {
			stack = append(stack, str[i])
		}
	}
	p := strings.Join(stack, "/")
	ans := "/" + p
	return ans
}

// https://leetcode.com/problems/simplify-path/
