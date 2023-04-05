package leet

func NumberOfSteps(num int) int {
	if num == 0 {
		return 0
	}
	if num%2 == 0 {
		return 1 + NumberOfSteps(num/2)
	}
	return 1 + NumberOfSteps(num-1)
}
