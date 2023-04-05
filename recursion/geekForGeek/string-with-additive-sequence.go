package geekForGeek

import "fmt"

func AdditiveSeq() {
	additiveSeq("235813", 0, "")
}

func additiveSeq(s string, j int, res string) {
	if j == len(s) {
		return
	}
	for i := j; i < len(s); i++ {
		fmt.Println(s[j : i+1])
		additiveSeq(s, j+1, s[j:i+1])
	}
}
