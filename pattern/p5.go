package pattern

import "fmt"

/*

*****
****
***
**
*
 */

func P5(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
