package recursion

import "fmt"

func InvertedTriangle(i, j, n int) {
	if i == n {
		return
	}
	fmt.Print("*")
	if j == n-1 {
		fmt.Println("")
		InvertedTriangle(i+1, i+1, n)
		return
	}
	InvertedTriangle(i, j+1, n)
}

func NormalTriangle(i, j, n int) {
	if i == n {
		return
	}
	fmt.Print("*")
	if j == i {
		fmt.Println("")
		NormalTriangle(i+1, 0, n)
		return
	}
	NormalTriangle(i, j+1, n)
}
