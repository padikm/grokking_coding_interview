package pattern

import (
	"fmt"
)

/*
1
0 1
1 0 1
0 1 0 1
1 0 1 0 1
*/

func P11(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print((j + i + 1) % 2)
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
