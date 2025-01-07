package pattern

import (
	"fmt"
)

/*
*
   ***
  *****
 *******
*********

*/

func P7(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < i*2+1; k++ {
			fmt.Print("*")
		}

		// for l := 0; l < n-i+1; l++ {
		// 	fmt.Print("")
		// }
		fmt.Println()
	}
}
