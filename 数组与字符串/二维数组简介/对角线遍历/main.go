package main

import "fmt"

func main() {
	fmt.Println(findDiagonalOrder([][]int{
		[]int{1,2,3},
		[]int{4,5,6},
		[]int{7,8,9},
	}))
}

func findDiagonalOrder(mat [][]int) []int {
	// 0,0; 0,1; 1,0; 2,0; 1,1; 0,2; 1,2; 2,1; 2,2
}