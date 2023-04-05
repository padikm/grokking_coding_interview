package geekForGeek

import "fmt"

func MinimumStepsMinimizeNPerGivenCondition(n int) {
	fmt.Println(minimumStepsMinimizeNPerGivenConditionGreedy(n))
}

func minimumStepsMinimizeNPerGivenConditionGreedy(n int) int {
	if n <= 1 {
		return 0
	}

	if n%2 == 0 {
		return 1 + minimumStepsMinimizeNPerGivenConditionGreedy(n/2)
	}

	if n%3 == 0 {
		return 1 + minimumStepsMinimizeNPerGivenConditionGreedy(n/3)
	}
	return 1 + minimumStepsMinimizeNPerGivenConditionGreedy(n-1)
}
func minimumStepsMinimizeNPerGivenCondition(n int) int {

	if n == 0 {
		return n
	}

	res := minimumStepsMinimizeNPerGivenCondition(n - 1)
	if n%2 == 0 {
		return 1 + min(res, minimumStepsMinimizeNPerGivenCondition(n/2))
	}

	if n%3 == 0 {
		return 1 + min(res, minimumStepsMinimizeNPerGivenCondition(n/3))
	}
	return res
}
