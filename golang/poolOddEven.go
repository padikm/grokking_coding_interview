package golang

import "fmt"

func PoolOddEven() {
	c := make(chan struct{})
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			<-c
			if i%2 == 0 {
				fmt.Println("Even ", i)
			}

		}

	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			if i%2 != 0 {
				fmt.Println("Odd ", i)
			}
			c <- struct{}{}
		}
		//c <- structtest{}{}
	}()

	wg.Wait()
}
