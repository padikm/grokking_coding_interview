package leet

var store = make(map[int]int)

func ClimbStairs(n int) int {
	if v, ok := store[n]; ok {
		return v
	}
	if n == 0 {
		return 1
	}

	if n < 0 {
		return 0
	}

	res := ClimbStairs(n-1) + ClimbStairs(n-2)
	store[n] = res
	return res

}

func ClimbStairsFib(n int) int {

	var res = make([]int, 0)
	if n == 1 {
		return 1
	}

	res = append(res, 1)
	res = append(res, 1)

	for i := 2; i <= n; i++ {
		res = append(res, res[i-1]+res[i-2])
	}

	return res[len(res)-1]
}
