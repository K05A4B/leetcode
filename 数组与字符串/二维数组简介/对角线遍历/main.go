package main

import (
	"fmt"
)

func main() {
	// 坐标顺序: 0,0; 0,1; 1,0; 2,0; 1,1; 0,2; 1,2; 2,1; 2,2
	// 返回: 124753689

	// var param = [][]int{
	// 	[]int{1,2,3},
	// 	[]int{4,5,6},
	// 	[]int{7,8,9},
	// }

	// 坐标顺序: 0,0; 0,1; 1,0; 1,1;
	// 返回: 1234
	
	var param = [][]int{
		[]int{1,2},
		[]int{3,4},
	}
	fmt.Println(findDiagonalOrder(param))
}

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