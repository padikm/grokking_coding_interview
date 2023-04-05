package recursion

import "fmt"

func PrintMaze() {
	fmt.Println(printMaze(3, 3, ""))

}

func printMaze(r, l int, path string) int {
	if r == 1 && l == 1 {
		fmt.Println(path)
		return 1
	}
	count := 0
	if r > 1 {
		count = count + printMaze(r-1, l, path+"D")
	}
	if l > 1 {
		count = count + printMaze(r, l-1, path+"R")
	}
	return count
}
