package leet

import "fmt"

func AllPathsSourceTarget(graph [][]int) [][]int {
	final := make([][]int, 0)
	aps(graph, []int{}, 0, &final)
	fmt.Println(final)
	return final

}

func aps(g [][]int, res []int, j int, final *[][]int) {
	if len(res) > 0 && res[0] != 0 {
		return
	}

	if j == len(g)-1 {
		res = append(res, j)
		tmp := make([]int, 0)
		for i := 0; i < len(res); i++ {
			tmp = append(tmp, res[i])
		}
		*final = append(*final, tmp)
		return
	}

	for i := 0; i < len(g[j]); i++ {
		res = append(res, j)
		aps(g, res, g[j][i], final)
		res = res[:len(res)-1]
	}
}
