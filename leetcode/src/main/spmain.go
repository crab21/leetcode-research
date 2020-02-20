package main

import (
	"bytes"
	"container/list"
	"fmt"
	"io"
	"log"
	"pipline"
	"strings"
	"time"
)
func sm() {
	now := time.Now().Unix()
	fmt.Print(now)
	i := list.New()
	for sp := 0; sp < 9999999; sp = sp + 1 {
		i.PushFront(sp)
	}

	producer := pipline.Producer(i)
	/*ch := pipline.Squre(producer)

	for ret := range ch {
		fmt.Println( ret)
	}*/

	c1 := pipline.Squre(producer)
	c2 := pipline.Squre(producer)
	c3 := pipline.Squre(producer)
	c4 := pipline.Squre(producer)
	c5 := pipline.Squre(producer)
	c6 := pipline.Squre(producer)
	c7 := pipline.Squre(producer)
	c8 := pipline.Squre(producer)
	c9 := pipline.Squre(producer)
	c10 := pipline.Squre(producer)
	c11 := pipline.Squre(producer)
	c12 := pipline.Squre(producer)
	c13 := pipline.Squre(producer)
	c14 := pipline.Squre(producer)

	ret := pipline.Merge(c1, c2, c3, c4, c5, c6, c7, c8, c9, c10, c11, c12, c13, c14)
	for rr := range ret {
		fmt.Println(rr)
	}
	// consumer
	//for ret := range pipline.Merge(c1, c2, c3) {
	//fmt.Println(ret)
	//}
	fmt.Println()

	fmt.Println("======", time.Now().Unix()-now)

}

func mpline() {
	n := map[string]int{
		"b": 1,
		"a": 2,
	}
	if v, ok := n["a"]; ok {
		print(v)
	}
}

func smp() {
	/*	var i byte
			i='a'
			fmt.Printf("%c",i)
			fmt.Println()

			var a uint8 = 96
			fmt.Printf("%08b",a)

			fmt.Println()



			var mpp = make(map[string]string)
			var sp = &mpp

			(*sp)["w"]="1"
			(*sp)["p"]="p"
			(*sp)["y"]="y"

			for _,v :=range (*sp){
				v = "ssss"
				fmt.Println(v)
			}
			fmt.Println((*sp))

		*//*
		var a int = 1
		//fmt.Printf("%b", a)
		fmt.Printf("%0[1]d", 30, a)*/

}

func multiwriter() {
	r := strings.NewReader("some io.Reader stream to be read\n")

	var buf1, buf2 bytes.Buffer
	w := io.MultiWriter(&buf1, &buf2)

	if _, err := io.Copy(w, r); err != nil {
		log.Fatal(err)
	}

	fmt.Print(buf1.String())

	fmt.Print(buf2.String())
}

func floatLearn() {
	var a, b, c float64
	a = 0.1
	b = 0.7
	fmt.Println(a)
	c = a + b
	fmt.Println(c)
	sprintf := fmt.Sprintf("%.2f", c)
	fmt.Println(sprintf)

}

type Student struct {
	Age int
}



func sp(ss *[]string) {

	//multiwriter()
	//floatLearn()
	/*var sps = []string{"1", "2", "3"}
	fmt.Println(sps)
	sp(&sps)
	fmt.Println(sps)*/
	//base.SliceSplit() ()
	//A1()
	/*sortAlgorithm.Aug()*/
	//1010
	//s :=fmt.Sprintf("%b",(10))
	/*fmt.Println(21& 7)
	sortAlgorithm.MakeListNode()*/
	/*kv := map[string]Student{"menglu": {Age: 21}}
	student := kv["menglu"]
	student.Age = 22

	s := []Student{{Age: 21}}
	s[0].Age = 22
	fmt.Println(kv, s)*/
	//ListMain()
	//var ss = []string{"ffsdfds"}
	//base.T1()
	//sp(&ss)
	*ss = append(*ss, "fsfsd111111111")
	fmt.Println(*ss)
}

