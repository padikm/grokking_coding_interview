package leet

import (
	"strconv"
)

func RestoreIpAddresses(s string) []string {
	ips := make([]string, 0)
	restoreIpAddrUtil(s, 0, "", 0, &ips)
	return ips
}

func restoreIpAddrUtil(s string, j int, res string, dot int, ips *[]string) {
	if dot == 4 && j == len(s) {
		*ips = append(*ips, res[:len(res)-1])
		return
	}

	for i := j; i < min(j+3, len(s)); i++ {
		iIp, err := strconv.Atoi(s[j : i+1])
		if err == nil && iIp <= 255 && (i == j || string(s[j]) != "0") {
			restoreIpAddrUtil(s, i+1, res+s[j:i+1]+".", dot+1, ips)
		}
	}
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
