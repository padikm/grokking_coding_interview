package slidingwindow

func NumOfSubarrays(arr []int, k int, threshold int) int {
    return numOfSubarrays1(arr,k,threshold)
}

func numOfSubarrays1(arr []int, k int, threshold int) int {
    res := 0
	for i:=0;i<len(arr);i++ {
		j:=i
		avg :=0
		for j<len(arr) && j-i+1 <= k && i+k-1 < len(arr) {
			avg = avg + arr[j]
			j++
		}
		avg = avg / k
		if i+k == j && avg >= threshold {
			res++
		}
	}
	return res
}