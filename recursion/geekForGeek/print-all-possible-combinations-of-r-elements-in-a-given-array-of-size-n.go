package geekForGeek

import "fmt"

func PrintAllPosCombK(a []string) []string {
	return printAllPosCombK1(a, 0, 2, "")
}

func printAllPosCombK(s []string, k int, res string) {
	if len(s) == 0 && len(res) == k {
		fmt.Println(res)
	}

	if len(s) == 0 {
		//fmt.Println(res)
		return
	}

	printAllPosCombK(s[1:len(s)], k, res+s[0])
	printAllPosCombK(s[1:len(s)], k, res)

}

func printAllPosCombK1(s []string, j, k int, res string) []string {
	a := make([]string, 0)
	if len(res) == k {
		a = append(a, res)
		//fmt.Println(res)
		return a
	}
	for i := j; i < len(s); i++ {
		a = append(a, printAllPosCombK1(s, i+1, k, res+s[i])...)
	}

	return a
}
