package pipline


import (
	"container/list"
	"sync"
)

func Producer(nums *list.List)<-chan int {
	out:=make(chan int)
	go func() {
		defer close(out)
		for e := (*nums).Front(); e != nil; e = e.Next() {
			out<-e.Value.(int)
		}
	}()
	return out
}

func Squre(inch <-chan int) <-chan int{
	out := make(chan int)
	go func() {
		defer close(out)
		for n:= range inch {
			out <- n*n
		}
	}()
	return out
}


func Merge(cs ...<-chan int) <-chan int {
	out := make(chan int,10000000)

	var wg sync.WaitGroup

	collect := func(in <-chan int) {
		defer wg.Done()
		for n := range in {
			out <- n
		}
	}

	wg.Add(len(cs))
	// FAN-IN
	for _, c := range cs {
		go collect(c)
	}

	// 错误方式：直接等待是bug，死锁，因为merge写了out，main却没有读
	// wg.Wait()
	// close(out)

	// 正确方式
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}