package base

import (
	"encoding/json"
	"fmt"
	"unsafe"
)

func truncate() {
	/*	arr := [14]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
		fmt.Print("原数组：", arr)

		fmt.Println("substr....")

		s3 := make([]int, 20)
		copy(s3, arr[2:5:9])

		s1 := arr[2:5:9]
		fmt.Printf("now len %d,now cap %d", len(s1), cap(s1))

		fmt.Println()
		fmt.Println(s3)
		s3[1] = 22222

		s2 := arr[1:7]
		fmt.Println(s2)
		fmt.Printf("now len %d,now cap %d", len(s2), cap(s2))
		fmt.Println()


	*/
	ss1 := make([]int, 10, 20)
	fmt.Printf("now len %d,now cap %d", len(ss1), cap(ss1))
	fmt.Println()
	fmt.Println(ss1)
	ss1[0] = 10
	ss1[1] = 10
	ss1[2] = 10
	ss1[3] = 10
	ss1[4] = 10
	ss1[5] = 10
	ss1[6] = 10
	ss1[7] = 10
	ss1[8] = 10
	ss1[9] = 10
	ss1 = append(ss1, 10)
	ss1 = append(ss1, 11)
	ss1 = append(ss1, 12)
	ss1 = append(ss1, 13)
	ss1 = append(ss1, 14)
	ss1 = append(ss1, 15)
	ss1 = append(ss1, 16)
	ss1 = append(ss1, 17)
	ss1 = append(ss1, 18)
	ss1 = append(ss1, 19)
	ss1 = append(ss1, 20)
	ss1 = append(ss1, 9)
	ss1 = append(ss1, 8)
	ss1 = append(ss1, 7)
	ss1 = append(ss1, 6)
	ss1 = append(ss1, 5)
	ss1 = append(ss1, 4)
	ss1 = append(ss1, 3)
	ss1 = append(ss1, 2)
	ss1 = append(ss1, 1)
	ss1 = append(ss1, 0)
	fmt.Printf("now len %d,now cap %d", len(ss1), cap(ss1))
	fmt.Println()
	ss1 = append(ss1, 101)
	ss1 = append(ss1, 102)
	fmt.Printf("now len %d,now cap %d", len(ss1), cap(ss1))

	fmt.Println("\nmap slice information....")
	mmp := make(map[int]int, 20)
	fmt.Println(mmp)

	fmt.Println("string slice information....")
	stp := make([]byte, 0, 2048)
	fmt.Println(stp)

	fmt.Println("slice new ")
	var ssp = []int{2: 999}
	fmt.Println(cap(ssp) == 0)
	ssp = append(ssp, 0004)
	fmt.Println(ssp)
	fmt.Printf("now len %d,now cap %d", len(ssp), cap(ssp))
	fmt.Println("-----------------><")

	/*	fmt.Println("=======================")
		sm1:=make(chan []string, 20)
		for i:=0;i<20;i+=1 {
			sm1<-[]string{"2222"}
		}
		for i:=0;i<20;i+=1 {
			fmt.Println(<-sm1)
		}*/
	sts := make([]float64, 0, 5)
	sts = append(sts, 1)
	fmt.Printf("now len %d,now cap %d ===\n", len(sts), cap(sts))

	sts = append(sts, 1)
	fmt.Printf("now len %d,now cap %d ===\n", len(sts), cap(sts))

	sts = append(sts, 1)
	fmt.Printf("now len %d,now cap %d ===\n", len(sts), cap(sts))

	sts = append(sts, 1)
	fmt.Printf("now len %d,now cap %d ===\n", len(sts), cap(sts))

	sts = append(sts, 1)
	fmt.Printf("now len %d,now cap %d ===\n", len(sts), cap(sts))
	sts = append(sts, 1)
	fmt.Printf("now len %d,now cap %d ===\n", len(sts), cap(sts))

}

func newAndMake() {
	var i *int
	var a int = 10
	i = &a

	fmt.Println(*i)

}
func SliceStart() {
	//truncate()
	//newAndMake()
	//ByteToString()
	rangeSlice()
}

func SliceSplit() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//s1 := slice[2:5]
	s1 := slice
	//fmt.Println(s1, len(s1),cap(s1))
	s2 := s1[2:6:7]
	fmt.Println(s2, len(s2), cap(s2))

	s2 = append(s2, 100)
	s2 = append(s2, 200)

	s1[2] = 20

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(slice)


	ssp:=make([]string, 10)
	fmt.Println(len(ssp),cap(ssp))
	ssp =append(ssp, "spppp")
	fmt.Println(ssp)
	fmt.Println(len(ssp),cap(ssp))

	json.Marshal("ss")

}


func ByteToString() {
	ss := []byte("sfdsfjsdkfjlsdfjsdlnslcnlsdfjlncnsdlkjflwefjwelfnl1831u482fjkencsd")


	sp := *(*string)(unsafe.Pointer(&ss))
	sps:=sp[2:10]
	spp :=sps+"111"
	fmt.Println(spp)
}

func rangeSlice(){
	/*var m = []int{1, 2, 3}
	for i := range m {
		go func(i int) {
			fmt.Print(i)
		}(i)

	}
	//block main 1ms to wait goroutine finished
	time.Sleep(time.Millisecond)*/
	var arr = [10]int{1, 1, 1}
	for i, n := range &arr {
		//just ignore i and n for simplify the example
		_ = i
		_ = n
		fmt.Println(i,"===",n)
	}
}


func SliceSprintf() {
	s := []byte("")

	s1 := append(s, 'a')
	s2 := append(s, 'b')

	// 如果有此行，打印的结果是 a b，否则打印的结果是b b
	fmt.Println(s1, "===", s2)
	fmt.Println(string(s1), string(s2))
}
