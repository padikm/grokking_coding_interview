package golang

import (
	"fmt"
	"sync"
)

func BufferChan() {
	c := make(chan string)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c <- "COMPLETED"
		}()
	}
	go func() {
		defer close(c)
		wg.Wait()
	}()
	count := 0
	for v := range c {
		if v == "COMPLETED" {
			count++
		}
	}
	fmt.Println(count)
}
