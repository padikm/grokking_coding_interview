package golang

import "fmt"

func CallGor() {
	go func() {
		fmt.Println("abc")
	}()

}
