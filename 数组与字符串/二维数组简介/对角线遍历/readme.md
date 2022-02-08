> 示例 1：

```
输入：mat = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,4,7,5,3,6,8,9]
```

> 示例 2：

```
输入：mat = [[1,2],[3,4]]
输出：[1,2,3,4]
```

> 提示：

+ m == mat.length
+ n == mat[i].length
+ 1 <= m, n <= 104
+ 1 <= m * n <= 104
+ -105 <= mat[i][j] <= 105


> 题目链接

[点我](https://leetcode-cn.com/leetbook/read/array-and-string/cuxq3/)

---

> 条件:

1. `i != 0 && (x + y - 1) % 2 == 0` 的情况下
    + `x == (rowLen - 1) && y == 0` 向右移动 (y++)
    + `y == 0` 向下移动 (x++)
    + 如果不满足以上条件则向 左下移动 (y--; x++)

2. 如果不满足以上条件则
    + `y == (colLen - 1) && x == 0` 向下移动 (x++)
    + `x == 0` 向右移动 (y++)
    + 如果不满足以上条件则向 右上下移动 (y++; x--)

> 代码

``` go
func findDiagonalOrder(mat [][]int) []int {
	var outputs []int
	var rowLen = len(mat)
	var colLen = len(mat[0])
	var x = 0
	var y = 0

	for i := 0; i < (rowLen * colLen); i++ {
		outputs = append(outputs, mat[x][y])
		if i != 0 && (x + y - 1) % 2 == 0 {
			if x == (rowLen - 1) && y == 0 {
				y++
			}else if y == 0 {
				x++ 
			} else {
				y--
				x++
			}
		} else {
			if y == (colLen - 1) && x == 0 {
				x++
			} else if x == 0 {
				y++
		    } else {
				x--
				y++
			}
		}
	}

	return outputs
}
```