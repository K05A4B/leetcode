package main

import "fmt"

func main() {
	// var strs = []string{"flower", "flow", "flight"}
	var strs = []string{"a", "ac"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	var outputs string
	var strsLen = len(strs)
	var charIndex = 0

	if strsLen == 0 || strs[0] == "" {
		return outputs
	}

	if strsLen == 1 && strs[0] != "" {
		return strs[0]
	}

	for {
		var str string
		var i   int
		for i, str = range strs {
			if (i + 1) == strsLen {
				continue
			}
			if len(str) <= charIndex {
				return outputs
			}
			if len(strs[i+1]) <= charIndex {
				return outputs
			}
			if strs[i+1][charIndex] != str[charIndex] {
				return outputs
			}
		}
		outputs += string(str[charIndex])
		charIndex++
	}
}
