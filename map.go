package main

import "fmt"

func main() {

	fmt.Println("Array of map")
	var a = map[string]string{"name": "Ruhul", "Age": "30"}
	var b = map[string]string{"name": "Emran", "Age": "96"}

	fmt.Println(a)
	fmt.Println(b)

	var arr = [3]map[string]string{}

	arr[0] = a
	arr[1] = b
	arr[2] = b

	fmt.Println(arr)
	fmt.Println(arr[0]["name"])

	fmt.Println()
	fmt.Println("Array of array")

	var arr1 = [3]string{"I am", "your", "friend"}
	var arr2 = [3]string{"I am", "no body"}
	var arr3 = [3]string{"What's", "that", "mean"}
	fmt.Println(arr1, arr2, arr3)

	var arrayOfArray = [3][3]string{arr1, arr2, arr3}

	fmt.Println(arrayOfArray)

	arrayOfArray[0][2] = "bandor"

	arrayOfArray[2] = [3]string{"KK", "SS", "99"}

	fmt.Println(arrayOfArray)
	fmt.Println(arrayOfArray[0])

	fmt.Println()
	fmt.Println("Map of Arrays/Slices")

	var moa = map[string][2]string{
		"Jan": [2]string{"a", "b"},
		"Feb": [2]string{"c", "d"},
	}

	fmt.Println(moa)

	mos := map[string][]string{
		"Apr": []string{"a", "b"},
		"May": []string{"c", "d", "e", "f", "g"},
	}

	fmt.Println(mos)

	fmt.Println()
	fmt.Println("Map of Map")

	var mom = map[string]map[string]string{
		"Jan": a,
		"Feb": b,
		"Mar": map[string]string{"key": "value", "key2": "value2"},
	}

	fmt.Println(mom)

	for i := range mom {
		fmt.Printf("(%s) %v\n", i, mom[i])
		for j := range mom[i] {
			fmt.Printf(" (%s) %s => %s\n", i, j, mom[i][j])
		}
	}

	//var a = make(map[string]string) // The map is empty now
	//a["brand"] = "Ford"
	//a["model"] = "Mustang"
	//a["year"] = "1964"
	//// a is no longer empty
	//
	//b := make(map[string]int)
	//b["Oslo"] = 1
	//b["Bergen"] = 2
	//b["Trondheim"] = 3
	//b["Stavanger"] = 4
	//
	//c := make(map[int]int)
	//c[10] = 1
	//c[11] = 2
	//c[12] = 3
	//c[13] = 4
	//
	//fmt.Printf("a\t%v\n", a)
	//fmt.Printf("b\t%v\n", b)
	//fmt.Printf("c\t%v\n", c[11])

	//var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964", "day": ""}
	//
	//val1, ok1 := a["brand"]
	//val2, ok2 := a["color"]
	//val3, ok3 := a["day"]
	//_, ok4 := a["model"]
	//
	//fmt.Println(val1, ok1)
	//fmt.Println(val2, ok2)
	//fmt.Println(val3, ok3)
	//fmt.Println(ok4)
	//
	//m := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}
	//
	//var b []string // defining the order
	//b = append(b, "one", "two", "three", "four")
	//
	//for k, v := range m {
	//	fmt.Printf("%v : %v, \n", k, v)
	//}
	//
	//fmt.Println()
	//
	//for _, element := range b { // loop with the defined order
	//	fmt.Printf("%v : %v, \n", element, a[element])
	//}

	//ch := make(chan string)
	//
	//// Start concurrent routines
	//go push("Moe", ch)
	//go push("Larry", ch)
	//go push("Curly", ch)
	//
	//// Read 3 results
	//// (Since our goroutines are concurrent,
	//// the order isn't guaranteed!)
	//fmt.Println(<-ch, <-ch, <-ch)

}

func push(name string, ch chan string) {
	msg := "Hey, " + name
	ch <- msg
}
