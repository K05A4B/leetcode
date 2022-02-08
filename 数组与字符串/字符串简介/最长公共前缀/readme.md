### 最长公共前缀

编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

> 示例 1：

```
输入：strs = ["flower","flow","flight"]
输出："fl"
```

示例 2：

```
输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。
```

提示：

+ 1 <= strs.length <= 200
+ 0 <= strs[i].length <= 200
+ strs[i] 仅由小写英文字母组成
  
> 链接：[](https://leetcode-cn.com/leetbook/read/array-and-string/ceda1/)

---

> 使用纵向扫描

循环扫描所有字符串的第 `charIndex` 个字符, 并且与下一个字符串的第 `charIndex`的字符比对, 如果不统一则立即返回 `outouts`

> 代码: 

```go
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

```