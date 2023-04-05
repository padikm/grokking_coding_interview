package leet

func FirstBadVersion(n int) int {
	found := -1
	h := n
	l := 0

	if n == 1 && isBadVersion(1) {
		return n
	}

	for l <= h {
		m := (l + h) / 2
		if isBadVersion(m) {
			found = m
			h = m - 1
		} else {
			l = m + 1
		}
	}
	return found
}

func isBadVersion(i int) bool {
	return i == 4
}
