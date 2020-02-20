package base

import (
	"fmt"
	"time"
)

func makeit() <-chan string {
	var sp = make(chan string, 2)
	sp <- "fdfsfs"
	sp <- "fdfsfs"
	fmt.Println("11")
	select {
	case sp<- "fdf":
	case <-time.After(1*time.Second):
		return sp


	}
	sp <- "fdfsfs"
	sp <- "fdfsfs"
	sp <- "fdfsfs"
	sp <- "fdfsfs"
	sp <- "fdfsfs"
	sp <- "fdfsfs"
	sp <- "fdfsfs"
	sp <- "fdfsfs"
	sp <- "fdfsfs"
	fmt.Println("fdsfjhskfhsafhsjkfuyr32783128318")
	return sp
}

func ChannelStart() {

	defer func() {
		if error := recover(); error != nil {
			fmt.Println(error)
		}
	}()
	var ss = make(chan string, 1)

	/*go func() {
		for ll := range ss {
			fmt.Println(ll)
			time.Sleep(1 * time.Second)
		}
	}()*/

	for {
		select {
		case ss <- "fdfsdf":
			fmt.Println("insert success")
		case <-time.After(2 * time.Second):
			fmt.Println("time out......")
		case <-makeit():
			fmt.Println("fdfsfsfsdfsdfdfdsfsd")

			time.Sleep(2 * time.Second)
		}


	}
}
