package main

import (
	"fmt"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	//s := "gopher"
	//fmt.Printf("Hello and welcome, %s!\n", s)
	//
	//for i := 1; i <= 5; i++ {
	//	fmt.Println("i =", 100/i)
	//}
	//
	//var X uint8 = 225
	//fmt.Println(X, X-3)
	//
	//var Y int16 = 32767
	//fmt.Println(Y+2, Y-2)
	//
	//for i, j := range "XabCd" {
	//	fmt.Printf("The index number of %U is %d\n", j, i)
	//}
	//
	//mmap := map[int]string{
	//	22: "Geeks",
	//	33: "GFG",
	//	44: "GeeksforGeeks",
	//}
	//for key, value := range mmap {
	//	fmt.Println(key, value)
	//}

	chnl := make(chan int)
	go func() {
		chnl <- 100
		chnl <- 1000
		chnl <- 10000
		chnl <- 100000
		close(chnl)
	}()
	for i := range chnl {
		fmt.Println(i)
	}

	baba := func() {

		fmt.Println("Welcome! to GeeksforGeeks")
	}

	baba()

}
