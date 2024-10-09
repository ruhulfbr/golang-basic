package main

import (
	"fmt"
)

var a int
var b int = 2
var c = 3

const PI float32 = 3.14

const (
	A int     = 1
	B float32 = 3.14
	C string  = "Hi!"
)

func main() {
	//s := "gopher"
	//fmt.Printf("Hello and welcome, %s!\n", s)
	//fmt.Println("I do not love GOLNG")
	//
	//var a string
	//var b int
	//var c bool
	//d := "aaa"
	//
	//fmt.Println("a", a)
	//fmt.Println("b", b)
	//fmt.Println("c", c)
	//fmt.Println("d", d)
	//fmt.Println("PI", PI)
	//
	//fmt.Println(A)
	//fmt.Println(B)
	//fmt.Println(C)
	//
	//var m string = "Hello"
	//var k int = 15
	//
	//fmt.Printf("m has value: %v and type: %T\n", m, m)
	//fmt.Printf("k has value: %v and type: %T\n", k, k)
	//fmt.Println("Sona boron pakhi re amar kajol boron akhi")
	//
	//var i = 15.5
	//var txt = "Hello World!"
	//
	//fmt.Printf("%v\n", i)
	//fmt.Printf("%#v\n", i)
	//fmt.Printf("%v%%\n", i)
	//fmt.Printf("%T\n", i)
	//
	//fmt.Printf("%v\n", txt)
	//fmt.Printf("%#v\n", txt)
	//fmt.Printf("%T\n", txt)
	//
	//// Array
	//fmt.Println()
	//fmt.Println("Learning array:")
	//fmt.Println()
	//
	//var arr1 = [3]int{1, 2, 3}
	//arr2 := [5]int{1, 2, 3}
	//arr3 := [5]string{"1", "2", "3"}
	//arr4 := [...]int{1, 2, 3}
	//
	//fmt.Println(arr1)
	//fmt.Println(arr2)
	//fmt.Printf("%v\n", arr3)
	//fmt.Printf("%v\n", arr4)
	//
	//// Slice
	//fmt.Println()
	//fmt.Println("Learning Slice:")
	//fmt.Println()
	//
	//myslice := [4]int{1, 2, 3}
	//fmt.Println(myslice, cap(myslice), len(myslice))
	//
	//fmt.Println("Create slice from an array:")
	//
	//array := [6]int{10, 11, 12, 13, 14, 15}
	//slice := array[2:6]
	//
	//fmt.Printf("myslice = %v\n", slice)
	//fmt.Printf("length = %d\n", len(slice))
	//fmt.Printf("capacity = %d\n", cap(slice))
	//
	//myslice2 := make([]int, 5, 11)
	//fmt.Printf("myslice2 = %v\n", myslice2)
	//fmt.Printf("length = %d\n", len(myslice2))
	//fmt.Printf("capacity = %d\n", cap(myslice2))
	//
	//prices := []int{10, 20, 30}
	//prices[2] = 50
	//
	//fmt.Println(prices[0])
	//fmt.Println(prices[2])
	//
	//myslice1 := []int{1, 2, 3, 4, 5, 6}
	//fmt.Printf("myslice1 = %v\n", myslice1)
	//fmt.Printf("length = %d\n", len(myslice1))
	//fmt.Printf("capacity = %d\n", cap(myslice1))
	//
	//myslice1 = append(myslice1, 20, 21)
	//fmt.Printf("myslice1 = %v\n", myslice1)
	//fmt.Printf("length = %d\n", len(myslice1))
	//fmt.Printf("capacity = %d\n", cap(myslice1))
	//
	//myslice4 := []int{1, 2, 3}
	//myslice5 := []int{4, 5, 6}
	//myslice6 := append(myslice4, myslice5...)
	//
	//fmt.Printf("myslice6=%v\n", myslice6)
	//fmt.Printf("length=%d\n", len(myslice6))
	//fmt.Printf("capacity=%d\n", cap(myslice6))
	//
	//// Create and use functions
	//fmt.Println(myFunction(1, 2.99))
	//fmt.Println(myFunction2(5, "Hello"))

	//_, d := myFunction2(5, "Hello")
	//fmt.Println(d)
	//
	//a, _ := myFunction2(5, "Hello")
	//fmt.Println(a)
	//
	//// Recursion Functions
	//testCount(1)

	fmt.Println(factorialRecursion(4))

}

func factorialRecursion(x float64) (y float64) {
	if x > 0 {
		y = x * factorialRecursion(x-1)
	} else {
		y = 1
	}
	return
}

func testCount(x int) int {
	if x == 11 {
		return 0
	}
	fmt.Println(x)
	return testCount(x + 1)
}

func myFunction2(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}

func myFunction(x float32, y float32) (result float32) {
	result = x + y
	return
}
