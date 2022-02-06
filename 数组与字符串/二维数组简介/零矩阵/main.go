package main

import "fmt"

func main() {
	// var param = [][]int{
	// 	[]int{1,1,1},
	// 	[]int{1,0,1},
	// 	[]int{1,1,1},
	// }
	var param = [][]int {
        []int{0,1,2,0},
        []int{3,4,5,2},
        []int{1,3,1,5},
	}
	setZeroes(param)
	fmt.Printf("%v\n%v\n%v\n", param[0], param[1], param[2])
}

func setZeroes(matrix [][]int)  {
	var XLen = len(matrix)
	var YLen = len(matrix[0])
	var CopyMatrix = make([][]int, XLen)
	var ZeroesPosition [][]int
	copy(CopyMatrix, matrix)
	// 寻找出所有0的位置
	for x, value := range CopyMatrix {
		for y, value := range value {
			if value == 0 {	
				ZeroesPosition = append(ZeroesPosition, []int{x, y})
			}
		}
	}
	for _, position := range ZeroesPosition {
		var x = position[0]
		var y = position[1]
		for i := 0; i < XLen; i++ {
			matrix[i][y] = 0
		}
		for j := 0; j < YLen; j++ {
			matrix[x][j] = 0
		}
	}
}

// func set(matrix [][]int, blackList * [][]int, x, y int) {
// 	var position []int
// 	position = append(position, x, y)
// 	*blackList = append(*blackList, position)
// 	matrix[position[0]][position[1]] = 0
// }

// func isExistInBlackList(blackList[][]int, x, y int) bool {
// 	for _, value := range blackList {
// 		if value[0] == x && value[1] == y {
// 			fmt.Println(false)
// 			return false
// 		}
// 	}
// 	return true
// }
