package recursion

import "fmt"

func Recursion(n int) {

	if n <=1 {
		return
	}
	fmt.Println(n)
	Recursion(n-1)
	fmt.Println(n)
	Recursion(n-1)
}
