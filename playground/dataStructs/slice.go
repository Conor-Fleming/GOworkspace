package main

import "fmt"

var x = []int{4, 5, 6, 7, 8, 9}

func slice() {
	//x := type{values} //composite literal format
	//x := []int{4, 5, 6, 7, 8, 9}
	fmt.Println(x)
}

func sliceRng() {
	for i := 0; i < len(x); i++ {
		fmt.Print(x[i], "\n")
	}

	for i, v := range x { //first value(i) will access index, second value(v) accesses value at index
		fmt.Println(i, v)
	}
}

func slicing() { //allows you do specify a starting index and ending index with ':'
	fmt.Println(x[1:4]) //access from index 1 to index 4
	fmt.Println(x[3:])  //access from index 3 through end of slice
}

func appndSlice() {
	x = append(x, 0, 100) //append values to slice x
	fmt.Println(x)
	fmt.Println("append tester slice to slice x")
	tester := []int{20, 25, 30} // declaring new slice to append onto x
	x = append(x, tester...)    //appending tester slice to x
	fmt.Println(x)

}
