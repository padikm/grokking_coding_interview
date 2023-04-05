package golang

import (
	"fmt"
	"runtime"
	"sync"
)

var value int
var wg sync.WaitGroup

func Race() {
	wg.Add(2)
	go incerement(1)
	go incerement(2)
	wg.Wait()
	fmt.Println(value)
}

func incerement(i int) {
	defer wg.Done()
	for c := 0; c < 2; c++ {
		counter := value
		runtime.Gosched()
		counter++
		value = counter
	}
}
