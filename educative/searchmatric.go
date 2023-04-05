// You can edit this code!
// Click here and start typing.
package educative

func searchMatrix(a [][]int, key int) bool {
	lastColumn := []int{}
	for i := 0; i < len(a); i++ {
		lastColumn = append(lastColumn, a[i][len(a[0])-1])
	}
	rowIndex := binarySearch(lastColumn, key)
	if rowIndex >= len(a) {
		return false
	}
	columnIndex := binarySearch(a[rowIndex], key)
	return a[rowIndex][columnIndex] == key
}

func binarySearch(a []int, key int) int {
	l := 0
	h := len(a)
	for l < h {
		mid := (l + h) / 2
		if a[mid] == key {
			return mid
		}

		if a[mid] > key {
			h = mid
		} else {
			l = mid + 1
		}
	}

	return h
}
