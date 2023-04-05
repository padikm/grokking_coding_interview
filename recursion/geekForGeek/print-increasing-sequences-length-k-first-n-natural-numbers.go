package geekForGeek

import (
	"fmt"
	"strconv"
)

func PrintIncreasingSeqGivenKAndN(n, k int) {
	printIncreasingSeqGivenKAndN(n, 1, k, "")
}

func printIncreasingSeqGivenKAndN(n, j, k int, res string) {
	if len(res) == k {
		fmt.Println(res)
		return
	}
	for i := j; i <= n; i++ {
		printIncreasingSeqGivenKAndN(n, i+1, k, res+strconv.Itoa(i))
	}
}
