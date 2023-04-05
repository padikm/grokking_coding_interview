package golang

import (
	"fmt"
	"time"
)

func Worker() {
	numJobs := 5
	jobs := make(chan struct{}, numJobs)
	chs := make(chan int, numJobs)
	for i := 0; i < 3; i++ {
		go Work(i, jobs, chs)
	}
	for i := 0; i < numJobs; i++ {
		jobs <- struct{}{}
	}
	close(jobs)
	time.Sleep(time.Second * 5)
	close(chs)
	//close(chs)
	//for a := 1; a <= numJobs; a++ {
	//	<-chs
	//}
	for v := range chs {
		fmt.Println(v)
	}
}

func Work(id int, jobs chan struct{}, chs chan int) {
	for i := range jobs {
		fmt.Println("Job ", id)
		fmt.Println(i)
		chs <- id * id
	}
}
