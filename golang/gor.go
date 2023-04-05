package golang

import (
	"fmt"
	"sync"
	"time"
)

func GoRoutine() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 'a'; i < 'a'+26; i++ {
			fmt.Printf("%c", i)
		}

	}()

	go func() {
		defer wg.Done()
		for i := 'A'; i < 'A'+26; i++ {
			fmt.Printf("%c", i)
		}
	}()

	wg.Wait()
	fmt.Println("Done")
}

func GoRoutine1() {
	//var c chan int
	c := make(chan int)
	go func() {
		c <- 1
		close(c)
	}()
	fmt.Println(<-c)

}

func GoRoutine2() {
	//var c chan int
	c := make(chan int)
	go func() {
		//defer close(c)
		//fmt.Println("1.1")
		//by the time <-c happens GoRoutine2 exits
		//so we cannot see the output 1.1
		// 2 things we can do
		//1. Sleep after c <- 1 in GoRoutine2
		//2. read in GoRoutine2 write in go routine
		if v, ok := <-c; ok {
			fmt.Println("1.1", v)
		}

	}()
	//fmt.Println("1.2")
	c <- 1
}

func GoRoutine3() {
	c := make(chan int)
	go func() {
		c <- 1
	}()
	for {
		select {
		case v := <-c:
			fmt.Println(v)
			break
		default:
			fmt.Println("running default case")
			time.Sleep(time.Second * 1)

		}

	}

}
