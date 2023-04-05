// You can edit this code!
// Click here and start typing.
package educative

func LongSum(a []int, l int) int {
	sum := 0
	for i := 0; i < l; i++ {
		sum = sum + a[i]
	}
	j := 0
	for i := l; i < len(a); i++ {
		tsum := sum - a[j] + a[i]
		if tsum > sum {
			sum = tsum
		}
		j++
	}
	return sum
}
