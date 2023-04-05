package leet

import (
	"math"
)

func FindClosestElements(arr []int, k int, x int) []int {
	low := 0
	high := len(arr) - 1
	for (high - low) >= k {
		if math.Abs(float64(arr[low]-x)) <= math.Abs(float64(arr[high]-x)) {
			high--
		} else {
			low++
		}
	}
	return arr[low : high+1]
}
