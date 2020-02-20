package main

import (
	"fmt"
	"time"
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
	stopCh := make(chan interface{})
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
	}
}
