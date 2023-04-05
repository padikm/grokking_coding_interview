package geekForGeek

import "fmt"

func Minimumnumberdeletionsmakesortedsequence(a []int) {
	res := 0
	for i := 0; i < len(a); i++ {
		count := 0
		pos := i
		for j := i; j < len(a); j++ {
			if a[j] >= a[pos] {
				pos = j
				count++
			}
		}
		if count > res {
			res = count
		}
	}

	fmt.Println(len(a) - res)
}

func MinimumNumberDeletionsMakeSortedSequence(a []int) {
	fmt.Println(len(a) - minimumNumberDeletionsMakeSortedSequence(a, 0, 0, 0, 0))
}

func minimumNumberDeletionsMakeSortedSequence(a []int, i, j int, count int, pos int) int {
	if len(a) == i {
		return count
	}

	if j == len(a) {
		return max(count, minimumNumberDeletionsMakeSortedSequence(a, i+1, i+1, 0, i+1))

	}
	if a[j] >= a[pos] {
		count++
		pos = j
	}
	return max(count, minimumNumberDeletionsMakeSortedSequence(a, i, j+1, count, pos))
}
