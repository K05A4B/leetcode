package main

import "fmt"

func main() {
	var param = [][]int{
		[]int{1,2,3},
		[]int{4,5,6},
		[]int{7,8,9},
	}
	rotate(param)
	fmt.Println(param)
}

func rotate(matrix [][]int) {
	var matrixLen = len(matrix)
	var CopyMatrix = make([][]int, matrixLen)
	var j = 0
	copy(CopyMatrix, matrix)
	for {
		var temp []int
		for i := (matrixLen - 1); i >= 0; i-- {
			temp = append(temp, CopyMatrix[i][j])
		}
		matrix[j] = temp
		j++
		if j >= matrixLen {
			break
		}
    }
}