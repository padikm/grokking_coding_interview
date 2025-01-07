package pattern

import "fmt"

/*

1 2 3 4 5
1 2 3 4
1 2 3
1 2
1
*/

func P6(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			fmt.Print(j + 1)
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
