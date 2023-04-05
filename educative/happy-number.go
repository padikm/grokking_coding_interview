package educative

func IsHappy(n int) bool {
	slow := sqrDigits(n)
	fast := sqrDigits(slow)
	for fast != slow {
		slow = sqrDigits(slow)
		fast = sqrDigits(sqrDigits(fast))
	}

	if slow == 1 {
		return true
	}
	return false

}

func sqrDigits(n int) int {
	sum := 0
	for n != 0 {
		rem := n % 10
		sum = sum + rem*rem
		n = n / 10
	}
	return sum
}
