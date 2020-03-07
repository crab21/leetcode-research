package main

import (
	"fmt"
)

func main() {

	/* sp :=[]int{}

	if sp == nil {

		fmt.Println(sp)
	}
	sp = append(sp, 1)
	fmt.Println(sp)*/
	//leetcode.Palindrome("babad")

	//var sp="s123456"
	//fmt.Println(sp[1:4])

	//palind := leetcode.LongestPalind("cabbanmabbam")
	//fmt.Println("palind")
	//leetcode.DPStart()
	//arr :=[]int{1,20,9,10,11,22,34,234,432,324}
	//sortAlgorithm.MergeForSort(arr,0,len(arr)-1)
	//fmt.Println(arr)
	//base.SliceStart()
	/*stopCh := make(chan interface{})
	go func() {
		var s = 0
		for {
			stopCh <- "1111"
			s++
			if s == 10 {
				close(stopCh)
				return
			}
			time.Sleep(1 * time.Second)
		}
	}()

	for {
		v, ok := <-stopCh

		if !ok {

			fmt.Println("closed")
			break
		}
		fmt.Println(v)
	}*/
/*	x, y := 1, 2
	fmt.Println(test(x, y, add))*/
	ss := "sfdsfjsdkfjlsdfjsdlnslcnlsdfjlncnsdlkjflwefjwelfnl1831u482fjkencsd"
	/*sp := *(*[]byte)(unsafe.Pointer(&ss))*/
	sp:=[]byte(ss)
	fmt.Println(sp)
	//debug.PrintStack()

}

type Callback func(x, y int) int

// 提供一个接口，让外部去实现
func test(x, y int, callback Callback) int {
	return callback(x, y)
}

// 回调函数的具体实现
func add(x, y int) int {
	return x + y
}
func adds(x, y int) int {
	return x * y
}
