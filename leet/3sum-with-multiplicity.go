package leet

import (
	"sort"
)

func Threesumwithmultiplicity(a []int, k int) int {
	sort.Ints(a)
	mod := 1000000007
	m := make(map[int]int)
	for i := 0; i < len(a); i++ {
		m[a[i]]++
	}
	count := 0
	for i := 0; i < len(a); i++ {
		//m[a[i]]--
		if i > 0 && a[i] == a[i-1] {
			continue
		}
		count += (findPair(a, i, k, a[i], m)) % mod
	}
	return count
}

func findPair(a []int, i, t, k int, m map[int]int) int {
	l := i + 1
	r := len(a) - 1
	mod := 1000000007
	count := 0
	j := 0
	for l < r {
		sum := k + a[l] + a[r]
		if sum == t {
			j++
			if a[i] == a[r] {
				count = (count + m[a[i]]*(m[a[i]]-1)*(m[a[i]]-2)/6) % mod
			} else if a[i] == a[l] {
				count = (count + m[a[r]]*m[a[i]]*(m[a[i]]-1)/2) % mod
			} else if a[l] == a[r] {
				count = (count + m[a[i]]*m[a[l]]*(m[a[l]]-1)/2) % mod
			} else {
				count += (m[a[i]] * m[a[l]] * m[a[r]]) % mod
			}
			//m[a[i]]--
			//count = count + m[k]*m[a[l]]*m[a[r]]
			l++
			r--
			for l < len(a) && a[l] == a[l-1] {
				l++
			}
			for r > 0 && a[r] == a[r+1] {
				r--
			}
		} else if sum < t {
			l++
		} else {
			r--
		}
	}
	return count
}
