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
*********
 *******
  *****
   ***
    *

*/

func P9(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			fmt.Print(" ")

		}
		for j := 0; j < i*2+1; j++ {
			fmt.Print("*")
		}

		for j := 0; j < n-i-1; j++ {
			fmt.Print(" ")
		}

		fmt.Println()
	}

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}

		for k := 0; k < (n-1-i)*2+1; k++ {
			fmt.Print("*")
		}

		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
