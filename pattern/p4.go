package pattern

import "fmt"

/*
1
2 2
3 3 3
4 4 4 4
5 5 5 5 5
*/
func P4(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print(i + 1)
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
