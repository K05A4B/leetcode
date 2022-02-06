package main

import (
	"fmt"
	// "sort"
)

func main() {
	var strs = []string{
		"flower","flow","flight",
		// "ab", "a",
		// "reflower","flow","flight",
	}

	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	var before string
	var Outputs = ""
	var char = ""

	if str := strs[0]; str != "" {
		before = string(str[0])
	} else {
		return ""
	}

	for {
		for _, str := range strs {
			for i := 0; i < len(str); i++ {
				char = string(str[i])
				if char == before {
					before = char
				}
				continue
			}
		}
		if char == before {
			Outputs += char
		} else {
			break
		}
    }
	return Outputs
}