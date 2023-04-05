package recursion

import "fmt"

func Print2DArrRecursion(a [][]int, i, j int) {
	if i == len(a) {
		return
	}
	if j < len(a[i]) {
		fmt.Println(a[i][j])
	}
	if len(a[i]) == j {
		Print2DArrRecursion(a, i+1, 0)
		return
	}
	Print2DArrRecursion(a, i, j+1)
}
