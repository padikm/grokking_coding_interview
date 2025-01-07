package pattern

import "fmt"

/*
*
**
***
****
*****
 */

func P2(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
