package geekForGeek

import "fmt"

func PrintAllCombOfStrGiveLen(a []string, k int) {
	printAllCombOfStrGiveLen(a, k, "")
}

func printAllCombOfStrGiveLen(a []string, k int, res string) {
	if k == 0 {
		fmt.Println(res)
		return
	}
	for i := 0; i < len(a); i++ {
		//res = res + a[i]
		printAllCombOfStrGiveLen(a, k-1, res+a[i])
		//res = res[0 : len(res)-1]
	}
}
